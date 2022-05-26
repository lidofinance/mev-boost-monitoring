package usecase

import (
	"context"

	"github.com/lidofinance/mev-boost-monitoring/internal/pkg/mev_boost"
	"github.com/lidofinance/mev-boost-monitoring/internal/pkg/mev_boost/entity"
)

type usecase struct {
	repo mev_boost.Repository
}

func New(repo mev_boost.Repository) mev_boost.Usecase {
	return &usecase{
		repo: repo,
	}
}

func (u *usecase) Get(ctx context.Context, ID int64) (*entity.Payload, error) {
	return u.repo.Get(ctx, ID)
}
