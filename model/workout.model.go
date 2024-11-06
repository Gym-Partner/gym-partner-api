package model

import "time"

type Workout struct {
	Id       int       `json:"id"`
	UserId   string    `json:"user_id"`
	UnitieId int       `json:"unitie_id"`
	Day      time.Time `json:"day"`
	Name     string    `json:"name"`
	Comment  string    `json:"comment"`
}

type Workouts []Workout
