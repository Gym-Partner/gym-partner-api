package router

import (
    "github.com/gin-gonic/gin"
    "gitlab.com/gym-partner1/api/gym-partner-api/core"
    "gitlab.com/gym-partner1/api/gym-partner-api/interfaces/controller"
    "gitlab.com/gym-partner1/api/gym-partner-api/middleware"
)

func Router(db *core.Database) *gin.Engine {
	route := gin.Default()
    route.Use(middleware.InitMiddleware(db.Logger))

    userController := controller.NewUserController(db)

    api := route.Group("/api")
    {
        v1Auth := api.Group("/v1", middleware.Auth())
        {
            v1Auth.GET("/user/getAll", userController.GetAll)
            v1Auth.GET("/user/getOne", userController.GetOne)
            v1Auth.PATCH("/user/update", userController.Update)
            v1Auth.DELETE("/user/delete", userController.Delete)
        }

        v1NoAuth := api.Group("/v1")
        {
            v1NoAuth.POST("/user/create", userController.Create)
            v1NoAuth.POST("/user/login", userController.Login)
            v1NoAuth.GET("/ping", userController.PING)
        }
    }

	return route
}