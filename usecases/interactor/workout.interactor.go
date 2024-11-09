package interactor

import (
	"gitlab.com/gym-partner1/api/gym-partner-api/database"
	"sync"

	"github.com/gin-gonic/gin"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
	"gitlab.com/gym-partner1/api/gym-partner-api/usecases/repository"
	"gitlab.com/gym-partner1/api/gym-partner-api/utils"
)

type WorkoutInteractor struct {
	IWorkoutRepository repository.IWorkoutRepository
	IUtils             utils.IUtils[model.Workout]
}

func (wi *WorkoutInteractor) Create(ctx *gin.Context) *core.Error {
	var wg sync.WaitGroup
	uid, _ := ctx.Get("uid")

	data, err := wi.IUtils.InjectBodyInModel(ctx)
	if err != nil {
		return err
	}
	data.ChargeData(*uid.(*string))

	if err := wi.IWorkoutRepository.CreateWorkout(data); err != nil {
		return err
	}

	errChan := make(chan *core.Error, len(data.UnitiesOfWorkout)*2)

	for _, unity := range data.UnitiesOfWorkout {
		if err := wi.IWorkoutRepository.CreateUnityOfWorkout(unity); err != nil {
			return err
		}

		for _, exercice := range unity.Exercices {
			wg.Add(1)

			go func(exercice model.Exercice) {
				defer wg.Done()

				if err := wi.IWorkoutRepository.CreateExcercice(exercice); err != nil {
					errChan <- err
				}
			}(exercice)
		}

		for _, serie := range unity.Series {
			wg.Add(1)

			go func(serie model.Serie) {
				defer wg.Done()

				if err := wi.IWorkoutRepository.CreateSerie(serie); err != nil {
					errChan <- err
				}
			}(serie)
		}
	}

	wg.Wait()
	close(errChan)

	if len(errChan) > 0 {
		return <-errChan
	}

	return nil
}

func (wi *WorkoutInteractor) GetOneByUserId(ctx *gin.Context) (model.Workout, *core.Error) {
	var emptyWorkout model.Workout
	var WG sync.WaitGroup
	uid, _ := ctx.Get("uid")

	workout, err := wi.IWorkoutRepository.GetOneWorkoutByUserId(*uid.(*string))
	if err != nil {
		return emptyWorkout, err
	}

	errChan := make(chan *core.Error, len(workout.UnitiesId)*2)
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
			WG.Add(1)

			go func(exerciceId string) {
				defer WG.Done()

				exercice, err := wi.IWorkoutRepository.GetExerciceById(exerciceId)
				if err != nil {
					errChan <- err
				}
				exercices = append(exercices, exercice)
			}(exerciceId)
		}

		for _, serieId := range unity.SerieId {
			WG.Add(1)

			go func(serieId string) {
				defer WG.Done()

				serie, err := wi.IWorkoutRepository.GetSerieById(serieId)
				if err != nil {
					errChan <- err
				}
				series = append(series, serie)
			}(serieId)
		}
	}
	WG.Wait()
	close(errChan)

	if len(errChan) > 0 {
		return emptyWorkout, <-errChan
	}

	newData := workout.SchemaToModel(unities, exercices, series)
	return newData, nil
}
