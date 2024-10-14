package utils

import (
    "github.com/gin-gonic/gin"
    "gitlab.com/gym-partner1/api/gym-partner-api/core"
    "gitlab.com/gym-partner1/api/gym-partner-api/domain/model"
    "golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, *core.Error) {
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", core.NewError(500, "Error to hash password", err)
	}

	return string(hasedPassword), nil
}

func InjectBodyInModel[T model.User](ctx *gin.Context) (T, *core.Error) {
	var data T

	if err := ctx.ShouldBind(&data); err != nil {
		return data, core.NewError(500, "Error to inject Resquest Body to model", err)
	}

	return data, nil
}