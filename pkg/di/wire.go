//go:build wireinject
// +build wireinject

package di

import (
	http "shiftsync/pkg/api"
	"shiftsync/pkg/api/handler"
	"shiftsync/pkg/usecases"

	"github.com/google/wire"
)

func InitializeAPI() *http.ServerHTTP {
	wire.Build(usecases.NewEmployeeUseCase, handler.NewEmployeeHandler, http.NewHTTPServer)

	return &http.ServerHTTP{}
}
