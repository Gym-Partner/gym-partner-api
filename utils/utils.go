package utils

import (
    "gitlab.com/gym-partner1/api/gym-partner-api/core"
    "golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, *core.Error) {
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", core.NewError(500, "Error to hash password", err)
	}

	return string(hasedPassword), nil
}