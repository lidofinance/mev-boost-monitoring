package mev_boost

import (
	"context"

	"github.com/flashbots/go-boost-utils/types"
	"github.com/lidofinance/mev-boost-monitoring/internal/pkg/mev_boost/entity"
)

type Repository interface {
	Create(ctx context.Context, headerPayload types.GetHeaderResponse) error
	Get(ctx context.Context) ([]entity.Header, error)
}
