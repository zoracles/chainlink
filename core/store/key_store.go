package store

import (
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"math/big"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/tidwall/gjson"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/store/models"
	"github.com/smartcontractkit/chainlink/core/utils"
	"go.uber.org/multierr"
)

// KeyStore manages a key storage directory on disk.
type KeyStore struct {
	*keystore.KeyStore
	keysDir  string // Path to directory containing ethereum key material
	PrivKeys map[common.Address]*ecdsa.PrivateKey
}

// NewKeyStore creates a keystore for the given directory.
func NewKeyStore(keyDir string) *KeyStore {
	ks := keystore.NewKeyStore(
		keyDir,
		keystore.StandardScryptN,
		keystore.StandardScryptP,
	)
	return &KeyStore{ks, keyDir, make(map[common.Address]*ecdsa.PrivateKey)}
}

// HasAccounts returns true if there are accounts located at the keystore
// directory.
func (ks *KeyStore) HasAccounts() bool {
	return len(ks.Accounts()) > 0
}

// privateKey searches ks.keysDir for key material matching account a, and
// attempts to decrypt it and return the corresponding private key.
func (ks *KeyStore) keyDataFor(a accounts.Account, phrase string) (*ecdsa.PrivateKey, error) {
	keyFiles, err := ioutil.ReadDir(ks.keysDir)
	if err != nil {
		return nil, errors.Wrapf(err, "while trying to list key files in %s", ks.keysDir)
	}
	for _, file := range keyFiles {
		fullPath := filepath.Join(ks.keysDir, file.Name())
		contents, err := ioutil.ReadFile(fullPath)
		if err != nil {
			return nil, errors.Wrapf(err, "while trying to read keyfile %s", fullPath)
		}
		if a.Address == common.HexToAddress(gjson.Get(string(contents), "address").String()) {
			key, err := keystore.DecryptKey(contents, phrase)
			if err != nil {
				return nil, errors.Wrapf(err, "while decrypting keydata")
			}
			if a.Address != crypto.PubkeyToAddress(key.PrivateKey.PublicKey) {
				return nil, fmt.Errorf("purported key material for %s has wrong address", a.Address)
			}
			return key.PrivateKey, nil
		}
	}
	return nil, fmt.Errorf("failed to find key file for %s", a.Address)
}

// Unlock uses the given password to try to unlock accounts located in the
// keystore directory, and recover their private keys to the KeyStore
// PrivateKeys map.
func (ks *KeyStore) Unlock(phrase string) error {
	var merr error
	for _, account := range ks.Accounts() {
		err := ks.KeyStore.Unlock(account, phrase)
		if err != nil {
			merr = multierr.Combine(merr, fmt.Errorf("invalid password for account %s", account.Address.Hex()), err)
		} else {
			logger.Infow(fmt.Sprint("Unlocked account ", account.Address.Hex()), "address", account.Address.Hex())
		}
		key, err := ks.keyDataFor(account, phrase)
		if err != nil {
			return errors.Wrapf(err, "while retrieving key material")
		}
		ks.PrivKeys[account.Address] = key
	}
	return merr
}

// NewAccount adds an account to the keystore
func (ks *KeyStore) NewAccount(passphrase string) (accounts.Account, error) {
	account, err := ks.KeyStore.NewAccount(passphrase)
	if err != nil {
		return accounts.Account{}, err
	}

	err = ks.KeyStore.Unlock(account, passphrase)
	if err != nil {
		return accounts.Account{}, err
	}

	return account, nil
}

// SignTx uses the unlocked account to sign the given transaction.
func (ks *KeyStore) SignTx(account accounts.Account, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
	return ks.KeyStore.SignTx(account, tx, chainID)
}

// Sign creates an HMAC from some input data using the account's private key
func (ks *KeyStore) Sign(input []byte) (models.Signature, error) {
	account, err := ks.GetFirstAccount()
	if err != nil {
		return models.Signature{}, err
	}
	hash, err := utils.Keccak256(input)
	if err != nil {
		return models.Signature{}, err
	}

	output, err := ks.KeyStore.SignHash(account, hash)
	if err != nil {
		return models.Signature{}, err
	}
	var signature models.Signature
	signature.SetBytes(output)
	return signature, nil
}

// GetFirstAccount returns the unlocked account in the KeyStore object. The client
// ensures that an account exists during authentication.
func (ks *KeyStore) GetFirstAccount() (accounts.Account, error) {
	if len(ks.Accounts()) == 0 {
		return accounts.Account{}, errors.New("No Ethereum Accounts configured")
	}
	return ks.Accounts()[0], nil
}

// GetAccounts returns all accounts
func (ks *KeyStore) GetAccounts() []accounts.Account {
	return ks.Accounts()
}
