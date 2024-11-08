package repository

import (
	"net/http"

	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
	"gorm.io/gorm"
)

type WorkoutRepository struct {
	DB  *gorm.DB
	Log *core.Log
}

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
	// fmt.Println("EXERCICE")
	// fmt.Println("#################################################################")
	// b, _ := json.MarshalIndent(data, "", " ")
	// fmt.Println(string(b))

	return nil
}

func (wr WorkoutRepository) CreateSerie(data model.Serie) *core.Error {
	// fmt.Println("SERIE")
	// fmt.Println("#################################################################")
	// b, _ := json.MarshalIndent(data, "", " ")
	// fmt.Println(string(b))

	return nil
}
