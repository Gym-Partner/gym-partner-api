package controller

import (
    "github.com/gin-gonic/gin"
    "gitlab.com/gym-partner1/api/gym-partner-api/core"
    "gitlab.com/gym-partner1/api/gym-partner-api/interfaces/repository"
    "gitlab.com/gym-partner1/api/gym-partner-api/usecases/interactor"
)

type UserController struct {
    UserInteractor interactor.UserInteractor
    Log *core.Log
}

func NewUserController(db *core.Database) *UserController {
    return &UserController{
        UserInteractor: interactor.UserInteractor{
            IUserRepository: repository.UserRepository{
                DB: db.Handler,
                Log: db.Logger,
            },
        },
        Log: db.Logger,
    }
}

func (uc *UserController) Create(ctx *gin.Context) {
    user, err := uc.UserInteractor.Create(ctx)
    if err != nil {
        ctx.JSON(500, gin.H{
            "message": err.Error(),
        })
        return
    }
    
    ctx.JSON(200, user.UserRespons())
}

func (uc *UserController) PING(ctx *gin.Context) {
    ctx.JSON(200, gin.H{
        "message": "PONG",
    })
}