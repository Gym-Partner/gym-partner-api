package model

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Users []User
type User struct {
	Id        string    `json:"id" gorm:"primaryKey, not null" swaggerignore:"true"`
	FirstName string    `json:"first_name" example:"test"`
	LastName  string    `json:"last_name" example:"test"`
	UserName  string    `json:"username" gorm:"column:username; not null" example:"test_test"`
	Email     string    `json:"email" gorm:"not null" example:"test@test.com"`
	Password  string    `json:"password" gorm:"not null" example:"aaaAAA111"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) Respons() gin.H {
	return gin.H{
		"data": gin.H{
			"id":         u.Id,
			"first_name": u.FirstName,
			"last_name":  u.LastName,
			"username":   u.UserName,
			"email":      u.Email,
		},
	}
}

func (u *User) GenerateUID() {
	u.Id = uuid.New().String()
}

func (u *User) GenerateTestStruct() {
	u.Id = uuid.New().String()
	u.FirstName = "Test"
	u.LastName = "Test"
	u.UserName = "test_test"
	u.Email = "test@gmail.com"
	u.Password = "aaaAAA111"
}

func (u *User) UserToAnother(data User) {
	u.Id = data.Id
	u.FirstName = data.FirstName
	u.LastName = data.LastName
	u.UserName = data.UserName
	u.Email = data.Email
	u.Password = data.Password
}

func (u *Users) Respons() gin.H {
	return gin.H{
		"data": u,
	}
}
