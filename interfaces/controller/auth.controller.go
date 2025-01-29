package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/interfaces/repository"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
	"gitlab.com/gym-partner1/api/gym-partner-api/usecases/interactor"
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
			IUtils: utils.Utils[model.Auth]{},
		},
		Log: db.Logger,
	}
}

func (ac *AuthController) Login(ctx *gin.Context) {}
