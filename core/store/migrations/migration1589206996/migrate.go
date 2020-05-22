package migration1589206996

import (
	"github.com/jinzhu/gorm"
)

// Migrate adds the requisite tables for the BulletproofTxManager
// I have tried to make an intelligent guess at the required indexes and
// constraints but this will undergo revision as the work progresses
func Migrate(tx *gorm.DB) error {
	return tx.Exec(`
		CREATE TYPE eth_transactions_attempt_state AS ENUM ('unconfirmed', 'confirmed');

	  	CREATE TABLE eth_transactions (
			id BIGSERIAL PRIMARY KEY,
			nonce bigint, 
			from_address bytea REFERENCES keys (address) NOT NULL,
			to_address bytea NOT NULL,
			encoded_payload bytea NOT NULL,
			value numeric(78, 0) NOT NULL,
			gas_limit bigint NOT NULL,
			error text,
			broadcast_at timestamptz,
			created_at timestamptz NOT NULL,
			-- NOTE: attempt_state is strictly speaking redundant but much needed from a performance standpoint
			attempt_state eth_transactions_attempt_state NOT NULL DEFAULT 'unconfirmed'::eth_transactions_attempt_state
		  );
		  
		ALTER TABLE eth_transactions ADD CONSTRAINT chk_from_address_length CHECK (
			octet_length(from_address) = 20
		);
		ALTER TABLE eth_transactions ADD CONSTRAINT chk_to_address_length CHECK (
			octet_length(to_address) = 20
		);

		CREATE UNIQUE INDEX idx_eth_transactions_nonce_from_address ON eth_transactions (nonce, from_address);
		CREATE INDEX idx_eth_transactions_attempt_state ON eth_transactions (attempt_state) WHERE attempt_state = 'unconfirmed';
		CREATE INDEX idx_eth_transactions_created_at ON eth_transactions USING BRIN (created_at);

		-- only one record allowed per account with a nonce but no broadcast_at
		CREATE UNIQUE INDEX idx_only_one_in_progress_tx_per_account ON eth_transactions (from_address) WHERE broadcast_at IS NULL AND nonce IS NOT NULL;

		ALTER TABLE eth_transactions ADD CONSTRAINT chk_nonce_requires_from_address CHECK (
			nonce IS NULL OR from_address IS NOT NULL
		);

		ALTER TABLE eth_transactions ADD CONSTRAINT chk_nonce_may_not_be_present_with_error CHECK (
			nonce IS NULL OR error IS NULL
		);

		ALTER TABLE eth_transactions ADD CONSTRAINT chk_broadcast_at_may_not_be_present_with_error CHECK (
			broadcast_at IS NULL OR error IS NULL
		);

		ALTER TABLE eth_transactions ADD CONSTRAINT chk_broadcast_at_requires_nonce CHECK (
			broadcast_at IS NULL OR nonce IS NOT NULL
		);

		CREATE TABLE eth_transaction_attempts (
			id BIGSERIAL PRIMARY KEY,
		 	eth_transaction_id bigint REFERENCES eth_transactions (id) NOT NULL,
		 	gas_price numeric(78,0) NOT NULL,
		 	signed_raw_tx bytea NOT NULL,
		 	hash bytea NOT NULL,
			broadcast_before_block_num bigint,
		 	created_at timestamptz NOT NULL
		);

		ALTER TABLE eth_transaction_attempts ADD CONSTRAINT chk_hash_length CHECK (
			octet_length(hash) = 32
		);

		CREATE UNIQUE INDEX idx_eth_transaction_attempts_hash ON eth_transaction_attempts (hash);
		CREATE INDEX idx_eth_transaction_attempts ON eth_transaction_attempts (eth_transaction_id);
		CREATE INDEX idx_eth_transactions_broadcast_before_block_num ON eth_transaction_attempts (broadcast_before_block_num);
		CREATE INDEX idx_eth_transaction_attempts_created_at ON eth_transaction_attempts USING BRIN (created_at);

		CREATE TABLE eth_task_run_transactions (
			task_run_id uuid NOT NULL REFERENCES task_runs (id) ON DELETE CASCADE,
			eth_transaction_id bigint NOT NULL REFERENCES eth_transactions (id) ON DELETE CASCADE
		);

		CREATE UNIQUE INDEX idx_eth_task_run_transactions_task_run_id ON eth_task_run_transactions (task_run_id);
		CREATE UNIQUE INDEX idx_eth_task_run_transactions_eth_transaction_id ON eth_task_run_transactions (eth_transaction_id);

		CREATE TABLE eth_receipts (
			id BIGSERIAL PRIMARY KEY,
			eth_transaction_attempt_id bigint REFERENCES eth_transaction_attempts (id) NOT NULL,
			transaction_hash bytea NOT NULL,
			block_hash bytea NOT NULL,
			block_number bigint NOT NULL,
			transaction_index bigint NOT NULL,
			receipt jsonb NOT NULL,
			created_at timestamptz NOT NULL
		);

		CREATE INDEX idx_eth_transactions_attempt_receipts_block_number ON eth_receipts (block_number);
		CREATE UNIQUE INDEX idx_unique_receipts ON eth_receipts (block_hash, transaction_hash);
		CREATE INDEX idx_eth_receipts_created_at ON eth_transaction_attempts USING BRIN (created_at);
	`).Error
}
