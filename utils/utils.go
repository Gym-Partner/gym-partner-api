package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"reflect"
	"strings"
	"time"
	"unicode"

	"gitlab.com/gym-partner1/api/gym-partner-api/database"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
	"golang.org/x/crypto/bcrypt"
)

type IUtils[T model.User | model.Workout | model.UserToLogin | model.Follow] interface {
	HashPassword(password string) (string, *core.Error)
	InjectBodyInModel(ctx *gin.Context) (T, *core.Error)
	Bind(target, patch interface{}) *core.Error
	GenerateUUID() string
	SchemaToModel(workout database.MigrateWorkout, unitie database.MigrateUnitiesOfWorkout, exercices database.MigrateExercices, series database.MigrateSeries) model.Workout
}

type Utils[T model.User | model.Workout | model.UserToLogin | model.Follow] struct{}

func (u Utils[T]) GenerateUUID() string {
	return uuid.New().String()
}

func (u Utils[T]) HashPassword(password string) (string, *core.Error) {
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", core.NewError(500, "Error to hash password")
	}

	return string(hasedPassword), nil
}

func (u Utils[T]) InjectBodyInModel(ctx *gin.Context) (T, *core.Error) {
	var data T

	if err := ctx.ShouldBind(&data); err != nil {
		return data, core.NewError(500, err.Error())
	}

	return data, nil
}

func (u Utils[T]) SchemaToModel(workout database.MigrateWorkout, unities database.MigrateUnitiesOfWorkout, exercices database.MigrateExercices, series database.MigrateSeries) model.Workout {
	var newWorkout model.Workout
	var newUnities model.UnitiesOfWorkout

	for _, unity := range unities {
		var newUnity model.UnityOfWorkout
		newExercices := make(model.Exercices, 0)
		newSeries := make(model.Series, 0)

		for _, exercice := range exercices {
			if contains(unity.ExerciceId, exercice.Id) {
				newExercice := model.Exercice{
					Id:         exercice.Id,
					Name:       exercice.Name,
					Equipement: exercice.Equipement,
				}
				newExercices = append(newExercices, newExercice)
			}
		}

		for _, serie := range series {
			if contains(unity.SerieId, serie.Id) {
				newSerie := model.Serie{
					Id:          serie.Id,
					Weight:      serie.Weight,
					Repetitions: serie.Repetitions,
					IsWarmUp:    serie.IsWarmUp,
				}
				newSeries = append(newSeries, newSerie)
			}
		}

		newUnity = model.UnityOfWorkout{
			Id:          unity.Id,
			Exercices:   newExercices,
			Series:      newSeries,
			NbSerie:     unity.NbSerie,
			Comment:     unity.Comment,
			RestTimeSec: unity.RestTimeSec,
		}

		newUnities = append(newUnities, newUnity)

		newExercices = nil
		newSeries = nil
	}

	newWorkout = model.Workout{
		Id:               workout.Id,
		UserId:           workout.UserId,
		UnitiesOfWorkout: newUnities,
		Day:              workout.Day,
		Name:             workout.Name,
		Comment:          workout.Comment,
	}

	return newWorkout
}

func contains(slice []string, item string) bool {
	for _, id := range slice {
		if id == item {
			return true
		}
	}
	return false
}

func (u Utils[T]) Bind(target, patch interface{}) *core.Error {
	patchValue := reflect.ValueOf(patch)
	if patchValue.Kind() != reflect.Struct {
		return core.NewError(500, "The patch was be always an struct")
	}

	targetValue := reflect.ValueOf(target)
	if targetValue.Kind() != reflect.Ptr || targetValue.Elem().Kind() != reflect.Struct {
		return core.NewError(500, "The target was be always an pointer of struct")
	}

	for i := 0; i < patchValue.NumField(); i++ {
		patchField := patchValue.Type().Field(i)
		patchFieldName := toTitle(patchField.Name)
		targetField := targetValue.Elem().FieldByName(patchFieldName)

		if !targetField.IsValid() || !targetField.CanSet() {
			// Le champ n'existe pas dans la cible ou ne peut pas être modifié
			continue
		}

		patchFieldValue := patchValue.Field(i)

		if !isEmptyValue(patchFieldValue) {
			if targetField.Type() == reflect.TypeOf(time.Time{}) && patchFieldValue.Type() == reflect.TypeOf(int64(0)) {
				targetField.Set(reflect.ValueOf(time.Unix(patchFieldValue.Interface().(int64), 0)))
			} else {
				targetField.Set(patchFieldValue)
			}
		}
	}

	return nil
}

func toTitle(s string) string {
	return strings.Join(strings.FieldsFunc(s, unicode.IsSpace), " ")
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}

	return false
}

func StructToReadCloser(data interface{}) (io.ReadCloser, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return io.NopCloser(bytes.NewReader(jsonData)), nil
}
