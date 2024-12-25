package test

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/interfaces/controller"
	"gitlab.com/gym-partner1/api/gym-partner-api/mock"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
	"gitlab.com/gym-partner1/api/gym-partner-api/utils"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWorkoutController_CREATE(t *testing.T) {
	var workout model.Workout
	workout.GenerateTestWorkout()

	setupTest := []struct {
		name         string
		setupMock    func(wcm *mock.WorkoutControllerMock, ctx *gin.Context)
		expectedCode int
		expectedBody gin.H
	}{
		{
			name: core.TestCONWorkoutCreateSuccess,
			setupMock: func(wcm *mock.WorkoutControllerMock, ctx *gin.Context) {
				wcm.On("Create", ctx).Return((*core.Error)(nil))
			},
			expectedCode: http.StatusCreated,
			expectedBody: nil,
		},
		{
			name: core.TestWorkoutCreateFailed,
			setupMock: func(wcm *mock.WorkoutControllerMock, ctx *gin.Context) {
				wcm.On("Create", ctx).Return(core.NewError(
					http.StatusInternalServerError,
					core.ErrDBCreateWorkout))
			},
			expectedCode: http.StatusInternalServerError,
			expectedBody: core.NewError(
				http.StatusInternalServerError,
				core.ErrDBCreateWorkout).Respons(),
		},
	}

	for _, value := range setupTest {
		t.Run(value.name, func(t *testing.T) {
			WorkoutControllerMock := new(mock.WorkoutControllerMock)

			wc := controller.MockWorkoutController(WorkoutControllerMock)

			gin.SetMode(gin.TestMode)
			res := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(res)

			buf, _ := utils.StructToReadCloser(workout)
			ctx.Request = &http.Request{
				Body: buf,
			}

			value.setupMock(WorkoutControllerMock, ctx)

			wc.Create(ctx)
			assert.Equal(t, value.expectedCode, res.Code)

			var response gin.H
			err := json.Unmarshal(res.Body.Bytes(), &response)
			transformResponse(response)

			assert.NoError(t, err)
			assert.Equal(t, value.expectedBody, response)

			WorkoutControllerMock.AssertExpectations(t)
		})
	}
}

func TestWorkoutController_GET_ONE(t *testing.T) {
	var workout model.Workout
	workout.GenerateTestWorkout()

	setupTest := []struct {
		name         string
		setupMock    func(wcm *mock.WorkoutControllerMock, ctx *gin.Context)
		expectedCode int
		expectedBody gin.H
	}{
		{
			name: core.TestCONWorkoutGetSuccess,
			setupMock: func(wcm *mock.WorkoutControllerMock, ctx *gin.Context) {
				wcm.On("GetOneByUserId", ctx).Return(workout, (*core.Error)(nil))
			},
			expectedCode: http.StatusOK,
			expectedBody: workout.Respons(),
		},
		{
			name: core.TestWorkoutGetFailed,
			setupMock: func(wcm *mock.WorkoutControllerMock, ctx *gin.Context) {
				wcm.On("GetOneByUserId", ctx).Return(model.Workout{}, core.NewError(
					http.StatusInternalServerError,
					core.ErrDBGetWorkout))
			},
			expectedCode: http.StatusInternalServerError,
			expectedBody: core.NewError(
				http.StatusInternalServerError,
				core.ErrDBGetWorkout).Respons(),
		},
	}

	for _, value := range setupTest {
		t.Run(value.name, func(t *testing.T) {
			WorkoutControllerMock := new(mock.WorkoutControllerMock)
			wc := controller.MockWorkoutController(WorkoutControllerMock)

			gin.SetMode(gin.TestMode)
			res := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(res)

			ctx.Request = &http.Request{
				Header: http.Header{
					"Authorization": []string{TOKEN},
				},
			}

			value.setupMock(WorkoutControllerMock, ctx)

			wc.GetOneByUserId(ctx)
			assert.Equal(t, value.expectedCode, res.Code)

			var response gin.H
			err := json.Unmarshal(res.Body.Bytes(), &response)
			transformResponse(response)

			assert.NoError(t, err)
			assert.Equal(t, value.expectedBody, response)

			WorkoutControllerMock.AssertExpectations(t)
		})
	}
}
