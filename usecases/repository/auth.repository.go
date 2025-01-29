package repository

import (
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
)

type IAuthRepository interface {
	GetUserIDByEmail(email string) (string, *core.Error)
	Create(data model.Auth) *core.Error
	Delete(uid string) *core.Error
}
