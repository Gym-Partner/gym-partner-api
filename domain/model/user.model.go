package model

import "go.mongodb.org/mongo-driver/bson"

type User struct {
    Id string `json:"id"`
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	UserName string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type Users []User

var UserProjection = bson.M{
	"id": 1,
	"firstname": 1,
	"lastname": 1,
	"username": 1,
	"email": 1,
	"password": 1,
}

func (u *User) NewUserFromData(data User) *User {
	return &User{
		Id: data.Id,
		FirstName: data.FirstName,
		LastName: data.LastName,
		UserName: data.UserName,
		Email: data.Email,
	}
}