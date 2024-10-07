package infra

import (
    "gitlab.com/Titouan-Esc/api_common/logger"
    "gitlab.com/Titouan-Esc/api_common/mongo"
    "gitlab.com/Titouan-Esc/api_common/router"
    "gitlab.com/gym-partner1/api/gym-partner-api/interfaces/controller"
)

func Dispatch(sql *mongo.Mongo, logServ, logSess *logger.Log) *router.Router {
	route := router.NewRouter()

	userController := controller.NewUserController(sql, logSess)

	route.AddCORS()

	return route
}