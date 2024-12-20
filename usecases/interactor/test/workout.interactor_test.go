package test

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/database"
	"gitlab.com/gym-partner1/api/gym-partner-api/mock"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
	"gitlab.com/gym-partner1/api/gym-partner-api/usecases/interactor"
	"gitlab.com/gym-partner1/api/gym-partner-api/utils"
	"net/http"
	"testing"
)

func TestWorkoutInteractor_CREATE(t *testing.T) {
	userId := uuid.New().String()
	var workout model.Workout
	workout.GenerateTestWorkout(userId)

	setupTest := []struct {
		name        string
		setupMock   func(workoutMock *mock.WorkoutInteractorMock, utilsMock *mock.UtilsMock[model.Workout], ctx *gin.Context)
		expectedRes *core.Error
	}{
		{
			name: core.TestINTWorkoutCreateSuccess,
			setupMock: func(workoutMock *mock.WorkoutInteractorMock, utilsMock *mock.UtilsMock[model.Workout], ctx *gin.Context) {
				utilsMock.On("InjectBodyInModel", ctx).Return(workout, (*core.Error)(nil))
				workoutMock.On("CreateWorkout", workout).Return((*core.Error)(nil))
				for _, unity := range workout.UnitiesOfWorkout {
					workoutMock.On("CreateUnityOfWorkout", unity).Return((*core.Error)(nil))
					for _, exercice := range unity.Exercices {
						workoutMock.On("CreateExcercice", exercice).Return((*core.Error)(nil))
					}
					for _, serie := range unity.Series {
						workoutMock.On("CreateSerie", serie).Return((*core.Error)(nil))
					}
				}
			},
			expectedRes: (*core.Error)(nil),
		},
		{
			name: core.TestWorkoutCreateFailed,
			setupMock: func(workoutMock *mock.WorkoutInteractorMock, utilsMock *mock.UtilsMock[model.Workout], ctx *gin.Context) {
				utilsMock.On("InjectBodyInModel", ctx).Return(workout, (*core.Error)(nil))
				workoutMock.On("CreateWorkout", workout).Return(
					core.NewError(http.StatusInternalServerError,
						core.ErrDBCreateWorkout))
			},
			expectedRes: core.NewError(http.StatusInternalServerError, core.ErrDBCreateWorkout),
		},
		{
			name: core.TestUnitiesOfWorkoutCreateFailed,
			setupMock: func(workoutMock *mock.WorkoutInteractorMock, utilsMock *mock.UtilsMock[model.Workout], ctx *gin.Context) {
				utilsMock.On("InjectBodyInModel", ctx).Return(workout, (*core.Error)(nil))
				workoutMock.On("CreateWorkout", workout).Return((*core.Error)(nil))
				for _, unity := range workout.UnitiesOfWorkout {
					workoutMock.On("CreateUnityOfWorkout", unity).Return(
						core.NewError(http.StatusInternalServerError,
							core.ErrDBCreateUnityOfWorkout)).Maybe()
				}
			},
			expectedRes: core.NewError(http.StatusInternalServerError, core.ErrDBCreateUnityOfWorkout),
		},
		{
			name: core.TestExercicesCreateFailed,
			setupMock: func(workoutMock *mock.WorkoutInteractorMock, utilsMock *mock.UtilsMock[model.Workout], ctx *gin.Context) {
				utilsMock.On("InjectBodyInModel", ctx).Return(workout, (*core.Error)(nil))
				workoutMock.On("CreateWorkout", workout).Return((*core.Error)(nil))
				for _, unity := range workout.UnitiesOfWorkout {
					workoutMock.On("CreateUnityOfWorkout", unity).Return((*core.Error)(nil)).Maybe()
					for _, exercice := range unity.Exercices {
						workoutMock.On("CreateExcercice", exercice).Return(
							core.NewError(http.StatusInternalServerError,
								core.ErrDBCreateExercice)).Maybe()
					}
				}
			},
			expectedRes: core.NewError(http.StatusInternalServerError, core.ErrDBCreateExercice),
		},
		{
			name: core.TestSeriesCreateFailed,
			setupMock: func(workoutMock *mock.WorkoutInteractorMock, utilsMock *mock.UtilsMock[model.Workout], ctx *gin.Context) {
				utilsMock.On("InjectBodyInModel", ctx).Return(workout, (*core.Error)(nil))
				workoutMock.On("CreateWorkout", workout).Return((*core.Error)(nil))
				for _, unity := range workout.UnitiesOfWorkout {
					workoutMock.On("CreateUnityOfWorkout", unity).Return((*core.Error)(nil)).Maybe()
					for _, exercice := range unity.Exercices {
						workoutMock.On("CreateExcercice", exercice).Return((*core.Error)(nil)).Maybe()
					}
					for _, serie := range unity.Series {
						workoutMock.On("CreateSerie", serie).Return(
							core.NewError(http.StatusInternalServerError,
								core.ErrDBCreateSerie)).Maybe()
					}
				}
			},
			expectedRes: core.NewError(http.StatusInternalServerError, core.ErrDBCreateSerie),
		},
	}

	for _, value := range setupTest {
		t.Run(value.name, func(t *testing.T) {
			workoutMock := new(mock.WorkoutInteractorMock)
			utilsMock := new(mock.UtilsMock[model.Workout])
			buf, _ := utils.StructToReadCloser(workout)
			context := &gin.Context{
				Request: &http.Request{
					Body: buf,
				},
			}
			context.Set("uid", &userId)

			wi := interactor.MockWorkoutInteractor(workoutMock, utilsMock)
			value.setupMock(workoutMock, utilsMock, context)
			result := wi.Create(context)

			assert.Equal(t, result, value.expectedRes)

			workoutMock.AssertExpectations(t)
			utilsMock.AssertExpectations(t)
		})
	}
}

func TestWorkoutInteractor_GETONEBYUSERID(t *testing.T) {
	userId := uuid.New().String()

	var workout model.Workout
	workout.GenerateTestWorkout(userId)

	var migrateWorkout database.MigrateWorkout
	migrateWorkout.GenerateForTest(userId)

	var migrateUnityOfWorkout database.MigrateUnityOfWorkout
	migrateUnityOfWorkout.GenerateForTest(migrateWorkout.UnitiesId)

	var migrateExercice database.MigrateExercice
	migrateExercice.GenerateForTest(migrateUnityOfWorkout.ExerciceId)

	var migrateSerie database.MigrateSerie
	migrateSerie.GenerateForTest(migrateUnityOfWorkout.SerieId)

	var migrateUnitiesOfWorkout database.MigrateUnitiesOfWorkout
	var migrateExercices database.MigrateExercices
	var migrateSeries database.MigrateSeries

	setupTest := []struct {
		name        string
		setupMock   func(workoutMock *mock.WorkoutInteractorMock, utilsMock *mock.UtilsMock[model.Workout], ctx *gin.Context)
		expectedRes model.Workout
		expectedErr *core.Error
	}{
		{
			name: core.TestINTWorkoutGetSuccess,
			setupMock: func(workoutMock *mock.WorkoutInteractorMock, utilsMock *mock.UtilsMock[model.Workout], ctx *gin.Context) {
				workoutMock.On("GetOneWorkoutByUserId", userId).Return(migrateWorkout, (*core.Error)(nil))
				for _, unityId := range migrateWorkout.UnitiesId {
					workoutMock.On("GetUntyById", unityId).Return(migrateUnityOfWorkout, (*core.Error)(nil))
					migrateUnitiesOfWorkout.GenerateForTest(migrateUnityOfWorkout)
					for _, exerciceId := range migrateUnityOfWorkout.ExerciceId {
						workoutMock.On("GetExerciceById", exerciceId).Return(migrateExercice, (*core.Error)(nil))
						migrateExercices.GenerateForTest(migrateExercice)
					}
					for _, serieId := range migrateUnityOfWorkout.SerieId {
						workoutMock.On("GetSerieById", serieId).Return(migrateSerie, (*core.Error)(nil))
						migrateSeries.GenerateForTest(migrateSerie)
					}
				}
				utilsMock.SchemaToModel(migrateWorkout, migrateUnitiesOfWorkout, migrateExercices, migrateSeries)
			},
			expectedRes: workout,
			expectedErr: (*core.Error)(nil),
		},
		{
			name: core.TestWorkoutGetFailed,
			setupMock: func(workoutMock *mock.WorkoutInteractorMock, utilsMock *mock.UtilsMock[model.Workout], ctx *gin.Context) {
				workoutMock.On("GetOneWorkoutByUserId", userId).Return(
					database.MigrateWorkout{},
					core.NewError(http.StatusInternalServerError, core.ErrDBGetWorkout)).Maybe()
			},
			expectedRes: model.Workout{},
			expectedErr: core.NewError(http.StatusInternalServerError, core.ErrDBGetWorkout),
		},
		{
			name: core.TestUnitiesOfWorkoutGetFailed,
			setupMock: func(workoutMock *mock.WorkoutInteractorMock, utilsMock *mock.UtilsMock[model.Workout], ctx *gin.Context) {
				workoutMock.On("GetOneWorkoutByUserId", userId).Return(migrateWorkout, (*core.Error)(nil))
				for _, unityId := range migrateWorkout.UnitiesId {
					workoutMock.On("GetUntyById", unityId).Return(
						database.MigrateUnityOfWorkout{},
						core.NewError(http.StatusInternalServerError, core.ErrDBGetUnityOfWorkout)).Maybe()
				}
			},
			expectedRes: model.Workout{},
			expectedErr: core.NewError(http.StatusInternalServerError, core.ErrDBGetUnityOfWorkout),
		},
		{
			name: core.TestExercicesGetFailed,
			setupMock: func(workoutMock *mock.WorkoutInteractorMock, utilsMock *mock.UtilsMock[model.Workout], ctx *gin.Context) {
				workoutMock.On("GetOneWorkoutByUserId", userId).Return(migrateWorkout, (*core.Error)(nil))
				for _, unityId := range migrateWorkout.UnitiesId {
					workoutMock.On("GetUntyById", unityId).Return(migrateUnityOfWorkout, (*core.Error)(nil)).Maybe()
					for _, exerciceId := range migrateUnityOfWorkout.ExerciceId {
						workoutMock.On("GetExerciceById", exerciceId).Return(
							database.MigrateExercice{},
							core.NewError(http.StatusInternalServerError, core.ErrDBGetExercice)).Maybe()
					}
				}
			},
			expectedRes: model.Workout{},
			expectedErr: core.NewError(http.StatusInternalServerError, core.ErrDBGetExercice),
		},
		{
			name: core.TestSeriesGetFailed,
			setupMock: func(workoutMock *mock.WorkoutInteractorMock, utilsMock *mock.UtilsMock[model.Workout], ctx *gin.Context) {
				workoutMock.On("GetOneWorkoutByUserId", userId).Return(migrateWorkout, (*core.Error)(nil))
				for _, unityId := range migrateWorkout.UnitiesId {
					workoutMock.On("GetUntyById", unityId).Return(migrateUnityOfWorkout, (*core.Error)(nil)).Maybe()
					for _, exerciceId := range migrateUnityOfWorkout.ExerciceId {
						workoutMock.On("GetExerciceById", exerciceId).Return(migrateExercice, (*core.Error)(nil)).Maybe()
					}
					for _, serieId := range migrateUnityOfWorkout.SerieId {
						workoutMock.On("GetSerieById", serieId).Return(
							database.MigrateSerie{},
							core.NewError(http.StatusInternalServerError, core.ErrDBGetSerie)).Maybe()
					}
				}
			},
			expectedRes: model.Workout{},
			expectedErr: core.NewError(http.StatusInternalServerError, core.ErrDBGetSerie),
		},
	}

	for _, value := range setupTest {
		t.Run(value.name, func(t *testing.T) {
			workoutMock := new(mock.WorkoutInteractorMock)
			utilsMock := new(mock.UtilsMock[model.Workout])
			context := &gin.Context{}
			context.Set("uid", &userId)

			wi := interactor.MockWorkoutInteractor(workoutMock, utilsMock)
			value.setupMock(workoutMock, utilsMock, context)
			users, err := wi.GetOneByUserId(context)

			assert.Equal(t, users, value.expectedRes)
			assert.Equal(t, err, value.expectedErr)

			workoutMock.AssertExpectations(t)
			utilsMock.AssertExpectations(t)
		})
	}
}
