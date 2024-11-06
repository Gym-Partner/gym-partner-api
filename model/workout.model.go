package model

import "time"

type Workouts []Workout
type Workout struct {
	Id               string           `json:"id"`
	UserId           string           `json:"user_id"`
	UnitiesOfWorkout []UnityOfWorkout `json:"unities_of_workout"`
	Day              time.Time        `json:"day"`
	Name             string           `json:"name"`
	Comment          string           `json:"comment"`
}

type UnityOfWorkout struct {
	Id          string     `json:"id"`
	Exercices   []Exercice `json:"exercices"`
	Series      []Serie    `json:"series"`
	NbSerie     int        `json:"nb_serie"`
	Comment     string     `json:"comment"`
	RestTimeSec time.Time  `json:"rest_time_sec"`
}

type Serie struct {
	Id          string `json:"id"`
	Weight      int    `json:"weight"`
	Repetitions int    `json:"repetitions"`
	IsWarmUp    bool   `json:"is_warm_up"`
}

type Exercice struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Equipement bool   `json:"equipement"`
}
