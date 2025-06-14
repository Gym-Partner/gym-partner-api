package repository

import (
	"fmt"
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

	if raw := wr.DB.
		Table(WORKOUTS_TABLE_NAME).
		Create(&newData); raw.Error != nil {
		wr.Log.Error(raw.Error.Error())
		return core.NewError(http.StatusInternalServerError, core.ErrDBCreateWorkout, raw.Error)
	}

	return nil
}

func (wr WorkoutRepository) CreateUnitiesOfWorkout(data model.UnityOfWorkout) *core.Error {
	newData := data.ModelToDbSchema()

	if raw := wr.DB.
		Table(UNITIES_TABLE_NAME).
		Create(&newData); raw.Error != nil {
		wr.Log.Error(raw.Error.Error())
		return core.NewError(http.StatusInternalServerError, core.ErrDBCreateUnityOfWorkout, raw.Error)
	}

	return nil
}

func (wr WorkoutRepository) CreateExercise(data model.Exercise) *core.Error {
	if raw := wr.DB.
		Table(EXERCISES_TABLE_NAME).
		Create(&data); raw.Error != nil {
		wr.Log.Error(raw.Error.Error())
		return core.NewError(http.StatusInternalServerError, core.ErrDBCreateExercise, raw.Error)
	}

	return nil
}

func (wr WorkoutRepository) CreateSeries(data model.Serie) *core.Error {
	if raw := wr.DB.
		Table(SERIES_TABLE_NAME).
		Create(&data); raw.Error != nil {
		wr.Log.Error(raw.Error.Error())
		return core.NewError(http.StatusInternalServerError, core.ErrDBCreateSeries, raw.Error)
	}

	return nil
}

// ------------------------------ READ ONE ------------------------------

func (wr WorkoutRepository) GetOneWorkoutsByUserId(uid string) (database.MigrateWorkout, *core.Error) {
	var workout database.MigrateWorkout

	if raw := wr.DB.
		Table(WORKOUTS_TABLE_NAME).
		Where("user_id = ?", uid).
		First(&workout); raw.Error != nil {
		wr.Log.Error(raw.Error.Error())
		return database.MigrateWorkout{}, core.NewError(http.StatusInternalServerError, core.ErrDBGetWorkout, raw.Error)
	}

	return workout, nil
}

func (wr WorkoutRepository) GetUnitiesById(id string) (database.MigrateUnityOfWorkout, *core.Error) {
	var unity database.MigrateUnityOfWorkout

	if raw := wr.DB.
		Table(UNITIES_TABLE_NAME).
		Where("id = ?", id).
		First(&unity); raw.Error != nil {
		wr.Log.Error(raw.Error.Error())
		return database.MigrateUnityOfWorkout{}, core.NewError(http.StatusInternalServerError, core.ErrDBGetUnityOfWorkout, raw.Error)
	}

	return unity, nil
}

func (wr WorkoutRepository) GetExerciseById(id string) (database.MigrateExercise, *core.Error) {
	var exercise database.MigrateExercise

	if raw := wr.DB.
		Table(EXERCISES_TABLE_NAME).
		Where("id = ?", id).
		First(&exercise); raw.Error != nil {
		wr.Log.Error(raw.Error.Error())
		return database.MigrateExercise{}, core.NewError(http.StatusInternalServerError, core.ErrDBGetExercise, raw.Error)
	}

	return exercise, nil
}

func (wr WorkoutRepository) GetSeriesById(id string) (database.MigrateSerie, *core.Error) {
	var serie database.MigrateSerie

	if raw := wr.DB.
		Table(SERIES_TABLE_NAME).
		Where("id = ?", id).
		First(&serie); raw.Error != nil {
		wr.Log.Error(raw.Error.Error())
		return database.MigrateSerie{}, core.NewError(http.StatusInternalServerError, core.ErrDBGetSeries, raw.Error)
	}

	return serie, nil
}

// ------------------------------ READ ALL ------------------------------

func (wr WorkoutRepository) GetAllWorkoutsByUserId(uid string) (database.MigrateWorkouts, *core.Error) {
	var workouts database.MigrateWorkouts

	if raw := wr.DB.
		Table(WORKOUTS_TABLE_NAME).
		Where("user_id = ?", uid).
		First(&workouts); raw.Error != nil {
		wr.Log.Error(core.ErrDBGetWorkouts, uid, raw.Error.Error())

		return workouts, core.NewError(
			http.StatusInternalServerError,
			core.ErrAppDBGetWorkouts,
			raw.Error)
	}

	return workouts, nil
}

// ------------------------------- UPDATE -------------------------------

func (wr WorkoutRepository) UpdateWorkouts(data model.Workout) *core.Error {
	if raw := wr.DB.
		Table(WORKOUTS_TABLE_NAME).
		Save(&data); raw.Error != nil {
		wr.Log.Error(fmt.Sprintf(core.ErrDBUpdateWorkout, data.UserId, raw.Error.Error()))
		return core.NewError(
			http.StatusInternalServerError,
			core.ErrAppDBUpdateWorkouts,
			raw.Error)
	}

	return nil
}

func (wr WorkoutRepository) UpdateUnitiesOfWorkout(data model.UnityOfWorkout) *core.Error {
	if raw := wr.DB.
		Table(UNITIES_TABLE_NAME).
		Save(&data); raw.Error != nil {
		wr.Log.Error(fmt.Sprintf(core.ErrDBUpdateUnitiesOfWorkouts, raw.Error.Error()))
		return core.NewError(
			http.StatusInternalServerError,
			core.ErrAppDBUpdateUnitiesOfWorkouts,
			raw.Error)
	}

	return nil
}

func (wr WorkoutRepository) UpdateExercise(data model.Exercise) *core.Error {
	if raw := wr.DB.
		Table(EXERCISES_TABLE_NAME).
		Save(&data); raw.Error != nil {
		wr.Log.Error(fmt.Sprintf(core.ErrDBUpdateExercises, raw.Error.Error()))
		return core.NewError(
			http.StatusInternalServerError,
			core.ErrAppDBUpdateExercises,
			raw.Error)
	}

	return nil
}

func (wr WorkoutRepository) UpdateSeries(data model.Serie) *core.Error {
	if raw := wr.DB.
		Table(SERIES_TABLE_NAME).
		Save(&data); raw.Error != nil {
		wr.Log.Error(fmt.Sprintf(core.ErrDBUpdateSeries, raw.Error.Error()))
		return core.NewError(
			http.StatusInternalServerError,
			core.ErrAppDBUpdateSeries,
			raw.Error)
	}

	return nil
}

// ------------------------------- DELETE -------------------------------

func (wr WorkoutRepository) DeleteWorkoutsByUserId(uid string) *core.Error {
	//TODO implement me
	panic("implement me")
}

func findByID[T any](wr WorkoutRepository, tableName, columnName, value string) (T, error) {
	var result T
	err := wr.DB.Table(tableName).Where(fmt.Sprintf("%s = ?", columnName), value).First(&result).Error
	return result, err
}
