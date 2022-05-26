package server

import (
	"github.com/lidofinance/go-template/internal/pkg/users"
	usersUsecase "github.com/lidofinance/go-template/internal/pkg/users/usecase"
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
