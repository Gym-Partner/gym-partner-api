package main

import (
	"github.com/Gym-Partner/api-common/config"
	"github.com/Gym-Partner/api-common/router"
	"gitlab.com/gym-partner1/api/gym-partner-api/internal/delivery"
	"gitlab.com/gym-partner1/api/gym-partner-api/internal/migrate"
)

func main() {
	conf := config.InitConfig(config.Options{
		EnableDatabase: true,
		Migrations: []any{
			// User Migration
			migrate.UserMigrate{},
			migrate.UserRoleMigrate{},
		},
		IsTest: false,
	})

	r := router.InitRouter(router.Options{
		Deps: &router.Dependencies{
			Database: conf.Database,
		},
		RegisterRoutes: delivery.RegisterRoutes,
	})

	conf.Run(r)
}
