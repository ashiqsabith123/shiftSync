//go:build wireinject
// +build wireinject

package di

import (
	http "shiftsync/pkg/api"
	"shiftsync/pkg/api/handler"
	"shiftsync/pkg/config"
	"shiftsync/pkg/db"
	"shiftsync/pkg/repository"
	"shiftsync/pkg/usecases"

	"github.com/google/wire"
)

func InitializeAPI(config config.Config) *http.ServerHTTP {
	wire.Build(
		db.ConnectToDatbase,
		repository.NewEmployeeRepository,
		usecases.NewEmployeeUseCase,
		handler.NewEmployeeHandler,
		http.NewHTTPServer)

	return &http.ServerHTTP{}
}
