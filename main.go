package main

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/viper"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/database"
	"gitlab.com/gym-partner1/api/gym-partner-api/router"
)

func main() {
	env := core.NewEnv()
	env.LoadEnv()

	log := core.NewLog(env.FilePath, false)
	log.ChargeLog()

	db := core.NewDatabase(log)
	if err := db.ModelMigrate(
		database.MigrateUser{},
		database.MigrateWorkout{},
		database.MigrateUnityOfWorkout{},
		database.MigrateExercise{},
		database.MigrateSerie{},
		database.MigrateAuth{},
		database.MigrateFollows{},
		database.MigrateUserImage{},
	); err != nil {
		log.Error(fmt.Sprintf(core.ErrMigrateModel, err.Error()))
		return
	}
	database.PsqlIndex(db.Handler)

	route := router.Router(db)
	address := viper.GetString("API_SERVER_HOST") + ":" + viper.GetString("API_SERVER_PORT")

	gymPartnerFigure := figure.NewFigure("Gym Partner API", "slant", true)
	gymPartnerFigure.Print()
	fmt.Println()
	fmt.Printf(fmt.Sprintf("Run to: http://%s/api/v1\n", address))

	if err := route.Run(address); err != nil {
		log.Error(fmt.Sprintf("[RUN] %s", err.Error()))
	}
}
