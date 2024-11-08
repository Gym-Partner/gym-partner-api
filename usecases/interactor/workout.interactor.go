package interactor

import (
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

func (wi *WorkoutInteractor) Create(ctx *gin.Context) (model.Workout, *core.Error) {
	var workout model.Workout
	uid, _ := ctx.Get("uid")

	data, err := wi.IUtils.InjectBodyInModel(ctx)
	if err != nil {
		return workout, err
	}
	data.ChargeData(*uid.(*string))

	if err := wi.IWorkoutRepository.CreateWorkout(data); err != nil {
		return workout, err
	}

	// First version
	for _, unity := range data.UnitiesOfWorkout {
		if err := wi.IWorkoutRepository.CreateUnityOfWorkout(unity); err != nil {
			return workout, err
		}

		for _, exercice := range unity.Exercices {
			if err := wi.IWorkoutRepository.CreateExcercice(exercice); err != nil {
				return workout, err
			}
		}

		for _, serie := range unity.Series {
			if err := wi.IWorkoutRepository.CreateSerie(serie); err != nil {
				return workout, err
			}
		}
	}

	// Second version
	// var wg sync.WaitGroup
	// errChan := make(chan *core.Error, len(data.UnitiesOfWorkout)*2)

	// for _, unity := range data.UnitiesOfWorkout {
	// 	if err := wi.IWorkoutRepository.CreateUnityOfWorkout(unity); err != nil {
	// 		return workout, err
	// 	}

	// 	for _, exercice := range unity.Exercices {
	// 		wg.Add(1)

	// 		go func(exercice model.Exercice) {
	// 			defer wg.Done()

	// 			if err := wi.IWorkoutRepository.CreateExcercice(exercice); err != nil {
	// 				errChan <- err
	// 			}
	// 		}(exercice)
	// 	}

	// 	for _, serie := range unity.Series {
	// 		wg.Add(1)

	// 		go func(serie model.Serie) {
	// 			defer wg.Done()

	// 			if err := wi.IWorkoutRepository.CreateSerie(serie); err != nil {
	// 				errChan <- err
	// 			}
	// 		}(serie)
	// 	}
	// }

	// wg.Wait()
	// close(errChan)

	// if len(errChan) > 0 {
	// 	return workout, <-errChan
	// }

	return data, nil
}
