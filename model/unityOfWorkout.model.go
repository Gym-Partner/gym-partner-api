package model

import "time"

type UnityOfWorkout struct {
	Id          string    `json:"id" gorm:"primaryKey, not null"`
	ExerciceId  string    `json:"exercice_id" gorm:"not null"`
	SerieId     string    `json:"serie_id" gorm:"not null"`
	NbSerie     int       `json:"nb_serie" gorm:"not null"`
	Comment     string    `json:"comment"`
	RestTimeSec time.Time `json:"rest_time_sec"`
}

type UnitiesOfWorkout []UnityOfWorkout
