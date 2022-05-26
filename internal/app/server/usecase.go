package server

import (
	"github.com/lidofinance/mev-boost-monitoring/internal/pkg/mev_boost"
	mevBoostUsecase "github.com/lidofinance/mev-boost-monitoring/internal/pkg/mev_boost/usecase"
)

type usecase struct {
	MevBoost mev_boost.Usecase
}

//nolint
func Usecase(
	repo *repository,
) *usecase {
	return &usecase{
		MevBoost: mevBoostUsecase.New(repo.MevBoost),
	}
}
