//go:build wireinject
package di

import (
	dbconnection "zocket-task/pkg/dbconnection"
	http "zocket-task/pkg/api"
    "zocket-task/pkg/api/handlers"
	"zocket-task/pkg/config"
	"zocket-task/pkg/controller"
	"zocket-task/pkg/datalayer"

	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(dbconnection.ConnectDatabase,

		datalayer.NewUserAuthenticationDL,
		controller.NewAuthController,
		handlers.NewAuthHandler,

		datalayer.NewTaskDL,
		controller.NewTaskController,
		handlers.NewTaskHandler,

		http.NewServerHTTP,
	)
	return &http.ServerHTTP{}, nil
}
