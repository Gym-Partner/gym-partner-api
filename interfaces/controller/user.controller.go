package controller

import (
    "gitlab.com/Titouan-Esc/api_common/logger"
    "gitlab.com/Titouan-Esc/api_common/mongo"
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

}