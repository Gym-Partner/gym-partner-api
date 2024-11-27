package controller

import (
	"gitlab.com/gym-partner1/api/gym-partner-api/mock"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/interfaces/repository"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
	"gitlab.com/gym-partner1/api/gym-partner-api/usecases/interactor"
	"gitlab.com/gym-partner1/api/gym-partner-api/utils"
)

type UserController struct {
	IUserInteractor interactor.IUserInteractor
	Log             *core.Log
}

// ------------------------------ Constructor ------------------------------

func NewUserController(db *core.Database) *UserController {
	cognito := core.NewCognito(db.Logger)

	return &UserController{
		IUserInteractor: &interactor.UserInteractor{
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

// ------------------------------ Mock Constructor ------------------------------

func MockUserController(UCMock *mock.UserControllerMock) *UserController {
	log := core.NewLog("/Users/oscar/Documents/gym-partner-env", true)
	log.ChargeLog()

	return &UserController{
		IUserInteractor: UCMock,
		Log:             log,
	}
}

// ------------------------------ CRUD ------------------------------

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
	user, err := uc.IUserInteractor.Create(ctx)
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
	users, err := uc.IUserInteractor.GetAll()
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
	user, err := uc.IUserInteractor.GetOne(ctx)
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
	if err := uc.IUserInteractor.Update(ctx); err != nil {
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
	if err := uc.IUserInteractor.Delete(ctx); err != nil {
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
	user, err := uc.IUserInteractor.GetOneByEmail(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Respons())
		return
	}

	token, err := uc.IUserInteractor.Login(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Respons())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
