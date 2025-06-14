package interactor

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/database"
	"gitlab.com/gym-partner1/api/gym-partner-api/mock"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
	"gitlab.com/gym-partner1/api/gym-partner-api/services/repository"
	"gitlab.com/gym-partner1/api/gym-partner-api/utils"
)

type IWorkoutInteractor interface {
	Create(ctx *gin.Context) *core.Error
	GetOneByUserId(ctx *gin.Context) (model.Workout, *core.Error)
	GetAllByUserId(ctx *gin.Context) (model.Workouts, *core.Error)
	Update(ctx *gin.Context) *core.Error
	Delete(ctx *gin.Context) *core.Error
}

type WorkoutInteractor struct {
	IWorkoutRepository repository.IWorkoutRepository
	IUtils             utils.IUtils[model.Workout]
}

func MockWorkoutInteractor(workoutMock *mock.WorkoutInteractorMock, utilsMock *mock.UtilsMock[model.Workout]) *WorkoutInteractor {
	return &WorkoutInteractor{
		IWorkoutRepository: workoutMock,
		IUtils:             utilsMock,
	}
}

func (wi *WorkoutInteractor) Create(ctx *gin.Context) *core.Error {
	data, err := wi.IUtils.InjectBodyInModel(ctx)
	if err != nil {
		return err
	}

	if reflect.TypeOf(wi.IWorkoutRepository) != reflect.TypeOf(&mock.WorkoutInteractorMock{}) {
		uid, _ := ctx.Get("uid")
		data.ChargeData(uid.(string), data.Day)
	}

	if err := wi.IWorkoutRepository.CreateWorkouts(data); err != nil {
		return err
	}

	for _, unity := range data.UnitiesOfWorkout {
		if err := wi.IWorkoutRepository.CreateUnitiesOfWorkout(unity); err != nil {
			return err
		}

		for _, exercise := range unity.Exercises {
			if err := wi.IWorkoutRepository.CreateExercise(exercise); err != nil {
				return err
			}
		}

		for _, series := range unity.Series {
			if err := wi.IWorkoutRepository.CreateSeries(series); err != nil {
				return err
			}
		}
	}

	return nil
}

func (wi *WorkoutInteractor) GetOneByUserId(ctx *gin.Context) (model.Workout, *core.Error) {
	var emptyWorkout model.Workout
	uid, _ := ctx.Get("uid")

	workout, err := wi.IWorkoutRepository.GetOneWorkoutsByValue("user_id", uid.(string))
	if err != nil {
		return emptyWorkout, err
	}

	var unities database.MigrateUnitiesOfWorkout
	var exercices database.MigrateExercises
	var series database.MigrateSeries

	for _, unityId := range workout.UnitiesId {
		unity, err := wi.IWorkoutRepository.GetUnitiesById(unityId)
		if err != nil {
			return emptyWorkout, err
		}
		unities = append(unities, unity)

		for _, exerciceId := range unity.ExercisesId {
			exercice, err := wi.IWorkoutRepository.GetExerciseById(exerciceId)
			if err != nil {
				return emptyWorkout, err
			}
			exercices = append(exercices, exercice)
		}

		for _, serieId := range unity.SeriesId {
			serie, err := wi.IWorkoutRepository.GetSeriesById(serieId)
			if err != nil {
				return emptyWorkout, err
			}
			series = append(series, serie)
		}
	}

	newData := wi.IUtils.SchemaToModel(workout, unities, exercices, series)
	return newData, nil
}

func (wi *WorkoutInteractor) GetAllByUserId(ctx *gin.Context) (model.Workouts, *core.Error) {
	var emptyWorkouts model.Workouts
	uid, _ := ctx.Get("uid")

	workouts, err := wi.IWorkoutRepository.GetAllWorkoutsByUserId(uid.(string))
	if err != nil {
		return emptyWorkouts, err
	}

	var result model.Workouts

	for _, workout := range workouts {
		var unities database.MigrateUnitiesOfWorkout
		var exercices database.MigrateExercises
		var series database.MigrateSeries

		for _, unityId := range workout.UnitiesId {
			unity, err := wi.IWorkoutRepository.GetUnitiesById(unityId)
			if err != nil {
				return emptyWorkouts, err
			}
			unities = append(unities, unity)

			for _, exerciceId := range unity.ExercisesId {
				exercice, err := wi.IWorkoutRepository.GetExerciseById(exerciceId)
				if err != nil {
					return emptyWorkouts, err
				}
				exercices = append(exercices, exercice)
			}

			for _, serieId := range unity.SeriesId {
				serie, err := wi.IWorkoutRepository.GetSeriesById(serieId)
				if err != nil {
					return emptyWorkouts, err
				}
				series = append(series, serie)
			}
		}

		newData := wi.IUtils.SchemaToModel(workout, unities, exercices, series)
		result = append(result, newData)
	}

	return result, nil
}

func (wi *WorkoutInteractor) Update(ctx *gin.Context) *core.Error {
	uid, _ := ctx.Get("uid")
	update, err := wi.IUtils.InjectBodyInModel(ctx)
	if err != nil {
		return err
	}

	exist := wi.IWorkoutRepository.IsExist(update.Id)
	if !exist {
		return core.NewError(
			http.StatusNotFound,
			fmt.Sprintf(core.ErrAppINTWorkoutsNotExist, update.Name))
	}

	update.UserId = uid.(string)
	if err := wi.IWorkoutRepository.UpdateWorkouts(update); err != nil {
		return err
	}

	for _, unity := range update.UnitiesOfWorkout {
		if unity.Id == "" {
			unity.GenerateUID()
		}

		if err := wi.IWorkoutRepository.UpdateUnitiesOfWorkout(unity); err != nil {
			return err
		}

		for _, exercises := range unity.Exercises {
			if exercises.Id == "" {
				exercises.GenerateUID()
			}

			if err := wi.IWorkoutRepository.UpdateExercise(exercises); err != nil {
				return err
			}
		}

		for _, series := range unity.Series {
			if series.Id == "" {
				series.GenerateUID()
			}

			if err := wi.IWorkoutRepository.UpdateSeries(series); err != nil {
				return err
			}
		}
	}

	return nil
}

func (wi *WorkoutInteractor) Delete(ctx *gin.Context) *core.Error {
	workout, err := wi.IUtils.InjectBodyInModel(ctx)
	if err != nil {
		return err
	}

	if exist := wi.IWorkoutRepository.IsExist(workout.Id); !exist {
		return core.NewError(
			http.StatusNotFound,
			fmt.Sprintf(core.ErrAppINTWorkoutsNotExist, workout.Name),
		)
	}

	w, err := wi.IWorkoutRepository.GetOneWorkoutsByValue("id", workout.Id)
	if err != nil {
		return err
	}

	var unityIDs []string
	var exerciseIDs []string
	var seriesIDs []string

	for _, unityID := range w.UnitiesId {
		unity, err := wi.IWorkoutRepository.GetUnitiesById(unityID)
		if err != nil {
			return err
		}
		unityIDs = append(unityIDs, unity.Id)
		exerciseIDs = append(exerciseIDs, unity.ExercisesId...)
		seriesIDs = append(seriesIDs, unity.SeriesId...)
	}

	for _, exID := range exerciseIDs {
		if err := wi.IWorkoutRepository.DeleteExercises(exID); err != nil {
			return err
		}
	}

	for _, sID := range seriesIDs {
		if err := wi.IWorkoutRepository.DeleteSeries(sID); err != nil {
			return err
		}
	}

	for _, uID := range unityIDs {
		if err := wi.IWorkoutRepository.DeleteUnities(uID); err != nil {
			return err
		}
	}

	if err := wi.IWorkoutRepository.DeleteWorkouts(workout.Id); err != nil {
		return err
	}

	return nil
}
