package model

import "gitlab.com/gym-partner1/api/gym-partner-api/database"

type Follow struct {
	Id        string   `json:"id"`
	UserId    string   `json:"user_id"`
	Followers []string `json:"followers"`
	Following []string `json:"following"`
}
type Follows []Follow

func (f *Follow) ModelToSchema() database.MigrateFollow {
	return database.MigrateFollow{
		Id:        f.Id,
		UserId:    f.UserId,
		Followers: f.Followers,
		Following: f.Following,
	}
}
