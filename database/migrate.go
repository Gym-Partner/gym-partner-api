package database

import (
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
	"time"

	"github.com/lib/pq"
)

type MigrateUser struct {
	Id        string    `json:"id" gorm:"primaryKey, not null"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Username  string    `json:"username" gorm:"column:username; not null"`
	Email     string    `json:"email" gorm:"not null"`
	Password  string    `json:"password" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

type MigrateWorkout struct {
	Id        string         `json:"id" gorm:"primaryKey, not null"`
	UserId    string         `json:"user_id" gorm:"not null"`
	UnitiesId pq.StringArray `json:"unities_id" gorm:"type:text[]; not null"`
	Day       time.Time      `json:"day" gorm:"autoCreateTime"`
	Name      string         `json:"name" gorm:"not null"`
	Comment   string         `json:"comment"`
}
type MigrateWorkouts []MigrateWorkout

func (mw MigrateWorkout) SchemaToModel(unities MigrateUnitiesOfWorkout, exercices MigrateExercices, series MigrateSeries) model.Workout {
	var workout model.Workout
	var newUnities model.UnitiesOfWorkout
	var newExercices model.Exercices
	var newSeries model.Series

	workout.Id = mw.Id
	workout.UserId = mw.UserId
	workout.Day = mw.Day
	workout.Name = mw.Name
	workout.Comment = mw.Comment

	for _, unity := range unities {
		var newUnity model.UnityOfWorkout

		for _, exercice := range exercices {
			var newExercice model.Exercice

			newExercice.Id = exercice.Id
			newExercice.Name = exercice.Name
			newExercice.Equipement = exercice.Equipement

			newExercices = append(newExercices, newExercice)
		}

		for _, serie := range series {
			var newSerie model.Serie

			newSerie.Id = serie.Id
			newSerie.Weight = serie.Weight
			newSerie.Repetitions = serie.Repetitions
			newSerie.IsWarmUp = serie.IsWarmUp

			newSeries = append(newSeries, newSerie)
		}

		newUnity.Id = unity.Id
		newUnity.Exercices = newExercices
		newUnity.Series = newSeries
		newUnity.NbSerie = unity.NbSerie
		newUnity.Comment = unity.Comment
		newUnity.RestTimeSec = unity.RestTimeSec

		newUnities = append(newUnities, newUnity)
	}

	return workout
}

type MigrateUnityOfWorkout struct {
	Id          string         `json:"id" gorm:"primaryKey, not null"`
	ExerciceId  pq.StringArray `json:"exercice_id" gorm:"type:text[]; not null"`
	SerieId     pq.StringArray `json:"serie_id" gorm:"type:text[]; not null"`
	NbSerie     int            `json:"nb_serie" gorm:"not null"`
	Comment     string         `json:"comment"`
	RestTimeSec time.Time      `json:"rest_time_sec"`
}
type MigrateUnitiesOfWorkout []MigrateUnityOfWorkout

type MigrateSerie struct {
	Id          string `json:"id" gorm:"primaryKey, not null"`
	Weight      int    `json:"weight" gorm:"not null"`
	Repetitions int    `json:"repitions" gorm:"not null"`
	IsWarmUp    bool   `json:"is_warm_up" gorm:"not null"`
}
type MigrateSeries []MigrateSerie

type MigrateExercice struct {
	Id         string `json:"id" gorm:"primaryKey, not null"`
	Name       string `json:"name" gorm:"not null"`
	Equipement bool   `json:"equipement" gorm:"not null"`
}
type MigrateExercices []MigrateExercice

func (MigrateUser) TableName() string {
	return "user"
}

func (MigrateWorkout) TableName() string {
	return "workout"
}

func (MigrateUnityOfWorkout) TableName() string {
	return "unity_of_workout"
}

func (MigrateSerie) TableName() string {
	return "serie"
}

func (MigrateExercice) TableName() string {
	return "exercice"
}
