package router

import (
    "github.com/gin-gonic/gin"
    "gitlab.com/gym-partner1/api/gym-partner-api/core"
    "gitlab.com/gym-partner1/api/gym-partner-api/interfaces/controller"
)

func Router(db *core.Database) *gin.Engine {
	route := gin.Default()

    userController := controller.NewUserController(db.Handler)

    route.POST("/user/create", userController.Create)
    route.GET("/ping", userController.PING)

	return route
}