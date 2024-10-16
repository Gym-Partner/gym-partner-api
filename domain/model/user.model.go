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

func (u *User) Respons() gin.H {
	return gin.H{
		"data": gin.H{
			"id": u.Id,
			"first_name": u.FirstName,
			"last_name": u.LastName,
			"username": u.UserName,
			"email": u.Email,
		},
	}
}

func (u *Users) Respons() gin.H {
	return gin.H{
		"data": u,
	}
}