package users

import (
	"context"

	"github.com/lidofinance/mev-boost-monitoring/internal/pkg/users/entity"
)

type Repository interface {
	Get(ctx context.Context, ID int64) (*entity.User, error)
	Create(ctx context.Context) (*int64, error)
}
