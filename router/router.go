package router

import (
    "github.com/gin-gonic/gin"
    "gitlab.com/gym-partner1/api/gym-partner-api/core"
    "gitlab.com/gym-partner1/api/gym-partner-api/interfaces/controller"
)

func Router(db *core.Database) *gin.Engine {
	route := gin.Default()

    userController := controller.NewUserController(db)

    api := route.Group("/api")
    {
        v1 := api.Group("/v1")
        {
            v1.POST("/user/create", userController.Create)
            v1.GET("/user/getAll", userController.GetAll)
            v1.POST("/user/getOne", userController.GetOne)
            v1.PATCH("/user/update", userController.Update)
            v1.DELETE("/user/delete", userController.Delete)

            v1.GET("/ping", userController.PING)
        }
    }

	return route
}