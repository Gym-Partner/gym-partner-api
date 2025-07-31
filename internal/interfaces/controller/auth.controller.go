package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/interfaces/repository"
	"gitlab.com/gym-partner1/api/gym-partner-api/middleware"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
	"gitlab.com/gym-partner1/api/gym-partner-api/services/interactor"
	"gitlab.com/gym-partner1/api/gym-partner-api/utils"
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

// Login godoc
// @Summary Login user
// @Description Login user with his credentials (email / password) and retrieve token / refresh_token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body model.UserToLogin true "User's credentials for login"
// @Success 200 {object} model.Auth "User login successfully"
// @Failure 401 {object} core.Error "Token generation error"
// @Failure 500 {object} core.Error "Internal server error
// @Router /auth/sign_in [post]
func (ac *AuthController) Login(ctx *gin.Context) {
	auth, err := ac.IAuthInteractor.Authenticate(ctx)
	if err != nil {
		ctx.JSON(err.Code, err.Respons())
		return
	}
	ctx.JSON(http.StatusOK, auth.Response())
}

// RefreshToken godoc
// @Summary Re-generate token
// @Description Re-generate user's token with his refresh token. Use it when token is expired
// @Tags Auth
// @Accept multipart/form-data
// @Produce json
// @Param refresh_token formData string true "User's refresh token"
// @Success 200 {object} model.Auth "User re-login successfully"
// @Failure 401 {object} core.Error "Token generation error"
// @Failure 500 {object} core.Error "Internal server error
// @Router /auth/refresh_token [post]
func (ac *AuthController) RefreshToken(ctx *gin.Context) {
	newAuth, err := ac.IAuthInteractor.RefreshAuthenticate(ctx)
	if err != nil {
		ctx.JSON(err.Code, err.Respons())
		return
	}
	ctx.JSON(http.StatusOK, newAuth.Response())
}
