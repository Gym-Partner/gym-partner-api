package interactor

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/middleware"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
	"gitlab.com/gym-partner1/api/gym-partner-api/usecases/repository"
	"gitlab.com/gym-partner1/api/gym-partner-api/utils"
)

type IAuthInteractor interface {
	Authenticate(ctx *gin.Context) (model.Auth, *core.Error)
	RefreshAuthenticate(ctx *gin.Context) (model.Auth, *core.Error)
}

type AuthInteractor struct {
	IAuthRepository repository.IAuthRepository
	IUtils          utils.IUtils[model.UserToLogin]
	IAuthMiddleware middleware.IAuthMiddleware
}

func (ai AuthInteractor) Authenticate(ctx *gin.Context) (model.Auth, *core.Error) {
	data, err := ai.IUtils.InjectBodyInModel(ctx)
	if err != nil {
		return model.Auth{}, err
	}

	userId, err := ai.IAuthRepository.GetUserIDByEmail(data.Email)
	if err != nil {
		return model.Auth{}, err
	}

	token, refresh, expiration, err := ai.generateJWTAndRefresh(userId)
	if err != nil {
		return model.Auth{}, err
	}

	auth := model.Auth{
		Id:           ai.IUtils.GenerateUUID(),
		UserId:       userId,
		Token:        token,
		RefreshToken: refresh,
		ExpiresAt:    expiration,
	}

	err = ai.IAuthRepository.Create(auth)
	if err != nil {
		return model.Auth{}, err
	}

	return auth, nil
}

func (ai AuthInteractor) RefreshAuthenticate(ctx *gin.Context) (model.Auth, *core.Error) {
	refresh_token := ctx.PostForm("refresh_token")

	existingAuth, err := ai.IAuthRepository.GetAuthByRefreshToken(refresh_token)
	if err != nil {
		return model.Auth{}, err
	}

	if err := ai.IAuthMiddleware.VerifyRefreshToken(existingAuth, refresh_token); err != nil {
		return model.Auth{}, err
	}

	newToken, newRefresh, expiration, err := ai.generateJWTAndRefresh(existingAuth.UserId)
	if err != nil {
		return model.Auth{}, err
	}

	updatedAuth := model.Auth{
		Id:           existingAuth.Id,
		UserId:       existingAuth.UserId,
		Token:        newToken,
		RefreshToken: newRefresh,
		ExpiresAt:    expiration,
	}

	if err := ai.IAuthRepository.Delete(existingAuth.UserId); err != nil {
		return model.Auth{}, err
	}

	if err := ai.IAuthRepository.Create(updatedAuth); err != nil {
		return model.Auth{}, err
	}

	return updatedAuth, nil
}

func (ai AuthInteractor) generateJWTAndRefresh(userId string) (token, refresh string, expiration time.Time, err *core.Error) {
	var expirationTime time.Time
	token, newErr := ai.IAuthMiddleware.GenerateJWT(userId, "TOKEN_SECRET", 7*24*time.Hour, &expirationTime)
	if newErr != nil {
		return "", "", time.Now(), core.NewError(http.StatusUnauthorized, newErr.Error())
	}

	refresh, otherErr := ai.IAuthMiddleware.GenerateJWT(userId, "REFRESH_TOKEN_SECRET", 17*24*time.Hour)
	if otherErr != nil {
		return "", "", time.Now(), core.NewError(http.StatusUnauthorized, otherErr.Error())
	}

	return token, refresh, expirationTime, nil
}
