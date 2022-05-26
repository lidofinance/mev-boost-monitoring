package users

import (
	"context"

	"github.com/lidofinance/mev-boost-monitoring/internal/pkg/users/entity"
)

type Usecase interface {
	Get(ctx context.Context, ID int64) (*entity.User, error)
}
