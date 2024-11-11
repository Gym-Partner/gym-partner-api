package repository

import (
	"net/http"

	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/database"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
	"gorm.io/gorm"
)

const (
	WORKOUT_TABLE_NAME  = "workout"
	UNITY_TABLE_NAME    = "unity_of_workout"
	EXERCICE_TABLE_NAME = "exercice"
	SERIE_TABLE_NAME    = "serie"
)

type WorkoutRepository struct {
	DB  *gorm.DB
	Log *core.Log
}

// ------------------------------ CREATE------------------------------

func (wr WorkoutRepository) CreateWorkout(data model.Workout) *core.Error {
	newData := data.ModelToDbSchema()

	if retour := wr.DB.Table(WORKOUT_TABLE_NAME).Create(&newData); retour.Error != nil {
		wr.Log.Error(retour.Error.Error())
		return core.NewError(http.StatusInternalServerError, core.ErrDBCreateWorkout, retour.Error)
	}

	return nil
}

func (wr WorkoutRepository) CreateUnityOfWorkout(data model.UnityOfWorkout) *core.Error {
	newData := data.ModelToDbSchema()

	if retour := wr.DB.Table(UNITY_TABLE_NAME).Create(&newData); retour.Error != nil {
		wr.Log.Error(retour.Error.Error())
		return core.NewError(http.StatusInternalServerError, core.ErrDBCreateUnityOfWorkout, retour.Error)
	}

	return nil
}

func (wr WorkoutRepository) CreateExcercice(data model.Exercice) *core.Error {
	if retour := wr.DB.Table(EXERCICE_TABLE_NAME).Create(&data); retour.Error != nil {
		wr.Log.Error(retour.Error.Error())
		return core.NewError(http.StatusInternalServerError, core.ErrDBCreateExercice, retour.Error)
	}

	return nil
}

func (wr WorkoutRepository) CreateSerie(data model.Serie) *core.Error {
	if retour := wr.DB.Table(SERIE_TABLE_NAME).Create(&data); retour.Error != nil {
		wr.Log.Error(retour.Error.Error())
		return core.NewError(http.StatusInternalServerError, core.ErrDBCreateSerie, retour.Error)
	}

	return nil
}

// ------------------------------ READ ONE ------------------------------

func (wr WorkoutRepository) GetOneWorkoutByUserId(uid string) (database.MigrateWorkout, *core.Error) {
	var workout database.MigrateWorkout

	if retour := wr.DB.Table(WORKOUT_TABLE_NAME).Where("user_id = ?", uid).First(&workout); retour.Error != nil {
		wr.Log.Error(retour.Error.Error())
		return database.MigrateWorkout{}, core.NewError(http.StatusInternalServerError, core.ErrDBGetWorkout, retour.Error)
	}

	return workout, nil
}

func (wr WorkoutRepository) GetUntyById(id string) (database.MigrateUnityOfWorkout, *core.Error) {
	var unity database.MigrateUnityOfWorkout

	if retour := wr.DB.Table(UNITY_TABLE_NAME).Where("id = ?", id).First(&unity); retour.Error != nil {
		wr.Log.Error(retour.Error.Error())
		return database.MigrateUnityOfWorkout{}, core.NewError(http.StatusInternalServerError, core.ErrDBGetUnityOfWorkout, retour.Error)
	}

	return unity, nil
}

func (wr WorkoutRepository) GetExerciceById(id string) (database.MigrateExercice, *core.Error) {
	var exercice database.MigrateExercice

	if retour := wr.DB.Table(EXERCICE_TABLE_NAME).Where("id = ?", id).First(&exercice); retour.Error != nil {
		wr.Log.Error(retour.Error.Error())
		return database.MigrateExercice{}, core.NewError(http.StatusInternalServerError, core.ErrDBGetExercice, retour.Error)
	}

	return exercice, nil
}

func (wr WorkoutRepository) GetSerieById(id string) (database.MigrateSerie, *core.Error) {
	var serie database.MigrateSerie

	if retour := wr.DB.Table(SERIE_TABLE_NAME).Where("id = ?", id).First(&serie); retour.Error != nil {
		wr.Log.Error(retour.Error.Error())
		return database.MigrateSerie{}, core.NewError(http.StatusInternalServerError, core.ErrDBGetSerie, retour.Error)
	}

	return serie, nil
}

// ------------------------------ READ ALL ------------------------------

func (wr WorkoutRepository) GetAllWorkoutByUserId(uid string) (database.MigrateWorkouts, *core.Error) {
	//TODO implement me
	panic("implement me")
}
