package mev_boost

import (
	"context"

	"github.com/lidofinance/mev-boost-monitoring/internal/pkg/mev_boost/entity"
)

type Repository interface {
	Create(ctx context.Context, in *entity.RelayPayload) error
	Get(ctx context.Context) ([]entity.RelayPayload, error)
}
