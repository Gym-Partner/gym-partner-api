package controller

import (
    "github.com/gin-gonic/gin"
    "gitlab.com/gym-partner1/api/gym-partner-api/core"
    "gitlab.com/gym-partner1/api/gym-partner-api/interfaces/repository"
    "gitlab.com/gym-partner1/api/gym-partner-api/usecases/interactor"
    "net/http"
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

// ------------------------------ CRUD ------------------------------

func (uc *UserController) Create(ctx *gin.Context) {
    user, err := uc.UserInteractor.Create(ctx)
    if err != nil {
        ctx.JSON(err.Code, err.Respons())
        return
    }

    ctx.JSON(http.StatusCreated, user.Respons())
}

func (uc *UserController) GetAll(ctx *gin.Context) {
    users, err := uc.UserInteractor.GetAll()
    if err != nil {
        ctx.JSON(err.Code, err.Respons())
        return
    }

    ctx.JSON(http.StatusOK, users.Respons())
}

func (uc *UserController) GetOne(ctx *gin.Context) {
    user, err := uc.UserInteractor.GetOne(ctx)
    if err != nil {
        ctx.JSON(err.Code, err.Respons())
        return
    }

    ctx.JSON(http.StatusOK, user.Respons())
}

func (uc *UserController) Update(ctx *gin.Context) {
    if err := uc.UserInteractor.Update(ctx); err != nil {
        ctx.JSON(err.Code, err.Respons())
        return
    }

    ctx.JSON(http.StatusOK, nil)
}

func (uc *UserController) Delete(ctx *gin.Context) {
    if err := uc.UserInteractor.Delete(ctx); err != nil {
        ctx.JSON(err.Code, err.Respons())
        return
    }

    ctx.JSON(http.StatusOK, nil)
}

// ------------------------------ AUTH ------------------------------

func (uc *UserController) Login(ctx *gin.Context) {
    user, err := uc.UserInteractor.GetOneByEmail(ctx)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, err.Respons())
        return
    }

    token, err := uc.UserInteractor.Login(ctx, user)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, err.Respons())
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "token": token,
    })
}

// ------------------------------ PING ------------------------------

func (uc *UserController) PING(ctx *gin.Context) {
    ctx.JSON(200, gin.H{
        "message": "PONG",
    })
}