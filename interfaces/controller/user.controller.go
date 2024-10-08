package controller

import (
    "gitlab.com/Titouan-Esc/api_common/controller"
    "gitlab.com/Titouan-Esc/api_common/logger"
    "gitlab.com/Titouan-Esc/api_common/mongo"
    "gitlab.com/Titouan-Esc/api_common/utils"
    "gitlab.com/gym-partner1/api/gym-partner-api/domain/model"
    "gitlab.com/gym-partner1/api/gym-partner-api/interfaces/repository"
    "gitlab.com/gym-partner1/api/gym-partner-api/usecases/interactor"
    "net/http"
)

type UserController struct {
    UserInteractor interactor.UserInteractor
    Logger *logger.Log
}

func NewUserController(sql *mongo.Mongo, log *logger.Log) *UserController {
    return &UserController{
        UserInteractor: interactor.UserInteractor{
            IUserRepository: repository.UserRepository{
                Sql: sql,
                Logger: log,
            },
        },
        Logger: log,
    }
}

func (uc *UserController) GetAll(res http.ResponseWriter, req *http.Request) {
    manager := controller.NewController(res, req, uc.Logger, false)
    if manager.Errors.Error {
        manager.StopRequest()
        return
    }

    users, err := uc.UserInteractor.GetAll()
    if err != nil {
        manager.Respons().Build(http.StatusInternalServerError, err.Error())
        return
    }

    manager.Respons().Build(http.StatusOK, users)
}

func (uc *UserController) Create(res http.ResponseWriter, req *http.Request) {
    manager := controller.NewController(res, req, uc.Logger, false)
    if manager.Errors.Error {
        manager.StopRequest()
        return
    }

    body := utils.ReadBody[model.User](manager.Body)

    user, err := uc.UserInteractor.Create(body)
    if err != nil {
        manager.Respons().Build(http.StatusInternalServerError, err.Error())
        return
    }

    mapBody := utils.StructToMap(user)

    if err = manager.Cognito.SignUp(mapBody); err != nil {
            manager.Respons().Build(http.StatusInternalServerError, err.Error())
            return
    }

    newUser := user.NewUserFromData(user)
    manager.Respons().Build(http.StatusOK, newUser)
}