package server

import (
	"github.com/jmoiron/sqlx"

	"github.com/lidofinance/mev-boost-monitoring/internal/pkg/mev_boost"
	mevBoostRepo "github.com/lidofinance/mev-boost-monitoring/internal/pkg/mev_boost/repository"
)

type repository struct {
	MevBoost mev_boost.Repository
}

//nolint
func Repository(db *sqlx.DB) *repository {
	return &repository{
		MevBoost: mevBoostRepo.New(db),
	}
}
