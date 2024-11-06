package model

import "time"

type Workout struct {
	Id       string    `json:"id" gorm:"primaryKey, not null"`
	UserId   string    `json:"user_id" gorm:"not null"`
	UnitieId string    `json:"unitie_id" gorm:"not null"`
	Day      time.Time `json:"day"`
	Name     string    `json:"name" gorm:"not null"`
	Comment  string    `json:"comment"`
}

type Workouts []Workout
