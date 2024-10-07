package model

type User struct {
    Id string `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	UserName string `json:"user_name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type Users []User