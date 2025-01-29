package interactor

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/middleware"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
	"gitlab.com/gym-partner1/api/gym-partner-api/usecases/repository"
	"gitlab.com/gym-partner1/api/gym-partner-api/utils"
	"net/http"
	"time"
)

type IAuthInteractor interface {
	Authenticate(ctx *gin.Context) (model.Auth, *core.Error)
}

type AuthInteractor struct {
	IAuthRepository repository.IAuthRepository
	IUtils          utils.IUtils[model.UserToLogin]
	IAuthMiddleware middleware.IAuthMiddleware
}

func (ai AuthInteractor) Authenticate(ctx *gin.Context) (model.Auth, *core.Error) {
	var expirationTime time.Time

	data, err := ai.IUtils.InjectBodyInModel(ctx)
	if err != nil {
		return model.Auth{}, err
	}

	userId, err := ai.IAuthRepository.GetUserIDByEmail(data.Email)
	if err != nil {
		return model.Auth{}, err
	}

	token, newErr := ai.IAuthMiddleware.GenerateJWT(userId, "TOKEN_SECRET", 7*24*time.Hour, &expirationTime)
	if newErr != nil {
		return model.Auth{}, core.NewError(http.StatusUnauthorized, newErr.Error())
	}

	refreshToken, otherErr := ai.IAuthMiddleware.GenerateJWT(userId, "REFRESH_TOKEN_SECRET", 17*24*time.Hour)
	if otherErr != nil {
		return model.Auth{}, core.NewError(http.StatusUnauthorized, otherErr.Error())
	}

	auth := model.Auth{
		Id:           ai.IUtils.GenerateUUID(),
		UserId:       userId,
		Token:        token,
		RefreshToken: refreshToken,
		ExpiresAt:    expirationTime,
	}

	err = ai.IAuthRepository.Create(auth)
	if err != nil {
		return model.Auth{}, err
	}

	return auth, nil
}
