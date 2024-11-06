package database

import "time"

type MigrateUser struct {
	Id        string `gorm:"primaryKey, not null`
	FirstName string
	LastName  string
	Username  string    `gorm:"column:username; not null"`
	Email     string    `gorm:"not null"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

type MigrateWorkout struct {
	Id       string    `gorm:"primaryKey, not null"`
	UserId   string    `gorm:"not null"`
	UnitieId string    `gorm:"not null"`
	Day      time.Time `gorm:"autoCreateTime"`
	Name     string    `gorm:"not null"`
	Comment  string
}

type MigrateUnityOfWorkout struct {
	Id          string `gorm:"primaryKey, not null"`
	ExerciceId  string `gorm:"not null"`
	SerieId     string `gorm:"not null"`
	NbSerie     int    `gorm:"not null"`
	Comment     string
	RestTimeSec time.Time
}

type MigrateSerie struct {
	Id          string `gorm:"primaryKey, not null"`
	Weight      int    `gorm:"not null"`
	Repetitions int    `gorm:"not null"`
	IsWarmUp    bool   `gorm:"not null"`
}

type MigrateExercice struct {
	Id         string `gorm:"primaryKey, not null"`
	Name       string `gorm:"not null"`
	Equipement bool   `gorm:"not null"`
}

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
