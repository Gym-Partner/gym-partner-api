package model

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
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

func (w *Workout) Respons() gin.H {
	return gin.H{
		"data": gin.H{
			"id":                 w.Id,
			"user_id":            w.UserId,
			"unities_of_workout": w.UnitiesOfWorkout.Response(),
			"day":                w.Day.Format("2006-01-02"),
			"name":               w.Name,
			"comment":            w.Comment,
		},
	}
}

func (w *Workouts) Respons() gin.H {
	var result []gin.H

	for _, workout := range *w {
		result = append(result, gin.H{
			"id":                 workout.Id,
			"user_id":            workout.UserId,
			"unities_of_workout": workout.UnitiesOfWorkout.Response(),
			"day":                workout.Day.Format("2006-01-02"),
			"name":               workout.Name,
			"comment":            workout.Comment,
		})
	}

	return gin.H{
		"data": result,
	}
}

func (w *Workout) GenerateUID() {
	w.Id = uuid.New().String()
}

func (w *Workout) ChargeData(uid string, day time.Time) {
	var WG sync.WaitGroup

	w.GenerateUID()
	w.UserId = uid
	w.Day = day

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

func (w *Workout) GenerateTestWorkout(uid ...string) {
	newUid := ""
	for _, v := range uid {
		newUid = v
	}

	if len(newUid) > 0 {
		w.UserId = newUid
	} else {
		w.UserId = "1234-5678-9123"
	}

	w.Id = uuid.New().String()
	//w.Day = time.Now()
	w.UnitiesOfWorkout = generateTestUnities(2)
	w.Name = "Name test"
	w.Comment = "Comment test"
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

func (uow *UnitiesOfWorkout) Response() []gin.H {
	response := make([]gin.H, len(*uow))

	for i, unity := range *uow {
		response[i] = gin.H{
			"id":            unity.Id,
			"exercices":     unity.Exercices.Response(),
			"series":        unity.Series.Response(),
			"nb_serie":      unity.NbSerie,
			"comment":       unity.Comment,
			"rest_time_sec": unity.RestTimeSec.Format("2006-01-02"),
		}
	}

	return response
}

func (uw *UnityOfWorkout) GenerateUID() {
	uw.Id = uuid.New().String()
}

func (uow *UnityOfWorkout) ModelToDbSchema() database.MigrateUnityOfWorkout {
	var exerciceId, serieId []string

	for _, exercice := range uow.Exercices {
		exerciceId = append(exerciceId, exercice.Id)
	}

	for _, serie := range uow.Series {
		serieId = append(serieId, serie.Id)
	}

	return database.MigrateUnityOfWorkout{
		Id:          uow.Id,
		ExerciceId:  exerciceId,
		SerieId:     serieId,
		NbSerie:     uow.NbSerie,
		Comment:     uow.Comment,
		RestTimeSec: uow.RestTimeSec,
	}
}

func (uow *UnityOfWorkout) GenerateTestUnity() {
	uow.Id = uuid.New().String()
	uow.Exercices = generateTestExercices(2)
	uow.Series = generateTestSeries(2)
	uow.NbSerie = 2
	uow.Comment = "Comment test unity of workout"
	uow.RestTimeSec = time.Now()
}

func generateTestUnities(iteration int) UnitiesOfWorkout {
	var unities UnitiesOfWorkout

	for i := 0; i < iteration; i++ {
		unity := UnityOfWorkout{
			Id:        uuid.New().String(),
			Exercices: generateTestExercices(iteration),
			Series:    generateTestSeries(iteration),
			NbSerie:   i,
			Comment:   fmt.Sprintf("Comment test: %d", i),
			//RestTimeSec: time.Now(),
		}

		unities = append(unities, unity)
	}

	return unities
}

// ------------------------------ SERIE ------------------------------

type Serie struct {
	Id          string `json:"id"`
	Weight      int    `json:"weight"`
	Repetitions int    `json:"repetitions"`
	IsWarmUp    bool   `json:"is_warm_up"`
}
type Series []Serie

func (s *Series) Response() []gin.H {
	response := make([]gin.H, len(*s))

	for i, serie := range *s {
		response[i] = gin.H{
			"id":          serie.Id,
			"weight":      serie.Weight,
			"repetitions": serie.Repetitions,
			"is_warm_up":  serie.IsWarmUp,
		}
	}

	return response
}

func (s *Serie) GenerateUID() {
	s.Id = uuid.New().String()
}

func (s *Serie) GenerateTest() {
	s.Id = uuid.New().String()
	s.Weight = 4
	s.Repetitions = 4
	s.IsWarmUp = true
}

func generateTestSeries(iteration int) Series {
	var series Series

	for i := 0; i < iteration; i++ {
		serie := Serie{
			Id:          uuid.New().String(),
			Weight:      i,
			Repetitions: i,
			IsWarmUp:    false,
		}

		series = append(series, serie)
	}

	return series
}

// ------------------------------ EXERCICE ------------------------------

type Exercice struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Equipement bool   `json:"equipement"`
}
type Exercices []Exercice

func (e *Exercices) Response() []gin.H {
	response := make([]gin.H, len(*e))

	for i, exercice := range *e {
		response[i] = gin.H{
			"id":         exercice.Id,
			"name":       exercice.Name,
			"equipement": exercice.Equipement,
		}
	}

	return response
}

func (e *Exercice) GenerateUID() {
	e.Id = uuid.New().String()
}

func (e *Exercice) GenerateTest() {
	e.Id = uuid.New().String()
	e.Name = "Name test exercice"
	e.Equipement = true
}

func generateTestExercices(iteration int) Exercices {
	var exercices Exercices

	for i := 0; i < iteration; i++ {
		exercice := Exercice{
			Id:         uuid.New().String(),
			Name:       fmt.Sprintf("Exercice %d", i),
			Equipement: true,
		}

		exercices = append(exercices, exercice)
	}

	return exercices
}
