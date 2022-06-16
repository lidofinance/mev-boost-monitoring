package mev_boost

import (
	"context"

	"github.com/lidofinance/mev-boost-monitoring/internal/pkg/mev_boost/entity"
)

type Repository interface {
	Create(ctx context.Context, in *entity.RelayPayload) error
	TotalRelays(ctx context.Context) (uint64, error)
	Paginated(ctx context.Context, offset, limit uint64) ([]entity.RelayPayload, error)
}
