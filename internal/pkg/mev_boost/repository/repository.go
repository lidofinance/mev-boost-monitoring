package repository

import (
	"context"
	"encoding/json"

	"github.com/flashbots/go-boost-utils/types"
	"github.com/jmoiron/sqlx"
	"github.com/lidofinance/mev-boost-monitoring/internal/pkg/mev_boost/entity"

	"github.com/lidofinance/mev-boost-monitoring/internal/pkg/mev_boost"
)

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) mev_boost.Repository {
	return &repo{
		db: db,
	}
}

func (r *repo) Create(ctx context.Context, headerPayload types.GetHeaderResponse) error {
	data, _ := json.Marshal(*headerPayload.Data)

	query := `insert into headers (version, data) values ($1, $2)`
	if _, createUserErr := r.db.ExecContext(ctx, query, headerPayload.Version, string(data)); createUserErr != nil {
		return createUserErr
	}

	return nil
}

func (r *repo) Get(ctx context.Context) ([]entity.Header, error) {
	var out []entity.Header

	query := `select * from headers limit 10`
	if err := r.db.SelectContext(ctx, &out, query); err != nil {
		return nil, err
	}

	return out, nil
}
