package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/interfaces/repository"
	"gitlab.com/gym-partner1/api/gym-partner-api/middleware"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
	"gitlab.com/gym-partner1/api/gym-partner-api/usecases/interactor"
	"gitlab.com/gym-partner1/api/gym-partner-api/utils"
	"net/http"
)

type AuthController struct {
	IAuthInteractor interactor.IAuthInteractor
	Log             *core.Log
}

func NewAuthController(db *core.Database) *AuthController {
	return &AuthController{
		IAuthInteractor: &interactor.AuthInteractor{
			IAuthRepository: repository.AuthRepository{
				DB:  db.Handler,
				Log: db.Logger,
			},
			IUtils:          utils.Utils[model.UserToLogin]{},
			IAuthMiddleware: middleware.AuthMiddleware{},
		},
		Log: db.Logger,
	}
}

func (ac *AuthController) Login(ctx *gin.Context) {
	auth, err := ac.IAuthInteractor.Authenticate(ctx)
	if err != nil {
		ctx.JSON(err.Code, err.Respons())
		return
	}

	ctx.JSON(http.StatusOK, auth.Response())
}

func (ac *AuthController) RefreshToken(ctx *gin.Context) {}
