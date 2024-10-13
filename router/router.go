package router

import (
    "github.com/gin-gonic/gin"
    "gitlab.com/gym-partner1/api/gym-partner-api/core"
    "gitlab.com/gym-partner1/api/gym-partner-api/interfaces/controller"
)

func Router(db *core.Database) *gin.Engine {
	route := gin.Default()

    userController := controller.NewUserController(db.Handler)

    route.GET("/ping", userController.PING)
//
//	route.AddCORS()
//
//	route.AddRoute(" Get All Users ", "GET", "/user/getAll", userController.GetAll, logServ)
//	route.AddRoute(" Get One User ", "GET", "/user/getOne", userController.GetOne, logServ)
//	route.AddRoute(" Create User ", "POST", "/user/create", userController.Create, logServ)
//	route.AddRoute(" Update User ", "PATCH", "/user/update", userController.Update, logServ)
//	route.AddRoute(" Delete User ", "DELETE", "/user/delete", userController.Delete, logServ)

	return route
}