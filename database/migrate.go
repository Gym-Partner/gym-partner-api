package database

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type MigrateUser struct {
	Id        string    `json:"id" gorm:"primaryKey;not null"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Username  string    `json:"username" gorm:"column:username;not null"`
	Email     string    `json:"email" gorm:"not null"`
	Password  string    `json:"password" gorm:"not null"`
	Age       int       `json:"age"`
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

func (mw *MigrateWorkout) GenerateForTest(userId string) {
	mw.Id = uuid.New().String()
	mw.UserId = userId
	mw.UnitiesId = pq.StringArray{
		uuid.New().String(),
		uuid.New().String(),
	}
	// mw.Day = time.Now()
	mw.Name = "Workout name test"
	mw.Comment = "Workout comment test"
}

type MigrateUnityOfWorkout struct {
	Id          string         `json:"id" gorm:"primaryKey;not null"`
	ExerciceId  pq.StringArray `json:"exercice_id" gorm:"type:text[]; not null"`
	SerieId     pq.StringArray `json:"serie_id" gorm:"type:text[]; not null"`
	NbSerie     int            `json:"nb_serie" gorm:"not null"`
	Comment     string         `json:"comment"`
	RestTimeSec int            `json:"rest_time_sec"`
}
type MigrateUnitiesOfWorkout []MigrateUnityOfWorkout

func (mu *MigrateUnityOfWorkout) GenerateForTest(ids pq.StringArray) {
	for _, value := range ids {
		mu.Id = value
		mu.ExerciceId = pq.StringArray{
			uuid.New().String(),
			uuid.New().String(),
		}
		mu.SerieId = pq.StringArray{
			uuid.New().String(),
			uuid.New().String(),
		}
		mu.NbSerie = 0
		mu.Comment = "Unity of workout comment test"
		// mu.RestTimeSec = time.Now()
	}
}

func (mus *MigrateUnitiesOfWorkout) GenerateForTest(unity MigrateUnityOfWorkout) {
	*mus = append(*mus, unity)
}

type MigrateSerie struct {
	Id          string `json:"id" gorm:"primaryKey;not null"`
	Weight      int    `json:"weight" gorm:"not null"`
	Repetitions int    `json:"repetitions" gorm:"not null"`
	IsWarmUp    bool   `json:"is_warm_up" gorm:"not null"`
}
type MigrateSeries []MigrateSerie

func (ms *MigrateSerie) GenerateForTest(ids pq.StringArray) {
	for _, value := range ids {
		ms.Id = value
		ms.Weight = 1
		ms.Repetitions = 1
		ms.IsWarmUp = true
	}
}

func (mss *MigrateSeries) GenerateForTest(serie MigrateSerie) {
	*mss = append(*mss, serie)
}

type MigrateExercice struct {
	Id         string `json:"id" gorm:"primaryKey;not null"`
	Name       string `json:"name" gorm:"not null"`
	Equipement bool   `json:"equipement" gorm:"not null"`
}
type MigrateExercices []MigrateExercice

func (me *MigrateExercice) GenerateForTest(ids pq.StringArray) {
	for _, value := range ids {
		me.Id = value
		me.Name = "Exercice name test"
		me.Equipement = true
	}
}

func (mes *MigrateExercices) GenerateForTest(exercice MigrateExercice) {
	*mes = append(*mes, exercice)
}

type MigrateAuth struct {
	Id           string    `json:"id" gorm:"primaryKey;not null"`
	UserId       string    `json:"user_id" gorm:"not null"`
	Token        string    `json:"token" gorm:"not null"`
	RefreshToken string    `json:"refresh_token" gorm:"not null"`
	ExpiresAt    time.Time `json:"expires_at" gorm:"not null"`
}

type MigrateFollows struct {
	FollowerId string    `json:"follower_id" gorm:"primaryKey; not null"`
	FollowedId string    `json:"followed_id" gorm:"primaryKey; not null"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
}

func (MigrateUser) TableName() string { return "user" }

func (MigrateWorkout) TableName() string { return "workout" }

func (MigrateUnityOfWorkout) TableName() string { return "unity_of_workout" }

func (MigrateSerie) TableName() string { return "serie" }

func (MigrateExercice) TableName() string { return "exercice" }

func (MigrateAuth) TableName() string { return "auth" }

func (MigrateFollows) TableName() string { return "follows" }
