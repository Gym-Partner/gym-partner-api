package model

type User struct {
    Id string `json:"id"`
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	UserName string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type Users []User
