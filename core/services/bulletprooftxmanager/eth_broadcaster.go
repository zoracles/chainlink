package bulletprooftxmanager

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/store"
	"github.com/smartcontractkit/chainlink/core/store/models"
	"github.com/smartcontractkit/chainlink/core/store/orm"

	gethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

const (
	// databasePollInterval indicates how long to wait each time before polling
	// the database for new eth_transactions to send
	databasePollInterval = 1 * time.Second

	// EthBroadcaster advisory lock class ID
	ethBroadcasterAdvisoryLockClassID = 0
)

type EthBroadcaster interface {
	Start() error
	Stop() error

	ProcessUnbroadcastEthTransactions(models.Key) error
}

// ethBroadcaster monitors eth_transactions for transactions that need to
// be broadcast, assigns nonces and ensures that at least one eth node
// somewhere has received the transaction successfully.
//
// This does not guarantee delivery! A whole host of other things can
// subsequently go wrong such as transctions being evicted from the mempool,
// eth nodes going offline etc. Responsibility for ensuring eventual inclusion
// into the chain falls on the shoulders of the ethConfirmer.
//
// What ethBroadcaster does guarantee is:
// - a monotic series of increasing nonces for eth_transactions that can all eventually be confirmed if you retry enough times
// - existence of a saved eth_transaction_attempt
type ethBroadcaster struct {
	store  *store.Store
	config orm.ConfigReader

	started    bool
	stateMutex sync.RWMutex

	chStop chan struct{}
	chDone chan struct{}
}

func NewEthBroadcaster(store *store.Store, config orm.ConfigReader) EthBroadcaster {
	return &ethBroadcaster{
		store:  store,
		config: config,
		chStop: make(chan struct{}),
		chDone: make(chan struct{}),
	}
}

func (eb *ethBroadcaster) Start() error {
	if !eb.config.EnableBulletproofTxManager() {
		return nil
	}

	eb.stateMutex.Lock()
	defer eb.stateMutex.Unlock()
	if eb.started {
		return errors.New("already started")
	}
	go eb.monitorEthTransactions()
	eb.started = true

	return nil
}

func (eb *ethBroadcaster) Stop() error {
	eb.stateMutex.Lock()
	defer eb.stateMutex.Unlock()
	if !eb.started {
		return nil
	}
	eb.started = false
	close(eb.chStop)
	<-eb.chDone

	return nil
}

func (eb *ethBroadcaster) monitorEthTransactions() {
	defer close(eb.chDone)
	for {
		pollDatabaseTimer := time.NewTimer(databasePollInterval)

		keys, err := eb.store.Keys()

		if err != nil {
			logger.Error(err)
		} else {
			var wg sync.WaitGroup

			// It is safe to process separate keys concurrently
			// NOTE: This design will block one key if another takes a really long time to execute
			for _, key := range keys {
				if key == nil {
					logger.Error("key was unexpectedly nil. This should never happen")
					continue
				}
				wg.Add(1)
				go func(k models.Key) {
					if err := eb.ProcessUnbroadcastEthTransactions(k); err != nil {
						// NOTE: retries if this function errors are unbounded,
						// since they can be due to things like network errors
						// etc
						logger.Errorf("Error in ProcessUnbroadcastEthTransactions: %s", err)
					}
					wg.Done()
				}(*key)
			}

			wg.Wait()
		}

		select {
		case <-eb.chStop:
			return
		// TODO: can add <-eb.trigger channel for allowing other goroutines to manually trigger it early
		case <-pollDatabaseTimer.C:
			continue
		}
	}
}

func (eb *ethBroadcaster) ProcessUnbroadcastEthTransactions(key models.Key) error {
	ctx := context.Background()
	conn, err := eb.store.GetRawDB().DB().Conn(ctx)
	if err != nil {
		logger.Error(err)
		return err
	}
	defer conn.Close()
	if err := eb.lock(ctx, conn, key.ID); err != nil {
		return err
	}
	defer eb.unlock(ctx, conn, key.ID)
	return eb.processUnbroadcastEthTransactions(key.Address.Address())
}

// TODO: write this doc
// NOTE: This MUST NOT be run concurrently for the same address or it could
// result in undefined state or deadlocks.
func (eb *ethBroadcaster) processUnbroadcastEthTransactions(fromAddress gethCommon.Address) error {
	logger.Debugf("ProcessUnbroadcastEthTransactions start for %s", fromAddress.Hex())

	if err := eb.handleAnyUnfinishedEthTransaction(fromAddress); err != nil {
		return err
	}

	for {
		etx, err := nextUnbroadcastTransactionWithNonce(eb.store, fromAddress)
		if err != nil {
			// Break loop
			return err
		}
		if etx == nil {
			logger.Debugf("ProcessUnbroadcastEthTransactions finish for %s", fromAddress.Hex())
			// Finished
			return nil
		}

		gasPrice := eb.config.EthGasPriceDefault()
		etxAttempt := &models.EthTransactionAttempt{}
		sendError := send(eb.store, etx, etxAttempt, gasPrice)

		if sendError.Fatal() {
			etx.Error = sendError.StrPtr()
			err := saveFatallyErroredTransaction(eb.store, etx)
			if err != nil {
				return err
			}
			continue
		} else if sendError.isNonceAlreadyUsedError() {
			if err := eb.handleExternalWalletUsedNonce(etx, etxAttempt); err != nil {
				return err
			}
			continue
		} else if sendError != nil {
			return sendError.err
		}

		if err := saveBroadcastTransaction(eb.store, etx, etxAttempt); err != nil {
			return err
		}
	}
}

// TODO: docs
func (eb *ethBroadcaster) handleAnyUnfinishedEthTransaction(fromAddress gethCommon.Address) error {
	unfinishedEthTransaction, err := getUnfinishedEthTransaction(eb.store, fromAddress)
	if err != nil {
		return err
	}
	if unfinishedEthTransaction != nil {
		if err := eb.handleUnfinishedEthTransaction(unfinishedEthTransaction); err != nil {
			return err
		}
	}
	return nil
}

// TODO: Document exactly what the potential implications are from this
func (eb *ethBroadcaster) handleExternalWalletUsedNonce(etx *models.EthTransaction, etxAttempt *models.EthTransactionAttempt) error {
	// At all costs we avoid possible gaps in the nonce sequence. This means we may fail to send transactions, or send them twice and have one revert
	logger.Errorf("nonce of %v was too low for eth_transaction %v. Address %s has been used by another wallet. This is NOT SUPPORTED by chainlink and can lead to lost or reverted transactions.", *etx.Nonce, etx.ID, etx.FromAddress.String())

	clonedEtx := cloneForRebroadcast(etx)

	return eb.store.Transaction(func(db *gorm.DB) error {
		// Handle this case by assuming the particular transaction is broadcast already and handing off to the confirmer
		// We MUST do this to avoid gaps in the nonce sequence

		// We cannot know when the transaction was broadcast so just assume it was at the time of creation
		broadcastAt := etx.CreatedAt
		etx.BroadcastAt = &broadcastAt
		if err := saveBroadcastTransaction(eb.store, etx, etxAttempt); err != nil {
			return err
		}
		return db.Save(&clonedEtx).Error
	})
}

// getUnfinishedEthTransaction returns either 0 or 1 transaction that was left in
// an unfinished state because something went screwy the last time. Most likely
// the node crashed in the middle of the ProcessUnbroadcastEthTransactions loop.
// It may or may not have been broadcast to an eth node.
func getUnfinishedEthTransaction(store *store.Store, fromAddress gethCommon.Address) (*models.EthTransaction, error) {
	etx := &models.EthTransaction{}
	err := store.GetRawDB().First(etx, "from_address = ? AND broadcast_at IS NULL AND nonce IS NOT NULL", fromAddress.Bytes()).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}
	return etx, err
}

// TODO: docs
func (eb *ethBroadcaster) handleUnfinishedEthTransaction(ethTransaction *models.EthTransaction) error {
	gasPrice := eb.config.EthGasPriceDefault()
	ethTransactionAttempt := &models.EthTransactionAttempt{}

	sendError := send(eb.store, ethTransaction, ethTransactionAttempt, gasPrice)
	if sendError.Fatal() {
		errString := sendError.Error()
		ethTransaction.Error = &errString
		return saveFatallyErroredTransaction(eb.store, ethTransaction)
	} else if sendError.isNonceAlreadyUsedError() {
		logger.Warnf("A transaction with nonce %v has already been confirmed. Either the node crashed on a previous run, or address %s has been used by another wallet. Assuming transaction was sent successfully", *ethTransaction.Nonce, ethTransaction.FromAddress.String())
		// Cannot really know BroadcastAt for certain since the node could have crashed an indeterminate time ago
		// CreatedAt is our best guess
		// NOTE: Could add additional column 'started_at' to do better but probably not very important
		broadcastAt := ethTransaction.CreatedAt
		ethTransaction.BroadcastAt = &broadcastAt
		return saveBroadcastTransaction(eb.store, ethTransaction, ethTransactionAttempt)
	} else if sendError != nil {
		return sendError
	}

	return saveBroadcastTransaction(eb.store, ethTransaction, ethTransactionAttempt)
}

// TODO: Write short doc
func nextUnbroadcastTransactionWithNonce(store *store.Store, fromAddress gethCommon.Address) (*models.EthTransaction, error) {
	ethTransaction := &models.EthTransaction{}
	if err := findNextUnbroadcastTransactionFromAddress(store.GetRawDB(), ethTransaction, fromAddress); err != nil {
		if gorm.IsRecordNotFoundError(err) {
			// Finish. No more unbroadcasted transactions left to process. Hoorah!
			return nil, nil
		}
		return nil, err
	}

	nonce, err := GetNextNonce(store.GetRawDB(), ethTransaction.FromAddress)
	if err != nil {
		return nil, err
	}
	ethTransaction.Nonce = &nonce
	if err := store.GetRawDB().Save(ethTransaction).Error; err != nil {
		return nil, err
	}
	return ethTransaction, nil
}

func findNextUnbroadcastTransactionFromAddress(tx *gorm.DB, ethTransaction *models.EthTransaction, fromAddress gethCommon.Address) error {
	return tx.
		Where("nonce IS NULL AND error IS NULL AND broadcast_at IS NULL AND from_address = ?", fromAddress).
		Order("created_at ASC, id ASC").
		First(ethTransaction).
		Error
}

func saveBroadcastTransaction(store *store.Store, ethTransaction *models.EthTransaction, attempt *models.EthTransactionAttempt) error {
	if ethTransaction.BroadcastAt == nil {
		return errors.New("broadcastAt must be set")
	}
	if ethTransaction.Nonce == nil {
		return errors.New("nonce must be set")
	}
	return store.Transaction(func(tx *gorm.DB) error {
		if err := IncrementNextNonce(tx, ethTransaction.FromAddress, *ethTransaction.Nonce); err != nil {
			logger.Error(err)
			return err
		}
		if err := tx.Save(ethTransaction).Error; err != nil {
			logger.Error(err)
			return err
		}
		err := tx.Save(attempt).Error
		if err != nil {
			logger.Error(err)
		}
		return err
	})
}

func saveTransactionWithoutNonce(store *store.Store, ethTransaction *models.EthTransaction) error {
	if ethTransaction.Nonce != nil {
		return errors.New("nonce must be nil")
	}
	if ethTransaction.BroadcastAt != nil {
		return errors.New("broadcastAt must be nil")
	}
	return store.GetRawDB().Save(ethTransaction).Error
}

func saveFatallyErroredTransaction(store *store.Store, ethTransaction *models.EthTransaction) error {
	if ethTransaction.Error == nil {
		return errors.New("error must be set")
	}
	if ethTransaction.Nonce == nil {
		return errors.New("expected transaction to have a nonce")
	}
	ethTransaction.Nonce = nil
	return store.GetRawDB().Save(ethTransaction).Error
}

// GetNextNonce returns keys.next_nonce for the given address
func GetNextNonce(db *gorm.DB, address gethCommon.Address) (int64, error) {
	var nonce *int64
	row := db.Raw("SELECT next_nonce FROM keys WHERE address = ?", address).Row()
	if err := row.Scan(&nonce); err != nil {
		logger.Error(err)
		return 0, err
	}
	return *nonce, nil
}

// IncrementNextNonce increments keys.next_nonce by 1
func IncrementNextNonce(db *gorm.DB, address gethCommon.Address, currentNonce int64) error {
	res := db.Exec("UPDATE keys SET next_nonce = next_nonce + 1 WHERE address = ? AND next_nonce = ?", address.Bytes(), currentNonce)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		// TODO: Should probably reload nonce from eth client in this case since some invariant has been violated and it's a complete disaster
		return errors.New("could not increment nonce because no rows matched query. Either the key is missing or the nonce has been modified by an external process. This is an unrecoverable error")
	}
	return nil
}

// GetDefaultAddress queries the database for the address of the primary default ethereum key
func GetDefaultAddress(store *store.Store) (gethCommon.Address, error) {
	defaultKey, err := getDefaultKey(store)
	if err != nil {
		return gethCommon.Address{}, err
	}
	return defaultKey.Address.Address(), err
}

// NOTE: We can add more advanced logic here later such as sorting by priority
// etc
func getDefaultKey(store *store.Store) (models.Key, error) {
	availableKeys, err := store.Keys()
	if err != nil {
		return models.Key{}, err
	}
	if len(availableKeys) == 0 {
		return models.Key{}, errors.New("no keys available")
	}
	return *availableKeys[0], nil
}

// TODO: Unit test?
func cloneForRebroadcast(etx *models.EthTransaction) models.EthTransaction {
	return models.EthTransaction{
		Nonce:          nil,
		FromAddress:    etx.FromAddress,
		ToAddress:      etx.ToAddress,
		EncodedPayload: etx.EncodedPayload,
		Value:          etx.Value,
		GasLimit:       etx.GasLimit,
		BroadcastAt:    nil,
	}
}

func (eb *ethBroadcaster) lock(ctx context.Context, conn *sql.Conn, keyID int32) error {
	gotLock := false
	rows, err := conn.QueryContext(ctx, "SELECT pg_try_advisory_lock($1, $2)", ethBroadcasterAdvisoryLockClassID, keyID)
	defer rows.Close()
	if err != nil {
		logger.Error(err)
		return err
	}
	gotRow := rows.Next()
	if !gotRow {
		return errors.New("query unexpectedly returned 0 rows")
	}
	if err := rows.Scan(&gotLock); err != nil {
		logger.Error(err)
		return err
	}
	if gotLock {
		return nil
	}
	return fmt.Errorf("could not get advisory lock for key %v", keyID)
}

func (eb *ethBroadcaster) unlock(ctx context.Context, conn *sql.Conn, keyID int32) error {
	_, err := conn.ExecContext(ctx, "SELECT pg_advisory_unlock($1, $2)", ethBroadcasterAdvisoryLockClassID, keyID)
	if err != nil {
		logger.Error(err)
	}
	return err
}
