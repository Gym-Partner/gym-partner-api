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
				mock.ExpectCommit()
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
