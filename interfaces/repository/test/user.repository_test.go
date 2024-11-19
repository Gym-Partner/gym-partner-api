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
	"time"
)

func SetupDb() (*gorm.DB, sqlmock.Sqlmock) {
	mockDB, mock, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDB,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})

	return db, mock
}

func TestUserRepository_INSERT(t *testing.T) {
	var user model.User
	user.GenerateTestStruct()
	user.CreatedAt = time.Now()

	db, mock := SetupDb()

	setupTest := []struct {
		name        string
		setupMock   func(mock sqlmock.Sqlmock)
		expectedRes model.User
		expectedErr *core.Error
	}{
		{
			name: core.TestREPCreateSuccess,
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(`INSERT INTO "user"`).
					WithArgs(
						user.Id,
						user.FirstName,
						user.LastName,
						user.UserName,
						user.Email,
						user.Password,
						sqlmock.AnyArg(),
					).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectedRes: user,
			expectedErr: (*core.Error)(nil),
		},
		{
			name: core.TestUserCreateFailed,
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(`INSERT INTO "user"`).
					WithArgs(
						user.Id,
						user.FirstName,
						user.LastName,
						user.UserName,
						user.Email,
						user.Password,
						sqlmock.AnyArg(),
					).
					WillReturnError(fmt.Errorf("database error"))
				mock.ExpectRollback()
			},
			expectedRes: model.User{},
			expectedErr: core.NewError(http.StatusInternalServerError, core.ErrDBCreateUser, fmt.Errorf("database error")),
		},
	}

	for _, value := range setupTest {
		t.Run(value.name, func(t *testing.T) {
			up := repository.MockUserRepository(db)
			value.setupMock(mock)

			result, err := up.Create(user)

			switch value.name {
			case core.TestREPCreateSuccess:
				assert.Nil(t, err)
				assert.NotEmpty(t, result)
				assert.Equal(t, result, value.expectedRes)
				assert.Equal(t, err, value.expectedErr)
			case core.TestUserCreateFailed:
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

func TestUserRepository_GETALL(t *testing.T) {
	var users model.Users
	users.GenerateTestStruct().AddCreatedAt()

	db, mock := SetupDb()

	setupTest := []struct {
		name        string
		setupMock   func(mock sqlmock.Sqlmock)
		expectedRes model.Users
		expectedErr *core.Error
	}{
		{
			name: core.TestREPGetAllSuccess,
			setupMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "username", "email"})

				for _, value := range users {
					rows.AddRow(value.Id, value.FirstName, value.LastName, value.UserName, value.Email)
				}

				mock.ExpectQuery(`SELECT id, first_name, last_name, username, email FROM "user"`).
					WillReturnRows(rows)
			},
			expectedRes: users,
			expectedErr: (*core.Error)(nil),
		},
		{
			name: core.TestUsersNotFound,
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`SELECT id, first_name, last_name, username, email FROM "user"`).
					WillReturnError(fmt.Errorf("database error"))
			},
			expectedRes: model.Users{},
			expectedErr: core.NewError(http.StatusInternalServerError, core.ErrDBGetAllUser, fmt.Errorf("database error")),
		},
	}

	for _, value := range setupTest {
		t.Run(value.name, func(t *testing.T) {
			up := repository.MockUserRepository(db)
			value.setupMock(mock)

			result, err := up.GetAll()

			switch value.name {
			case core.TestREPGetAllSuccess:
				assert.Nil(t, err)
				assert.NotEmpty(t, result)
				assert.Equal(t, result, value.expectedRes)
				assert.Equal(t, err, value.expectedErr)
			case core.TestUsersNotFound:
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
