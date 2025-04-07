package interactor

import (
	"reflect"

	"github.com/gin-gonic/gin"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/database"
	"gitlab.com/gym-partner1/api/gym-partner-api/mock"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
	"gitlab.com/gym-partner1/api/gym-partner-api/usecases/repository"
	"gitlab.com/gym-partner1/api/gym-partner-api/utils"
)

type IWorkoutInteractor interface {
	Create(ctx *gin.Context) *core.Error
	GetOneByUserId(ctx *gin.Context) (model.Workout, *core.Error)
	GetAllByUserId(ctx *gin.Context) (model.Workouts, *core.Error)
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
		data.ChargeData(*uid.(*string), data.Day)
	}

	if err := wi.IWorkoutRepository.CreateWorkout(data); err != nil {
		return err
	}

	for _, unity := range data.UnitiesOfWorkout {
		if err := wi.IWorkoutRepository.CreateUnityOfWorkout(unity); err != nil {
			return err
		}

		for _, exercice := range unity.Exercices {
			if err := wi.IWorkoutRepository.CreateExcercice(exercice); err != nil {
				return err
			}
		}

		for _, serie := range unity.Series {
			if err := wi.IWorkoutRepository.CreateSerie(serie); err != nil {
				return err
			}
		}
	}

	return nil
}

func (wi *WorkoutInteractor) GetOneByUserId(ctx *gin.Context) (model.Workout, *core.Error) {
	var emptyWorkout model.Workout
	uid, _ := ctx.Get("uid")

	workout, err := wi.IWorkoutRepository.GetOneWorkoutByUserId(*uid.(*string))
	if err != nil {
		return emptyWorkout, err
	}

	var unities database.MigrateUnitiesOfWorkout
	var exercices database.MigrateExercices
	var series database.MigrateSeries

	for _, unityId := range workout.UnitiesId {
		unity, err := wi.IWorkoutRepository.GetUntyById(unityId)
		if err != nil {
			return emptyWorkout, err
		}
		unities = append(unities, unity)

		for _, exerciceId := range unity.ExerciceId {
			exercice, err := wi.IWorkoutRepository.GetExerciceById(exerciceId)
			if err != nil {
				return emptyWorkout, err
			}
			exercices = append(exercices, exercice)
		}

		for _, serieId := range unity.SerieId {
			serie, err := wi.IWorkoutRepository.GetSerieById(serieId)
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
	//TODO implement me
	panic("implement me")
}
