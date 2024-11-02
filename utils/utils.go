package utils

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/domain/model"
	"golang.org/x/crypto/bcrypt"
	"io"
	"reflect"
	"strings"
	"time"
	"unicode"
)

type IUtils[T model.User] interface {
	HashPassword(password string) (string, *core.Error)
	InjectBodyInModel(ctx *gin.Context) (T, *core.Error)
	Bind(target, patch interface{}) *core.Error
}

type Utils[T model.User] struct{}

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
		return data, core.NewError(500, "Error to inject Resquest Body to model")
	}

	return data, nil
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
		patchFieldName := ToTitle(patchField.Name)
		targetField := targetValue.Elem().FieldByName(patchFieldName)

		if !targetField.IsValid() || !targetField.CanSet() {
			// Le champ n'existe pas dans la cible ou ne peut pas être modifié
			continue
		}

		patchFieldValue := patchValue.Field(i)

		if !IsEmptyValue(patchFieldValue) {
			if targetField.Type() == reflect.TypeOf(time.Time{}) && patchFieldValue.Type() == reflect.TypeOf(int64(0)) {
				targetField.Set(reflect.ValueOf(time.Unix(patchFieldValue.Interface().(int64), 0)))
			} else {
				targetField.Set(patchFieldValue)
			}
		}
	}

	return nil
}

func ToTitle(s string) string {
	return strings.Join(strings.FieldsFunc(s, unicode.IsSpace), " ")
}

func IsEmptyValue(v reflect.Value) bool {
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
