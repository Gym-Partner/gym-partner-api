package model

type Follow struct {
	Id        string   `json:"id"`
	UserId    string   `json:"user_id"`
	Followers []string `json:"followers"`
	Following []string `json:"following"`
}
type Follows []Follow
