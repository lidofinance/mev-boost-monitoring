package usecase

import (
	"context"

	"github.com/lidofinance/mev-boost-monitoring/internal/pkg/users"
	"github.com/lidofinance/mev-boost-monitoring/internal/pkg/users/entity"
)

type usecase struct {
	repo users.Repository
}

func New(repo users.Repository) users.Usecase {
	return &usecase{
		repo: repo,
	}
}

func (u *usecase) Get(ctx context.Context, ID int64) (*entity.User, error) {
	return u.repo.Get(ctx, ID)
}
