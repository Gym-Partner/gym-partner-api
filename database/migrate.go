package database

import (
	"github.com/lib/pq"
	"time"
)

type MigrateUser struct {
	Id        string    `json:"id" gorm:"primaryKey;not null"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Username  string    `json:"username" gorm:"column:username;not null"`
	Email     string    `json:"email" gorm:"not null"`
	Password  string    `json:"password" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}
type MigrateUsers []MigrateUser

type MigrateWorkout struct {
	Id        string         `json:"id" gorm:"primaryKey;not null"`
	UserId    string         `json:"user_id" gorm:"not null"`
	UnitiesId pq.StringArray `json:"unities_id" gorm:"type:text[];not null"`
	Day       time.Time      `json:"day" gorm:"autoCreateTime"`
	Name      string         `json:"name" gorm:"not null"`
	Comment   string         `json:"comment"`
}
type MigrateWorkouts []MigrateWorkout

type MigrateUnityOfWorkout struct {
	Id          string         `json:"id" gorm:"primaryKey;not null"`
	ExerciceId  pq.StringArray `json:"exercice_id" gorm:"type:text[]; not null"`
	SerieId     pq.StringArray `json:"serie_id" gorm:"type:text[]; not null"`
	NbSerie     int            `json:"nb_serie" gorm:"not null"`
	Comment     string         `json:"comment"`
	RestTimeSec time.Time      `json:"rest_time_sec"`
}
type MigrateUnitiesOfWorkout []MigrateUnityOfWorkout

type MigrateSerie struct {
	Id          string `json:"id" gorm:"primaryKey;not null"`
	Weight      int    `json:"weight" gorm:"not null"`
	Repetitions int    `json:"repitions" gorm:"not null"`
	IsWarmUp    bool   `json:"is_warm_up" gorm:"not null"`
}
type MigrateSeries []MigrateSerie

type MigrateExercice struct {
	Id         string `json:"id" gorm:"primaryKey;not null"`
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
