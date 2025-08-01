package main

import (
	"fmt"

	"github.com/spf13/viper"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
	"gitlab.com/gym-partner1/api/gym-partner-api/router"
)

func main() {
	env := core.NewEnv()
	env.LoadEnv()

	log := core.NewLog(env.FilePath)
	log.ChargeLog()

	db := core.NewDatabase(log)
	if err := db.ModelMigrate(model.User{}, model.Workout{}, model.UnityOfWorkout{}, model.Exercice{}, model.Serie{}); err != nil {
		log.Error(fmt.Sprintf(core.ErrMigrateModel, err.Error()))
		return
	}

	route := router.Router(db)
	address := viper.GetString("API_SERVER_HOST") + ":" + viper.GetString("API_SERVER_PORT")

	if err := route.Run(address); err != nil {
		log.Error(fmt.Sprintf("[RUN] %s", err.Error()))
	}
}
