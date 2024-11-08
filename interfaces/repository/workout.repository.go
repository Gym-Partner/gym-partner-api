package repository

import (
	"net/http"

	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/database"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
	"gorm.io/gorm"
)

type WorkoutRepository struct {
	DB  *gorm.DB
	Log *core.Log
}

// ------------------------------ CREATE------------------------------

func (wr WorkoutRepository) CreateWorkout(data model.Workout) *core.Error {
	newData := data.ModelToDbSchema()

	if retour := wr.DB.Table("workout").Create(&newData); retour.Error != nil {
		wr.Log.Error(retour.Error.Error())
		return core.NewError(http.StatusInternalServerError, core.ErrDBCreateWorkout, retour.Error)
	}

	return nil
}

func (wr WorkoutRepository) CreateUnityOfWorkout(data model.UnityOfWorkout) *core.Error {
	newData := data.ModelToDbSchema()

	if retour := wr.DB.Table("unity_of_workout").Create(&newData); retour.Error != nil {
		wr.Log.Error(retour.Error.Error())
		return core.NewError(http.StatusInternalServerError, core.ErrDBCreateUnityOfWorkout, retour.Error)
	}

	return nil
}

func (wr WorkoutRepository) CreateExcercice(data model.Exercice) *core.Error {
	if retour := wr.DB.Table("exercice").Create(&data); retour.Error != nil {
		wr.Log.Error(retour.Error.Error())
		return core.NewError(http.StatusInternalServerError, core.ErrDBCreateExercice, retour.Error)
	}

	return nil
}

func (wr WorkoutRepository) CreateSerie(data model.Serie) *core.Error {
	if retour := wr.DB.Table("serie").Create(&data); retour.Error != nil {
		wr.Log.Error(retour.Error.Error())
		return core.NewError(http.StatusInternalServerError, core.ErrDBCreateSerie, retour.Error)
	}

	return nil
}

// ------------------------------ READ ONE ------------------------------

func (wr WorkoutRepository) GetOneWorkoutByUserId(uid string) (database.MigrateWorkout, *core.Error) {
	panic("implement me")
}

func (wr WorkoutRepository) GetUntiesById(ids []string) (database.MigrateUnitiesOfWorkout, *core.Error) {
	panic("implement me")
}

func (wr WorkoutRepository) GetExercicesById(ids []string) (database.MigrateExercices, *core.Error) {
	panic("implement me")
}

func (wr WorkoutRepository) GetSeriesById(ids []string) (database.MigrateSeries, *core.Error) {
	panic("implement me")
}

// ------------------------------ READ ALL ------------------------------

func (wr WorkoutRepository) GetAllWorkoutByUserId(uid string) (database.MigrateWorkouts, *core.Error) {
	panic("implement me")
}
