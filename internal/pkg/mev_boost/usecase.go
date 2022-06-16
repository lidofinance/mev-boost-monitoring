package mev_boost

import (
	"context"

	"github.com/lidofinance/mev-boost-monitoring/internal/pkg/mev_boost/entity"
)

type Usecase interface {
	Create(ctx context.Context, in *entity.RelayPayload) error
	Paginated(ctx context.Context, currentPage, perPage uint64) (*entity.RelayPayloadPaginated, error)
}
