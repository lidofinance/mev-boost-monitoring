package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/lidofinance/mev-boost-monitoring/internal/pkg/mev_boost"
	"github.com/lidofinance/mev-boost-monitoring/internal/pkg/mev_boost/entity"
)

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) mev_boost.Repository {
	return &repo{
		db: db,
	}
}

func (r *repo) Create(ctx context.Context, in *entity.RelayPayload) error {
	query := `
insert into relay_payloads (slot, block_number, block_hash, fee_recipient, transactions_root, pubkey, signature, relay_addr, relay_timestamp) 
                    values ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	if _, err := r.db.ExecContext(ctx, query,
		in.SlotNumber, in.BlockNumber, in.BlockHash, in.FeeRecipient, in.TransactionsRoot, in.Pubkey,
		in.Signature, in.RelayAddr, in.RelayTimestamp,
	); err != nil {
		return err
	}

	return nil
}

func (r *repo) Get(ctx context.Context) ([]entity.RelayPayload, error) {
	var out []entity.RelayPayload

	query := `select * from relay_payloads limit 10`
	if err := r.db.SelectContext(ctx, &out, query); err != nil {
		return nil, err
	}

	return out, nil
}
