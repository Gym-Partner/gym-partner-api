package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/interfaces/repository"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
	"gitlab.com/gym-partner1/api/gym-partner-api/usecases/interactor"
	"gitlab.com/gym-partner1/api/gym-partner-api/utils"
)

type UserController struct {
	UserInteractor interactor.UserInteractor
	Log            *core.Log
}

func NewUserController(db *core.Database) *UserController {
	cognito := core.NewCognito(db.Logger)

	return &UserController{
		UserInteractor: interactor.UserInteractor{
			IUserRepository: repository.UserRepository{
				DB:  db.Handler,
				Log: db.Logger,
			},
			IUtils:   utils.Utils[model.User]{},
			ICognito: cognito,
		},
		Log: db.Logger,
	}
}

// ------------------------------ CRUD ------------------------------

// @BasePath /api/v1

// Create godoc
// @Summary Create a new user
// @Schemes
// @Description Create new user in database and return the created user withour the password
// @Tags User
// @Accept json
// @Produce application/json
// @Param user body model.User{} true "User data"
// @Success 201 {object} model.User{} "User successfully created"
// @Failure 500 {object} core.Error{} "Internal server error"
// @Router /user/create [post]
func (uc *UserController) Create(ctx *gin.Context) {
	user, err := uc.UserInteractor.Create(ctx)
	if err != nil {
		ctx.JSON(err.Code, err.Respons())
		return
	}

	ctx.JSON(http.StatusCreated, user.Respons())
}

// GetAll godoc
// @Summary Retrieve all user
// @Schemes
// @Description Retreive all user in database and return this without password
// @Tags User
// @Produce application/json
// @Param Authorization header string true "User Token"
// @Success 200 {object} model.Users{} "Users successfully retrieves"
// @Failure 500 {object} core.Error{} "Internal server error"
// @Router /user/getAll [get]
func (uc *UserController) GetAll(ctx *gin.Context) {
	users, err := uc.UserInteractor.GetAll()
	if err != nil {
		ctx.JSON(err.Code, err.Respons())
		return
	}

	ctx.JSON(http.StatusOK, users.Respons())
}

// GetOne godoc
// @Summary Retrieve one user
// @Schemes
// @Description Retrieve one user with id in token and return this without password
// @Tags User
// @Produce application/json
// @Param Authorization header string true "User Token"
// @Success 200 {object} model.User{} "User successfully retrieve"
// @Failure 500 {object} core.Error{} "Internal server error"
// @Router /user/getOne [get]
func (uc *UserController) GetOne(ctx *gin.Context) {
	user, err := uc.UserInteractor.GetOne(ctx)
	if err != nil {
		ctx.JSON(err.Code, err.Respons())
		return
	}

	ctx.JSON(http.StatusOK, user.Respons())
}

// Update godoc
// @Summary Update one user
// @Schemes
// @Description Update on user and return nil
// @Tags User
// @Param Authorization header string true "User Token"
// @Param user_update body model.User{} true "User data update"
// @Success 200 {object} nil "User successfully updated"
// @Failure 500 {object} core.Error{} "Internal server error"
// @Router /user/update [patch]
func (uc *UserController) Update(ctx *gin.Context) {
	if err := uc.UserInteractor.Update(ctx); err != nil {
		ctx.JSON(err.Code, err.Respons())
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

// Delete godoc
// @Summary Delete one user
// @Schemes
// @Description Delete one user and return nil
// @Tags User
// @Param Authorization header string true "User Token"
// @Success 200 {object} nil "User successfully deleted"
// @Failure 500 {object} core.Error{} "Internal server error"
// @Router /user/delete [delete]
func (uc *UserController) Delete(ctx *gin.Context) {
	if err := uc.UserInteractor.Delete(ctx); err != nil {
		ctx.JSON(err.Code, err.Respons())
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

// ------------------------------ AUTH ------------------------------

// Login godoc
// @Summary Sign in one user
// @Schemes
// @Description Sign in one user with this credentials and return user's token
// @Tags User
// @Produce application/json
// @Param credentials body model.Login{} true "User's credentials"
// @Success 200 {object} string "User's token"
// @Failure 500 {object} core.Error{} "Internal server error"
// @Router /user/login [post]
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

// PING godoc
// @Summary Do ping
// @Schemes
// @Description Do ping for test connection with the API
// @Tags PING
// @Produce application/json
// @Success 200 {string} json "PONG"
// @Router /ping [get]
func (uc *UserController) PING(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "PONG",
	})
}
