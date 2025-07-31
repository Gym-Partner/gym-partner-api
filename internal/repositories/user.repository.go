package repositories

import (
	"github.com/Gym-Partner/api-common/status"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB      *gorm.DB
	Catalog *status.StatusCatalog
}

func NewUserRepository(db *gorm.DB, catalog *status.StatusCatalog) *UserRepository {
	return &UserRepository{
		DB:      db,
		Catalog: catalog,
	}
}
