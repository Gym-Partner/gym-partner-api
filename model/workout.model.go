package model

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Workouts []Workout
type Workout struct {
	Id               string           `json:"id"`
	UserId           string           `json:"user_id"`
	UnitiesOfWorkout []UnityOfWorkout `json:"unities_of_workout"`
	Day              time.Time        `json:"day"`
	Name             string           `json:"name"`
	Comment          string           `json:"comment"`
}

func (w *Workout) GenerateUID() {
	w.Id = uuid.New().String()
}

type UnityOfWorkout struct {
	Id          string     `json:"id"`
	Exercices   []Exercice `json:"exercices"`
	Series      []Serie    `json:"series"`
	NbSerie     int        `json:"nb_serie"`
	Comment     string     `json:"comment"`
	RestTimeSec time.Time  `json:"rest_time_sec"`
}

func (uow *UnityOfWorkout) GenerateUID() {
	uow.Id = uuid.New().String()
}

type Serie struct {
	Id          string `json:"id"`
	Weight      int    `json:"weight"`
	Repetitions int    `json:"repetitions"`
	IsWarmUp    bool   `json:"is_warm_up"`
}

func (s *Serie) GenerateUID() {
	s.Id = uuid.New().String()
}

type Exercice struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Equipement bool   `json:"equipement"`
}

func (e *Exercice) GenerateUID() {
	e.Id = uuid.New().String()
}

func (w *Workout) Respons() gin.H {
	return gin.H{
		"data": w,
	}
}

func (w *Workout) ChargeData(uid string) {
	w.GenerateUID()
	w.UserId = uid
	w.Day = time.Now()

	for i := range w.UnitiesOfWorkout {
		w.UnitiesOfWorkout[i].GenerateUID()

		for j := range w.UnitiesOfWorkout[i].Exercices {
			w.UnitiesOfWorkout[i].Exercices[j].GenerateUID()
		}

		for k := range w.UnitiesOfWorkout[i].Series {
			w.UnitiesOfWorkout[i].Series[k].GenerateUID()
		}
	}
}
