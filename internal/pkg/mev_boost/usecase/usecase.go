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

func (u *usecase) Create(ctx context.Context, in *entity.RelayPayload) error {
	return u.repo.Create(ctx, in)
}

func (u *usecase) Get(ctx context.Context) ([]entity.RelayPayload, error) {
	return u.repo.Get(ctx)
}
