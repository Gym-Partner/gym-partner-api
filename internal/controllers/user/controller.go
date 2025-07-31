package user

import (
	"github.com/Gym-Partner/api-common/database"
	"github.com/Gym-Partner/api-common/status"
	"gitlab.com/gym-partner1/api/gym-partner-api/internal/repositories"
	"gitlab.com/gym-partner1/api/gym-partner-api/internal/services"
)

// Controller provides HTTP handlers for auth-related operations.
// It delegates business logic to the injected IService implementation.
type Controller struct {
	IService services.IUserService
}

func New(db *database.Database, catalog *status.StatusCatalog) *Controller {
	repo := repositories.NewUserRepository(db.Handler, catalog)
	svc := services.NewUserService(repo, catalog)
	return &Controller{
		IService: svc,
	}
}
