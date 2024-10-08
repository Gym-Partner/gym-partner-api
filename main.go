package main

import (
    "fmt"
    "github.com/spf13/viper"
    "gitlab.com/Titouan-Esc/api_common/env"
    "gitlab.com/Titouan-Esc/api_common/logger"
    "gitlab.com/Titouan-Esc/api_common/mongo"
    "gitlab.com/gym-partner1/api/gym-partner-api/infra"
    "net/http"
)

func main() {
	environment := env.NewEnv()
	environment.LoadEnv()

	logServ := logger.GetLoggers().Server
	logSess := logger.GetLoggers().Session

	db := mongo.NewMongo(logServ)

	router := infra.Dispatch(db, logServ, logSess)

	address := fmt.Sprintf("%s:%s", viper.GetString("API_SERVER_HOST"), viper.GetString("API_SERVER_PORT"))

	fmt.Println("Server running on : ", address + viper.GetString("API_PREFIX"))
	if err := http.ListenAndServe(address, router.Handle); err != nil {
		logServ.Error("[RUN] ", err.Error())
		return
	}
}