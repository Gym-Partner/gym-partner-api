package repository

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/database"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
	"gorm.io/gorm"
)

const (
	WORKOUTS_TABLE_NAME  = "workouts"
	UNITIES_TABLE_NAME   = "unities_of_workout"
	EXERCISES_TABLE_NAME = "exercises"
	SERIES_TABLE_NAME    = "series"
)

type WorkoutRepository struct {
	DB  *gorm.DB
	Log *core.Log
}

func MockWorkoutRepository(db *gorm.DB) *WorkoutRepository {
	log := core.NewLog("/Users/oscar/Documents/gym-partner-env", true)
	log.ChargeLog()

	return &WorkoutRepository{
		DB:  db,
		Log: log,
	}
}

// ------------------------------ IS EXIST-------------------------------

func (wr WorkoutRepository) IsExist(ctx *gin.Context) bool {
	//TODO implement me
	panic("implement me")
}

// -------------------------------- CREATE-------------------------------

func (wr WorkoutRepository) CreateWorkouts(data model.Workout) *core.Error {
	newData := data.ModelToDbSchema()

	if retour := wr.DB.Table(WORKOUTS_TABLE_NAME).Create(&newData); retour.Error != nil {
		wr.Log.Error(retour.Error.Error())
		return core.NewError(http.StatusInternalServerError, core.ErrDBCreateWorkout, retour.Error)
	}

	return nil
}

func (wr WorkoutRepository) CreateUnitiesOfWorkout(data model.UnityOfWorkout) *core.Error {
	newData := data.ModelToDbSchema()

	if retour := wr.DB.Table(UNITIES_TABLE_NAME).Create(&newData); retour.Error != nil {
		wr.Log.Error(retour.Error.Error())
		return core.NewError(http.StatusInternalServerError, core.ErrDBCreateUnityOfWorkout, retour.Error)
	}

	return nil
}

func (wr WorkoutRepository) CreateExercise(data model.Exercise) *core.Error {
	if retour := wr.DB.Table(EXERCISES_TABLE_NAME).Create(&data); retour.Error != nil {
		wr.Log.Error(retour.Error.Error())
		return core.NewError(http.StatusInternalServerError, core.ErrDBCreateExercice, retour.Error)
	}

	return nil
}

func (wr WorkoutRepository) CreateSeries(data model.Serie) *core.Error {
	if retour := wr.DB.Table(SERIES_TABLE_NAME).Create(&data); retour.Error != nil {
		wr.Log.Error(retour.Error.Error())
		return core.NewError(http.StatusInternalServerError, core.ErrDBCreateSerie, retour.Error)
	}

	return nil
}

// ------------------------------ READ ONE ------------------------------

func (wr WorkoutRepository) GetOneWorkoutsByUserId(uid string) (database.MigrateWorkout, *core.Error) {
	var workout database.MigrateWorkout

	if retour := wr.DB.Table(WORKOUTS_TABLE_NAME).Where("user_id = ?", uid).Find(&workout); retour.Error != nil {
		wr.Log.Error(retour.Error.Error())
		return database.MigrateWorkout{}, core.NewError(http.StatusInternalServerError, core.ErrDBGetWorkout, retour.Error)
	}

	return workout, nil
}

func (wr WorkoutRepository) GetUnitiesById(id string) (database.MigrateUnityOfWorkout, *core.Error) {
	var unity database.MigrateUnityOfWorkout

	if retour := wr.DB.Table(UNITIES_TABLE_NAME).Where("id = ?", id).Find(&unity); retour.Error != nil {
		wr.Log.Error(retour.Error.Error())
		return database.MigrateUnityOfWorkout{}, core.NewError(http.StatusInternalServerError, core.ErrDBGetUnityOfWorkout, retour.Error)
	}

	return unity, nil
}

func (wr WorkoutRepository) GetExerciseById(id string) (database.MigrateExercise, *core.Error) {
	var exercice database.MigrateExercise

	if retour := wr.DB.Table(EXERCISES_TABLE_NAME).Where("id = ?", id).Find(&exercice); retour.Error != nil {
		wr.Log.Error(retour.Error.Error())
		return database.MigrateExercise{}, core.NewError(http.StatusInternalServerError, core.ErrDBGetExercice, retour.Error)
	}

	return exercice, nil
}

func (wr WorkoutRepository) GetSeriesById(id string) (database.MigrateSerie, *core.Error) {
	var serie database.MigrateSerie

	if retour := wr.DB.Table(SERIES_TABLE_NAME).Where("id = ?", id).Find(&serie); retour.Error != nil {
		wr.Log.Error(retour.Error.Error())
		return database.MigrateSerie{}, core.NewError(http.StatusInternalServerError, core.ErrDBGetSerie, retour.Error)
	}

	return serie, nil
}

// ------------------------------ READ ALL ------------------------------

func (wr WorkoutRepository) GetAllWorkoutsByUserId(uid string) (database.MigrateWorkouts, *core.Error) {
	var workouts database.MigrateWorkouts

	if retour := wr.DB.Table(WORKOUTS_TABLE_NAME).Where("user_id = ?", uid).Find(&workouts); retour.Error != nil {
		wr.Log.Error(core.ErrDBGetWorkouts, uid, retour.Error.Error())

		return workouts, core.NewError(
			http.StatusInternalServerError,
			core.ErrAppDBGetWorkouts,
			retour.Error)
	}

	return workouts, nil
}

// ------------------------------- UPDATE -------------------------------

func (wr WorkoutRepository) UpdateWorkouts(data model.Workout) *core.Error {
	//TODO implement me
	panic("implement me")
}

func (wr WorkoutRepository) UpdateUnitiesOfWorkout(data model.UnityOfWorkout) *core.Error {
	//TODO implement me
	panic("implement me")
}

func (wr WorkoutRepository) UpdateExercise(data model.Exercise) *core.Error {
	//TODO implement me
	panic("implement me")
}

func (wr WorkoutRepository) UpdateSeries(data model.Serie) *core.Error {
	//TODO implement me
	panic("implement me")
}

// ------------------------------- DELETE -------------------------------

func (wr WorkoutRepository) DeleteWorkoutsByUserId(uid string) *core.Error {
	//TODO implement me
	panic("implement me")
}
