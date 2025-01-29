package repository

import (
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gorm.io/gorm"
)

type AuthRepository struct {
	DB  *gorm.DB
	Log *core.Log
}
