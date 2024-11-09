package model

import (
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.com/gym-partner1/api/gym-partner-api/database"
)

// ------------------------------ WORKOUT ------------------------------

type Workout struct {
	Id               string           `json:"id"`
	UserId           string           `json:"user_id"`
	UnitiesOfWorkout UnitiesOfWorkout `json:"unities_of_workout"`
	Day              time.Time        `json:"day"`
	Name             string           `json:"name"`
	Comment          string           `json:"comment"`
}
type Workouts []Workout

func (w *Workout) GenerateUID() {
	w.Id = uuid.New().String()
}

func (w *Workout) Respons() gin.H {
	return gin.H{
		"data": w,
	}
}

func (w *Workout) ChargeData(uid string) {
	var WG sync.WaitGroup

	w.GenerateUID()
	w.UserId = uid
	w.Day = time.Now()

	for i := range w.UnitiesOfWorkout {
		WG.Add(1)

		go func(i int) {
			defer WG.Done()
			w.UnitiesOfWorkout[i].GenerateUID()

			var exerciceWG sync.WaitGroup
			for j := range w.UnitiesOfWorkout[i].Exercices {
				exerciceWG.Add(1)

				go func(j int) {
					defer exerciceWG.Done()
					w.UnitiesOfWorkout[i].Exercices[j].GenerateUID()
				}(j)
			}
			exerciceWG.Wait()

			var serieWG sync.WaitGroup
			for k := range w.UnitiesOfWorkout[i].Series {
				serieWG.Add(1)

				go func(k int) {
					defer serieWG.Done()
					w.UnitiesOfWorkout[i].Series[k].GenerateUID()
				}(k)
			}
			serieWG.Wait()
		}(i)
	}

	WG.Wait()
}

func (w *Workout) ModelToDbSchema() database.MigrateWorkout {
	var unitiesIds []string

	for _, unity := range w.UnitiesOfWorkout {
		unitiesIds = append(unitiesIds, unity.Id)
	}

	return database.MigrateWorkout{
		Id:        w.Id,
		UserId:    w.UserId,
		UnitiesId: unitiesIds,
		Day:       w.Day,
		Name:      w.Name,
		Comment:   w.Comment,
	}
}

// ------------------------------ Unity Of Workout ------------------------------

type UnityOfWorkout struct {
	Id          string    `json:"id"`
	Exercices   Exercices `json:"exercices"`
	Series      Series    `json:"series"`
	NbSerie     int       `json:"nb_serie"`
	Comment     string    `json:"comment"`
	RestTimeSec time.Time `json:"rest_time_sec"`
}
type UnitiesOfWorkout []UnityOfWorkout

func (uow *UnityOfWorkout) GenerateUID() {
	uow.Id = uuid.New().String()
}

func (uow *UnityOfWorkout) ModelToDbSchema() database.MigrateUnityOfWorkout {
	var exerciceId, serieId []string

	exerciceChan := make(chan []string)
	serieChan := make(chan []string)

	go func() {
		var ids []string

		for _, exercice := range uow.Exercices {
			ids = append(ids, exercice.Id)
		}

		exerciceChan <- ids
		close(exerciceChan)
	}()

	go func() {
		var ids []string

		for _, serie := range uow.Series {
			ids = append(ids, serie.Id)
		}

		serieChan <- ids
		close(serieChan)
	}()

	exerciceId = <-exerciceChan
	serieId = <-serieChan

	return database.MigrateUnityOfWorkout{
		Id:          uow.Id,
		ExerciceId:  exerciceId,
		SerieId:     serieId,
		NbSerie:     uow.NbSerie,
		Comment:     uow.Comment,
		RestTimeSec: uow.RestTimeSec,
	}
}

// ------------------------------ SERIE ------------------------------

type Serie struct {
	Id          string `json:"id"`
	Weight      int    `json:"weight"`
	Repetitions int    `json:"repetitions"`
	IsWarmUp    bool   `json:"is_warm_up"`
}
type Series []Serie

func (s *Serie) GenerateUID() {
	s.Id = uuid.New().String()
}

// ------------------------------ EXERCICE ------------------------------

type Exercice struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Equipement bool   `json:"equipement"`
}
type Exercices []Exercice

func (e *Exercice) GenerateUID() {
	e.Id = uuid.New().String()
}
