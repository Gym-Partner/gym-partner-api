package main

import (
    "fmt"
    "github.com/spf13/viper"
    "gitlab.com/gym-partner1/api/gym-partner-api/core"
    "gitlab.com/gym-partner1/api/gym-partner-api/router"
)

func main() {
	env := core.NewEnv()
	env.LoadEnv()

	log := core.NewLog(env.FilePath)
	log.ChargeLog()

	db := core.NewDatabase(log)

	route := router.Router(db)
	address := viper.GetString("API_SERVER_HOST") + ":" + viper.GetString("API_SERVER_PORT")

//	if err := http.ListenAndServeTLS(address, viper.GetString("API_FULLCHAIN"), viper.GetString("API_PRIVKEY"), route.Handler()); err != nil {
//		log.Error(fmt.Sprintf("[RUN] %s", err.Error()))
//	}

	if err := route.RunTLS(address, viper.GetString("API_FULLCHAIN"), viper.GetString("API_PRIVKEY")); err != nil {
		log.Error(fmt.Sprintf("[RUN] %s", err.Error()))
	}
}