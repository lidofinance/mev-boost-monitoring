package server

import (
	"github.com/lidofinance/mev-boost-monitoring/internal/pkg/users"
	usersUsecase "github.com/lidofinance/mev-boost-monitoring/internal/pkg/users/usecase"
)

type usecase struct {
	User users.Usecase
}

//nolint
func Usecase(
	repo *repository,
) *usecase {
	return &usecase{
		User: usersUsecase.New(repo.User),
	}
}
