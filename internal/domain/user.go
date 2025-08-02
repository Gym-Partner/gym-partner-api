package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID  `json:"id" swaggerignore:"true"`
	FirstName string     `json:"first_name" example:"Titouan"`
	LastName  string     `json:"last_name" example:"Escorneboueu"`
	Username  string     `json:"username" example:"don_oscar_anton"`
	Email     string     `json:"email" example:"titouan.esc@icloud.com"`
	Phone     string     `json:"phone" example:"+33672135172"`
	Password  string     `json:"password" example:"aaaAAA111"`
	Followers []string   `json:"followers,omitempty" gorm:"-"`
	Following []string   `json:"following,omitempty" gorm:"-"`
	Image     string     `json:"image,omitempty" gorm:"-"`
	UserRoles []UserRole `json:"user_roles" gorm:"-"`
	CreatedAt time.Time  `json:"created_at"`
}

type Users []User

type UserRole struct {
	ID     uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"user_id"`
	RoleID uuid.UUID `json:"role_id"`
}

func (u *User) GenerateID() {
	u.ID = uuid.New()
}

func (ur *UserRole) GenerateID() {
	ur.ID = uuid.New()
}

func (u *User) Response() gin.H {
	data := gin.H{
		"id":         u.ID.String(),
		"first_name": u.FirstName,
		"last_name":  u.LastName,
		"username":   u.Username,
		"email":      u.Email,
		"phone":      u.Phone,
		"followers":  u.Followers,
		"following":  u.Following,
		"image":      u.Image,
		"user_roles": u.UserRoles,
		"created_at": u.CreatedAt,
	}

	return gin.H{
		"data": data,
	}
}

func (u *Users) Response() gin.H {
	var result []gin.H

	for _, user := range *u {
		result = append(result, user.Response())
	}

	return gin.H{
		"data": result,
	}
}
