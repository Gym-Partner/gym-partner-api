package main

import (
    "gitlab.com/gym-partner1/api/gym-partner-api/core"
    "gitlab.com/gym-partner1/api/gym-partner-api/router"
    "net/http"
)

func main() {
	env := core.NewEnv()
	env.LoadEnv()

	log := core.NewLog(env.FilePath)
	log.ChargeLog()

	db := core.NewDatabase(log)

	route := router.Router(db)

	http.ListenAndServe(":4200", route)
}