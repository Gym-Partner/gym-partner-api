package migrate

import (
	"github.com/google/uuid"
	"time"
)

type UserMigrate struct {
	ID        uuid.UUID `gorm:"primaryKey;not null"`
	FirstName string
	LastName  string
	Username  string `gorm:"column:username;not null;unique"`
	Email     string `gorm:"not null;unique"`
	Phone     string
	Password  string
	Roles     []UserRoleMigrate `gorm:"foreignkey:UserID"`
	CreatedAt time.Time         `gorm:"autoCreateTime"`
}

type UserRoleMigrate struct {
	ID     uuid.UUID `gorm:"primaryKey;not null"`
	UserID uuid.UUID `gorm:"index;not null"`
	RoleID uuid.UUID `gorm:"index;not null"`
}

func (UserMigrate) TableName() string     { return "users" }
func (UserRoleMigrate) TableName() string { return "user_roles" }
