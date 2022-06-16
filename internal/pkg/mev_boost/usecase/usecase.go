package usecase

import (
	"context"
	"fmt"

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

func (u *usecase) Paginated(ctx context.Context, currentPage, perPage uint64) (*entity.RelayPayloadPaginated, error) {
	total, totalErr := u.repo.TotalRelays(ctx)
	if totalErr != nil {
		return nil, totalErr
	}

	next := ((currentPage + 1) - 1) * perPage
	prev := uint64(0)
	if currentPage > 1 {
		prev = ((currentPage - 1) - 1) * perPage
	}
	from := (currentPage - 1) * perPage
	if from == 0 {
		from = 1
	}

	to := perPage * currentPage
	last := (total - 1) * perPage

	offset := (perPage * currentPage) - perPage

	data, dbErr := u.repo.Paginated(ctx, offset, perPage)
	if dbErr != nil {
		return nil, dbErr
	}

	return &entity.RelayPayloadPaginated{
		Total:        total,
		PerPage:      perPage,
		CurrentPage:  currentPage,
		LastPage:     last,
		FirstPageURL: fmt.Sprintf(`/payload/%d/%d`, 1, perPage),
		LastPageUrl:  fmt.Sprintf(`/payload/%d/%d`, last, perPage),
		NextPageURL:  fmt.Sprintf(`/payload/%d/%d`, next, perPage),
		PrevPageURL:  fmt.Sprintf(`/payload/%d/%d`, prev, perPage),
		From:         from,
		To:           to,
		Data:         data,
	}, nil
}
