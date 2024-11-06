package model

import "time"

type UnityOfWorkout struct {
	Id          int       `json:"id"`
	ExerciceId  int       `json:"exercice_id"`
	SerieId     int       `json:"serie_id"`
	NbSerie     int       `json:"nb_serie"`
	Comment     string    `json:"comment"`
	RestTimeSec time.Time `json:"rest_time_sec"`
}

type UnitiesOfWorkout []UnityOfWorkout
