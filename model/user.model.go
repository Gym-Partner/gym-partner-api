package model

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type User struct {
	Id        string    `json:"id" gorm:"primaryKey, not null" swaggerignore:"true"`
	FirstName string    `json:"first_name" example:"test"`
	LastName  string    `json:"last_name" example:"test"`
	UserName  string    `json:"username" gorm:"column:username; not null" example:"test_test"`
	Email     string    `json:"email" gorm:"not null" example:"test@test.com"`
	Password  string    `json:"password" gorm:"not null" example:"aaaAAA111"`
	Age       int       `json:"age" example:"24"`
	Followers []string  `json:"followers" gorm:"-"`
	Following []string  `json:"following" gorm:"-"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}
type Users []User

type NewUsers struct {
	Users Users
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserImage struct {
	Id       string `json:"id"`
	UserId   string `json:"user_id"`
	ImageURL string `json:"image_url"`
}

func (u *User) Respons() gin.H {
	return gin.H{
		"data": gin.H{
			"id":         u.Id,
			"first_name": u.FirstName,
			"last_name":  u.LastName,
			"username":   u.UserName,
			"email":      u.Email,
			"age":        u.Age,
			"followers":  u.Followers,
			"following":  u.Following,
		},
	}
}

func (u *User) GenerateUID() {
	u.Id = uuid.New().String()
}

func (u *User) GenerateTestStruct(uid ...string) {
	newUid := ""
	for _, v := range uid {
		newUid = v
	}

	if len(newUid) > 0 {
		u.Id = newUid
	} else {
		u.Id = uuid.New().String()
	}

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
	var result []gin.H

	for _, user := range *u {
		result = append(result, gin.H{
			"id":         user.Id,
			"first_name": user.FirstName,
			"last_name":  user.LastName,
			"username":   user.UserName,
			"email":      user.Email,
			"age":        user.Age,
			"followers":  user.Followers,
			"following":  user.Following,
		})
	}

	return gin.H{
		"data": result,
	}
}

func (u *Users) GenerateTestStruct() *NewUsers {
	*u = Users{
		{
			Id:        uuid.New().String(),
			FirstName: "Test1",
			LastName:  "Test1",
			UserName:  "test_test1",
			Email:     "test1@test.com",
			// Password:  "aaaAAA111",
		},
		{
			Id:        uuid.New().String(),
			FirstName: "Test2",
			LastName:  "Test2",
			UserName:  "test_test2",
			Email:     "test2@test.com",
			// Password:  "aaaAAA222",
		},
		{
			Id:        uuid.New().String(),
			FirstName: "Test3",
			LastName:  "Test3",
			UserName:  "test_test3",
			Email:     "test3@test.com",
			// Password:  "aaaAAA333",
		},
	}

	return &NewUsers{
		Users: *u,
	}
}

func (u *NewUsers) AddCreatedAt() {
	for _, user := range u.Users {
		user.CreatedAt = time.Now()
	}
}

func (u *UserImage) Response() gin.H {
	return gin.H{
		"data": gin.H{
			"id":        u.Id,
			"user_id":   u.UserId,
			"image_url": u.ImageURL,
		},
	}
}
