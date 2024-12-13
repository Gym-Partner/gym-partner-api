package test

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
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
			name: core.TestUnitiesOfWorkoutFailed,
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
			name: core.TestExercicesFailed,
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
			name: core.TestExercicesFailed,
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
								core.ErrDBCreateExercice)).Maybe()
					}
				}
			},
			expectedRes: core.NewError(http.StatusInternalServerError, core.ErrDBCreateExercice),
		},
	}

	for _, value := range setupTest {
		t.Run(value.name, func(t *testing.T) {
			workoutMock := new(mock.WorkoutInteractorMock)
			utilsMock := new(mock.UtilsMock[model.Workout])

			wi := interactor.MockWorkoutInteractor(workoutMock, utilsMock)

			buf, _ := utils.StructToReadCloser(workout)
			context := &gin.Context{
				Request: &http.Request{
					Body: buf,
				},
			}
			context.Set("uid", &userId)

			value.setupMock(workoutMock, utilsMock, context)

			result := wi.Create(context)

			assert.Equal(t, result, value.expectedRes)

			workoutMock.AssertExpectations(t)
			utilsMock.AssertExpectations(t)
		})
	}
}

func TestWorkoutInteractor_GETONEBYUSERID(t *testing.T) {}
