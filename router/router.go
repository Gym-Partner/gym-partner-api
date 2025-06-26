package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/zsais/go-gin-prometheus"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/docs"
	"gitlab.com/gym-partner1/api/gym-partner-api/interfaces/controller"
	"gitlab.com/gym-partner1/api/gym-partner-api/middleware"
)

func Router(db *core.Database, rabbit *core.RabbitMQ) *gin.Engine {
	route := gin.Default()
	route.Use(middleware.InitMiddleware(db.Logger))
	docs.SwaggerInfo.BasePath = "/api/v1"

	prometheus := ginprometheus.NewPrometheus("gin")
	prometheus.Use(route)

	route.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://foo.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	userController := controller.NewUserController(db, rabbit)
	workoutController := controller.NewWorkoutController(db)
	authController := controller.NewAuthController(db)
	followController := controller.NewFollowController(db)

	api := route.Group("/api")
	{
		v1Auth := api.Group("/v1", middleware.Auth())
		{
			// ###########################################################
			//							USER
			// ###########################################################
			v1Auth.GET("/user/get_all", userController.GetAll)
			v1Auth.GET("/user/get_one", userController.GetOne)
			v1Auth.POST("/users/get_users", userController.GetUsers)
			v1Auth.GET("/users/search", userController.Search)
			v1Auth.POST("/user/upload_image", userController.UploadImage)
			v1Auth.PATCH("/user/update", userController.Update)
			v1Auth.DELETE("/user/delete", userController.Delete)

			// ###########################################################
			//							WORKOUT
			// ###########################################################
			v1Auth.GET("/user/workout/get_one", workoutController.GetOneByUserId)
			v1Auth.GET("/user/workout/get_all", workoutController.GetAllByUserId)
			v1Auth.POST("/workout/create", workoutController.Create)
			v1Auth.PUT("/workout/update", workoutController.Update)
			v1Auth.DELETE("/workout/delete", workoutController.Delete)

			// ###########################################################
			//							AUTH
			// ###########################################################
			v1Auth.POST("/auth/refresh", authController.RefreshToken)

			// ###########################################################
			//							FOLLOWER
			// ###########################################################
			v1Auth.POST("/user/follower/add_follower", followController.AddFollower)
			v1Auth.POST("/user/follower/remove_follower", followController.RemoveFollower)
		}

		v1NoAuth := api.Group("/v1")
		{
			v1NoAuth.POST("/user/create", userController.Create)
			v1NoAuth.GET("/user/rabbitmq_test", userController.RabbitMQTest)
			v1NoAuth.POST("/auth/sign_in", authController.Login)
			v1NoAuth.POST("/auth/refresh_token", authController.RefreshToken)
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
