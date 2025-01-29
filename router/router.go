package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/docs"
	"gitlab.com/gym-partner1/api/gym-partner-api/interfaces/controller"
	"gitlab.com/gym-partner1/api/gym-partner-api/middleware"
	"time"
)

func Router(db *core.Database) *gin.Engine {
	route := gin.Default()
	route.Use(middleware.InitMiddleware(db.Logger))
	docs.SwaggerInfo.BasePath = "/api/v1"

	route.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://foo.com"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	userController := controller.NewUserController(db)
	workoutController := controller.NewWorkoutController(db)
	authController := controller.NewAuthController(db)

	api := route.Group("/api")
	{
		v1Auth := api.Group("/v1", middleware.Auth())
		{
			v1Auth.GET("/user/getAll", userController.GetAll)
			v1Auth.GET("/user/getOne", userController.GetOne)
			v1Auth.PATCH("/user/update", userController.Update)
			v1Auth.DELETE("/user/delete", userController.Delete)
			v1Auth.GET("/user/workout/getOne", workoutController.GetOneByUserId)
			v1Auth.POST("/workout/create", workoutController.Create)
			v1Auth.POST("/auth/refresh", authController.RefreshToken)
		}

		v1NoAuth := api.Group("/v1")
		{
			v1NoAuth.POST("/user/create", userController.Create)
			v1NoAuth.POST("/user/login", userController.Login)
			v1NoAuth.POST("/auth/login", authController.Login)
			v1NoAuth.GET("/ping", func(context *gin.Context) {
				context.JSON(200, gin.H{
					"message": "PONG",
				})
			})
		}
	}
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return route
}
