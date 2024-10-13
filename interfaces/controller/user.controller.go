package controller

import (
    "github.com/gin-gonic/gin"
    "gitlab.com/gym-partner1/api/gym-partner-api/interfaces/repository"
    "gitlab.com/gym-partner1/api/gym-partner-api/usecases/interactor"
    "gorm.io/gorm"
)

type UserController struct {
    UserInteractor interactor.UserInteractor
}

func NewUserController(sql *gorm.DB) *UserController {
    return &UserController{
        UserInteractor: interactor.UserInteractor{
            IUserRepository: repository.UserRepository{
                Sql: sql,
            },
        },
    }
}

//func (uc *UserController) GetAll(res http.ResponseWriter, req *http.Request) {
//    manager := controller.NewController(res, req, uc.Logger, false)
//    if manager.Errors.Error {
//        manager.StopRequest()
//        return
//    }
//
//    users, err := uc.UserInteractor.GetAll()
//    if err != nil {
//        manager.Respons().Build(http.StatusInternalServerError, err.Error())
//        return
//    }
//
//    manager.Respons().Build(http.StatusOK, users)
//}
//
//func (uc *UserController) GetOne(res http.ResponseWriter, req *http.Request) {
//    manager := controller.NewController(res, req, uc.Logger, false)
//    if manager.Errors.Error {
//        manager.StopRequest()
//        return
//    }
//
//    user, err := uc.UserInteractor.GetOneById(manager.UserId)
//    if err != nil {
//        manager.Respons().Build(http.StatusInternalServerError, err)
//        return
//    }
//
//    manager.Respons().Build(http.StatusOK, user)
//}
//
//func (uc *UserController) Create(res http.ResponseWriter, req *http.Request) {
//    manager := controller.NewController(res, req, uc.Logger, false)
//    if manager.Errors.Error {
//        manager.StopRequest()
//        return
//    }
//
//    body := utils.ReadBody[model.User](manager.Body)
//
//    user, err := uc.UserInteractor.Create(body)
//    if err != nil {
//        manager.Respons().Build(http.StatusInternalServerError, err.Error())
//        return
//    }
//
//    mapBody := utils.StructToMap(user)
//
//    if err = manager.Cognito.SignUp(mapBody); err != nil {
//            manager.Respons().Build(http.StatusInternalServerError, err.Error())
//            return
//    }
//
//    newUser := user.NewUserFromData(user)
//    manager.Respons().Build(http.StatusCreated, newUser)
//}
//
//func (uc *UserController) Update(res http.ResponseWriter, req *http.Request) {
//    manager := controller.NewController(res, req, uc.Logger, false)
//    if manager.Errors.Error {
//        manager.StopRequest()
//        return
//    }
//
//    body := utils.ReadBody[model.User](manager.Body)
//    //body.Id = manager.UserId
//
//    if err := uc.UserInteractor.Update(body); err != nil {
//        manager.Respons().Build(http.StatusInternalServerError, err)
//        return
//    }
//
//    manager.Respons().Build(http.StatusNoContent, nil)
//}
//
//func (uc *UserController) Delete(res http.ResponseWriter, req *http.Request) {
//    manager := controller.NewController(res, req, uc.Logger, false)
//    if manager.Errors.Error {
//        manager.StopRequest()
//        return
//    }
//
//    body := utils.ReadBody[model.User](manager.Body)
//
//    if err := uc.UserInteractor.Delete(body.Id); err != nil {
//        manager.Respons().Build(http.StatusInternalServerError, err)
//        return
//    }
//
//    manager.Respons().Build(http.StatusNoContent, nil)
//}

func (uc *UserController) PING(ctx *gin.Context) {
    ctx.JSON(200, gin.H{
        "message": "PONG",
    })
}