package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/docs"
	"gitlab.com/gym-partner1/api/gym-partner-api/interfaces/controller"
	"gitlab.com/gym-partner1/api/gym-partner-api/middleware"
)

func Router(db *core.Database) *gin.Engine {
	route := gin.Default()
	route.Use(middleware.InitMiddleware(db.Logger))
	docs.SwaggerInfo.BasePath = "/api/v1"

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
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return route
}
