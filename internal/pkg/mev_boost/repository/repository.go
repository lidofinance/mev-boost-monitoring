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

func (r *repo) Get(ctx context.Context, ID int64) (*entity.Payload, error) {
	var out entity.Payload
	err := r.db.GetContext(ctx, &out, `select * from users where id = $1`, ID)

	return &out, err
}

func (r *repo) Create(ctx context.Context) (*int64, error) {
	var ID int64

	query := `insert into users () returning id;`
	if createUserErr := r.db.GetContext(ctx, &ID, query); createUserErr != nil {
		return nil, createUserErr
	}

	return &ID, nil
}
