package test

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/interfaces/repository"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"testing"
)

func setUpDB() (*gorm.DB, sqlmock.Sqlmock) {
	mockDB, mock, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDB,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})

	return db, mock
}

//func prepareTest(name string) []struct {
//	name        string
//	setupMock   func(mock sqlmock.Sqlmock)
//	expectedRes *core.Error
//} {
//	return []struct {
//		name        string
//		setupMock   func(mock sqlmock.Sqlmock)
//		expectedRes *core.Error
//	}{
//		{
//			name: name,
//		},
//	}
//}

func TestWorkoutRepository_CREATE_WORKOUT(t *testing.T) {
	var workout model.Workout
	workout.GenerateTestWorkout()
	db, mock := setUpDB()

	setupTest := []struct {
		name        string
		setupMock   func(mock sqlmock.Sqlmock)
		expectedRes *core.Error
	}{
		{
			name: core.TestREPWorkoutCreateSuccess,
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(`INSERT INTO "workout"`).
					WithArgs(
						workout.Id,
						workout.UserId,
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
						workout.Name,
						workout.Comment).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectedRes: (*core.Error)(nil),
		},
		{
			name: core.TestWorkoutCreateFailed,
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(`INSERT INTO "workout"`).
					WithArgs(
						workout.Id,
						workout.UserId,
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
						workout.Name,
						workout.Comment).
					WillReturnError(fmt.Errorf("database error"))
				mock.ExpectRollback()
			},
			expectedRes: core.NewError(http.StatusInternalServerError, core.ErrDBCreateWorkout, fmt.Errorf("database error")),
		},
	}

	for _, value := range setupTest {
		t.Run(value.name, func(t *testing.T) {
			wp := repository.MockWorkoutRepository(db)
			value.setupMock(mock)

			err := wp.CreateWorkout(workout)
			assert.Equal(t, err, value.expectedRes)
		})

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	}
}

func TestWorkoutRepository_CREATE_UNITY_OF_WORKOUT(t *testing.T) {
	var unity model.UnityOfWorkout
	unity.GenerateTestUnity()
	db, mock := setUpDB()

	setupTest := []struct {
		name        string
		setupMock   func(mock sqlmock.Sqlmock)
		expectedRes *core.Error
	}{
		{
			name: core.TestREPUnityOfWorkoutCreateSuccess,
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(`INSERT INTO "unity_of_workout"`).
					WithArgs(
						unity.Id,
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
						unity.NbSerie,
						unity.Comment,
						sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectedRes: (*core.Error)(nil),
		},
		{
			name: core.TestUnitiesOfWorkoutCreateFailed,
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(`INSERT INTO "unity_of_workout"`).
					WithArgs(
						unity.Id,
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
						unity.NbSerie,
						unity.Comment,
						sqlmock.AnyArg()).
					WillReturnError(fmt.Errorf("database error"))
				mock.ExpectRollback()
			},
			expectedRes: core.NewError(http.StatusInternalServerError, core.ErrDBCreateUnityOfWorkout, fmt.Errorf("database error")),
		},
	}

	for _, value := range setupTest {
		t.Run(value.name, func(t *testing.T) {
			wp := repository.MockWorkoutRepository(db)
			value.setupMock(mock)

			err := wp.CreateUnityOfWorkout(unity)
			assert.Equal(t, err, value.expectedRes)
		})

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	}
}

func TestWorkoutRepository_CREATE_EXERCICE(t *testing.T) {
	var exercice model.Exercice
	exercice.GenerateTest()

	db, mock := setUpDB()

	setupTest := []struct {
		name        string
		setupMock   func(mock sqlmock.Sqlmock)
		expectedRes *core.Error
	}{
		{
			name: core.TestREPExerciceCreateSuccess,
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(`INSERT INTO "exercice"`).
					WithArgs(
						exercice.Id,
						exercice.Name,
						exercice.Equipement).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectedRes: (*core.Error)(nil),
		},
		{
			name: core.TestExercicesCreateFailed,
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(`INSERT INTO "exercice"`).
					WithArgs(
						exercice.Id,
						exercice.Name,
						exercice.Equipement).
					WillReturnError(fmt.Errorf("database error"))
				mock.ExpectRollback()
			},
			expectedRes: core.NewError(http.StatusInternalServerError, core.ErrDBCreateExercice, fmt.Errorf("database error")),
		},
	}

	for _, value := range setupTest {
		t.Run(value.name, func(t *testing.T) {
			wp := repository.MockWorkoutRepository(db)
			value.setupMock(mock)

			err := wp.CreateExcercice(exercice)
			assert.Equal(t, err, value.expectedRes)
		})

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	}
}

func TestWorkoutRepository_CREATE_SERIE(t *testing.T) {
	var serie model.Serie
	serie.GenerateTest()

	db, mock := setUpDB()

	setupTest := []struct {
		name        string
		setupMock   func(mock sqlmock.Sqlmock)
		expectedRes *core.Error
	}{
		{
			name: core.TestREPSerieCreateSuccess,
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(`INSERT INTO "serie"`).
					WithArgs(
						serie.Id,
						serie.Weight,
						serie.Repetitions,
						serie.IsWarmUp).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectedRes: (*core.Error)(nil),
		},
		{
			name: core.TestSeriesCreateFailed,
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(`INSERT INTO "serie"`).
					WithArgs(
						serie.Id,
						serie.Weight,
						serie.Repetitions,
						serie.IsWarmUp).
					WillReturnError(fmt.Errorf("database error"))
				mock.ExpectRollback()
			},
			expectedRes: core.NewError(http.StatusInternalServerError, core.ErrDBCreateSerie, fmt.Errorf("database error")),
		},
	}

	for _, value := range setupTest {
		t.Run(value.name, func(t *testing.T) {
			wp := repository.MockWorkoutRepository(db)
			value.setupMock(mock)

			err := wp.CreateSerie(serie)
			assert.Equal(t, err, value.expectedRes)
		})

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	}
}
