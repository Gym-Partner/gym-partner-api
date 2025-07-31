package services

import (
	"github.com/Gym-Partner/api-common/status"
	"gitlab.com/gym-partner1/api/gym-partner-api/internal/repositories"
)

type UserService struct {
	IRepository repositories.IUserRepository
	Catalog     *status.StatusCatalog
}

func NewUserService(repo repositories.IUserRepository, catalog *status.StatusCatalog) *UserService {
	return &UserService{
		IRepository: repo,
		Catalog:     catalog,
	}
}
