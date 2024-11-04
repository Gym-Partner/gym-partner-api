package model

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type User struct {
	Id        string `json:"id" gorm:"index:idx_name,unique" swaggerignore:"true"`
	FirstName string `json:"first_name" example:"test"`
	LastName  string `json:"last_name" example:"test"`
	UserName  string `json:"username" gorm:"column:username" example:"test_test"`
	Email     string `json:"email" example:"test@test.com"`
	Password  string `json:"password" example:"aaaAAA111"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Users []User

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

func (u *User) GenerateTestStruct() {
	u.Id = uuid.New().String()
	u.FirstName = "Test"
	u.LastName = "Test"
	u.UserName = "test_test"
	u.Email = "test@gmail.com"
	u.Password = "aaaAAA111"
}

func (u *Users) Respons() gin.H {
	return gin.H{
		"data": u,
	}
}
