package interactor

import (
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
	"gitlab.com/gym-partner1/api/gym-partner-api/usecases/repository"
	"gitlab.com/gym-partner1/api/gym-partner-api/utils"
)

type IAuthInteractor interface {
}

type AuthInteractor struct {
	IAuthRepository repository.IAuthRepository
	IUtils          utils.IUtils[model.Auth]
}
