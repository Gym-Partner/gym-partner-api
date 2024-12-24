package test

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/database"
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

func TestWorkoutRepository_GET_WORKOUT(t *testing.T) {
	userId := uuid.New().String()
	var workout database.MigrateWorkout
	workout.GenerateForTest(userId)

	db, mock := setUpDB()

	setupTest := []struct {
		name        string
		setupMock   func(mock sqlmock.Sqlmock)
		expectedRes database.MigrateWorkout
		expectedErr *core.Error
	}{
		{
			name: core.TestREPWorkoutGetSuccess,
			setupMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "user_id", "unities_id", "day", "name", "comment"}).
					AddRow(workout.Id, workout.UserId, workout.UnitiesId, workout.Day, workout.Name, workout.Comment)

				mock.ExpectQuery(`SELECT (.+) FROM \"workout\" WHERE user_id = (.+)`).
					WithArgs(userId).
					WillReturnRows(rows)
			},
			expectedRes: workout,
			expectedErr: (*core.Error)(nil),
		},
		{
			name: core.TestWorkoutGetFailed,
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`SELECT (.+) FROM \"workout\" WHERE user_id =(.+)`).
					WillReturnError(fmt.Errorf("database error"))
			},
			expectedRes: database.MigrateWorkout{},
			expectedErr: core.NewError(http.StatusInternalServerError, core.ErrDBGetWorkout, fmt.Errorf("database error")),
		},
	}

	for _, value := range setupTest {
		t.Run(value.name, func(t *testing.T) {
			wp := repository.MockWorkoutRepository(db)
			value.setupMock(mock)

			result, err := wp.GetOneWorkoutByUserId(userId)

			switch value.name {
			case core.TestREPWorkoutGetSuccess:
				assert.Nil(t, err)
				assert.NotEmpty(t, result)
				assert.Equal(t, result, value.expectedRes)
				assert.Equal(t, err, value.expectedErr)
			case core.TestWorkoutGetFailed:
				assert.NotNil(t, err)
				assert.Empty(t, result)
				assert.Equal(t, result, value.expectedRes)
				assert.Equal(t, err, value.expectedErr)
			}
		})
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	}
}

func TestWorkoutRepository_GET_UNITY_OF_WORKOUT(t *testing.T) {
	var unity database.MigrateUnityOfWorkout
	unity.GenerateForTest(pq.StringArray{
		uuid.New().String(),
		uuid.New().String(),
	})

	db, mock := setUpDB()

	setupTest := []struct {
		name        string
		setupMock   func(mock sqlmock.Sqlmock)
		expectedRes database.MigrateUnityOfWorkout
		expectedErr *core.Error
	}{
		{
			name: core.TestREPUnityOfWorkoutGetSuccess,
			setupMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "exercice_id", "serie_id", "nb_serie", "comment", "rest_time_sec"}).
					AddRow(unity.Id, unity.ExerciceId, unity.SerieId, unity.NbSerie, unity.Comment, unity.RestTimeSec)

				mock.ExpectQuery(`SELECT (.+) FROM \"unity_of_workout\" WHERE id = (.+)`).
					WithArgs(unity.Id).
					WillReturnRows(rows)
			},
			expectedRes: unity,
			expectedErr: (*core.Error)(nil),
		},
		{
			name: core.TestUnitiesOfWorkoutGetFailed,
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`SELECT (.+) FROM \"unity_of_workout\" WHERE id = (.+)`).
					WillReturnError(fmt.Errorf("database error"))
			},
			expectedRes: database.MigrateUnityOfWorkout{},
			expectedErr: core.NewError(
				http.StatusInternalServerError,
				core.ErrDBGetUnityOfWorkout,
				fmt.Errorf("database error")),
		},
	}

	for _, value := range setupTest {
		t.Run(value.name, func(t *testing.T) {
			wp := repository.MockWorkoutRepository(db)
			value.setupMock(mock)

			result, err := wp.GetUntyById(unity.Id)

			switch value.name {
			case core.TestREPUnityOfWorkoutGetSuccess:
				assert.Nil(t, err)
				assert.NotEmpty(t, result)
				assert.Equal(t, result, value.expectedRes)
				assert.Equal(t, err, value.expectedErr)
			case core.TestUnitiesOfWorkoutGetFailed:
				assert.NotNil(t, err)
				assert.Empty(t, result)
				assert.Equal(t, result, value.expectedRes)
				assert.Equal(t, err, value.expectedErr)
			}
		})
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	}
}

func TestWorkoutRepository_GET_EXERCICE(t *testing.T) {
	var exercice database.MigrateExercice
	exercice.GenerateForTest(pq.StringArray{
		uuid.New().String(),
		uuid.New().String(),
	})

	db, mock := setUpDB()

	setupTest := []struct {
		name        string
		setupMock   func(mock sqlmock.Sqlmock)
		expectedRes database.MigrateExercice
		expectedErr *core.Error
	}{
		{
			name: core.TestREPExerciceGetSuccess,
			setupMock: func(mock sqlmock.Sqlmock) {
				row := sqlmock.NewRows([]string{"id", "name", "equipement"}).
					AddRow(exercice.Id, exercice.Name, exercice.Equipement)

				mock.ExpectQuery(`SELECT (.+) FROM \"exercice\" WHERE id = (.+)`).
					WithArgs(exercice.Id).
					WillReturnRows(row)
			},
			expectedRes: exercice,
			expectedErr: (*core.Error)(nil),
		},
		{
			name: core.TestExercicesGetFailed,
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`SELECT (.+) FROM \"exercice\" WHERE id = (.+)`).
					WillReturnError(fmt.Errorf("database error"))
			},
			expectedRes: database.MigrateExercice{},
			expectedErr: core.NewError(
				http.StatusInternalServerError,
				core.ErrDBGetExercice,
				fmt.Errorf("database error")),
		},
	}

	for _, value := range setupTest {
		t.Run(value.name, func(t *testing.T) {
			wp := repository.MockWorkoutRepository(db)
			value.setupMock(mock)

			result, err := wp.GetExerciceById(exercice.Id)

			switch value.name {
			case core.TestREPExerciceGetSuccess:
				assert.Nil(t, err)
				assert.NotEmpty(t, result)
				assert.Equal(t, result, value.expectedRes)
				assert.Equal(t, err, value.expectedErr)
			case core.TestExercicesGetFailed:
				assert.NotNil(t, err)
				assert.Empty(t, result)
				assert.Equal(t, result, value.expectedRes)
				assert.Equal(t, err, value.expectedErr)
			}
		})
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	}
}

func TestWorkoutRepository_GET_SERIE(t *testing.T) {
	var serie database.MigrateSerie
	serie.GenerateForTest(pq.StringArray{
		uuid.New().String(),
		uuid.New().String(),
	})

	db, mock := setUpDB()

	setupTest := []struct {
		name        string
		setupMock   func(mock sqlmock.Sqlmock)
		expectedRes database.MigrateSerie
		expectedErr *core.Error
	}{
		{
			name: core.TESTREPSerieGetSuccess,
			setupMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "weight", "repetitions", "is_warm_up"}).
					AddRow(serie.Id, serie.Weight, serie.Repetitions, serie.IsWarmUp)

				mock.ExpectQuery(`SELECT (.+) FROM \"serie\" WHERE id = (.+)`).
					WithArgs(serie.Id).
					WillReturnRows(rows)
			},
			expectedRes: serie,
			expectedErr: (*core.Error)(nil),
		},
		{
			name: core.TestSeriesGetFailed,
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`SELECT (.+) FROM \"serie\" WHERE id = (.+)`).
					WillReturnError(fmt.Errorf("database error"))
			},
			expectedRes: database.MigrateSerie{},
			expectedErr: core.NewError(
				http.StatusInternalServerError,
				core.ErrDBGetSerie,
				fmt.Errorf("database error")),
		},
	}

	for _, value := range setupTest {
		t.Run(value.name, func(t *testing.T) {
			wp := repository.MockWorkoutRepository(db)
			value.setupMock(mock)

			result, err := wp.GetSerieById(serie.Id)

			switch value.name {
			case core.TESTREPSerieGetSuccess:
				assert.Nil(t, err)
				assert.NotEmpty(t, result)
				assert.Equal(t, result, value.expectedRes)
				assert.Equal(t, err, value.expectedErr)
			case core.TestSeriesGetFailed:
				assert.NotNil(t, err)
				assert.Empty(t, result)
				assert.Equal(t, result, value.expectedRes)
				assert.Equal(t, err, value.expectedErr)
			}
		})
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	}
}
