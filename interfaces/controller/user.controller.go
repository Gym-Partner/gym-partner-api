package controller

import (
    "github.com/gin-gonic/gin"
    "gitlab.com/gym-partner1/api/gym-partner-api/core"
    "gitlab.com/gym-partner1/api/gym-partner-api/interfaces/repository"
    "gitlab.com/gym-partner1/api/gym-partner-api/usecases/interactor"
)

type UserController struct {
    UserInteractor interactor.UserInteractor
}

func NewUserController(db *core.Database) *UserController {
    return &UserController{
        UserInteractor: interactor.UserInteractor{
            IUserRepository: repository.UserRepository{
                DB: db.Handler,
                Log: db.Logger,
            },
        },
    }
}

func (uc *UserController) Create(ctx *gin.Context) {
    
}

func (uc *UserController) PING(ctx *gin.Context) {
    ctx.JSON(200, gin.H{
        "message": "PONG",
    })
}