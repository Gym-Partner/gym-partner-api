package main

import (
	"github.com/Gym-Partner/api-common/config"
	"github.com/Gym-Partner/api-common/router"
	"gitlab.com/gym-partner1/api/gym-partner-api/internal/delivery"
)

func main() {
	conf := config.InitConfig(config.Options{
		EnableDatabase: true,
		Migrations:     []any{},
		IsTest:         false,
	})

	r := router.InitRouter(router.Options{
		Deps: &router.Dependencies{
			Database: conf.Database,
		},
		RegisterRoutes: delivery.RegisterRoutes,
	})

	conf.Run(r)
}
