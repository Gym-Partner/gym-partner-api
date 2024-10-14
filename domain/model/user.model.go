package model

import "github.com/gin-gonic/gin"

type User struct {
    Id string `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	UserName string `json:"username" gorm:"column:username"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type Users []User

func (u *User) UserRespons() gin.H {
	return gin.H{
		"data": u,
	}
}

func (u *Users) UsersRespons() gin.H {
	return gin.H{
		"data": u,
	}
}