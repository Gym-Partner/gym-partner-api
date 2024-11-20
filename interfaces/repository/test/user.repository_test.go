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

func TestUserRepository_ISEXIST(t *testing.T) {
	db, mock := SetupDb()

	setupTest := []struct {
		name        string
		data        string
		OPT         string
		setupMock   func(sqlmock.Sqlmock)
		expectedRes bool
	}{
		{
			name: core.TestREPIsExistByIdSuccess,
			data: "1234",
			OPT:  "ID",
			setupMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "email", "password", "created_at"}).
					AddRow("1234", "Test", "Test", "test@gmail.com", "aaaAAA111", time.Now())

				mock.ExpectQuery(`SELECT (.+) FROM \"user\" WHERE id =(.+)`).
					WithArgs("1234").
					WillReturnRows(rows)
			},
			expectedRes: true,
		},
		{
			name: core.TestUserNotExistFailed,
			data: "5678",
			OPT:  "ID",
			setupMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "email", "password", "created_at"})
				mock.ExpectQuery(`SELECT (.+) FROM \"user\" WHERE id =(.+)`).
					WithArgs("5678").
					WillReturnRows(rows)
			},
			expectedRes: false,
		},
		{
			name: core.TestREPIsExistByEmailSuccess,
			data: "test@gmail.com",
			OPT:  "EMAIL",
			setupMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "email", "password", "created_at"}).
					AddRow("1234", "Test", "Test", "test@gmail.com", "aaaAAA111", time.Now())

				mock.ExpectQuery(`SELECT (.+) FROM \"user\" WHERE email =(.+)`).
					WithArgs("test@gmail.com").
					WillReturnRows(rows)
			},
			expectedRes: true,
		},
		{
			name: core.TestUserNotExistFailed,
			data: "unknown@gmail.com",
			OPT:  "EMAIL",
			setupMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "email", "password", "created_at"})
				mock.ExpectQuery(`SELECT (.+) FROM \"user\" WHERE email =(.+)`).
					WithArgs("unknown@gmail.com").
					WillReturnRows(rows)
			},
			expectedRes: false,
		},
	}

	for _, value := range setupTest {
		t.Run(value.name, func(t *testing.T) {
			up := repository.MockUserRepository(db)
			value.setupMock(mock)

			result := up.IsExist(value.data, value.OPT)
			assert.Equal(t, result, value.expectedRes)
		})

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	}
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

func TestUserRepository_GETONEBYID(t *testing.T) {
	var user model.User
	user.GenerateTestStruct()

	db, mock := SetupDb()

	setupTest := []struct {
		name        string
		setupMock   func(mock sqlmock.Sqlmock)
		expectedRes model.User
		expectedErr *core.Error
	}{
		{
			name: core.TestREPGetOneByIdSuccess,
			setupMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "username", "email", "password", "created_at"}).
					AddRow(user.Id, user.FirstName, user.LastName, user.UserName, user.Email, user.Password, user.CreatedAt)

				mock.ExpectQuery(`SELECT (.+) FROM \"user\" WHERE id =(.+)`).
					WithArgs(user.Id).
					WillReturnRows(rows)
			},
			expectedRes: user,
			expectedErr: (*core.Error)(nil),
		},
		{
			name: core.TestUserNotFound,
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`SELECT (.+) FROM \"user\" WHERE id =(.+)`).
					WillReturnError(fmt.Errorf("database error"))
			},
			expectedRes: model.User{},
			expectedErr: core.NewError(http.StatusNotFound, fmt.Sprintf(core.ErrDBGetOneUser, user.Id), fmt.Errorf("database error")),
		},
	}

	for _, value := range setupTest {
		t.Run(value.name, func(t *testing.T) {
			up := repository.MockUserRepository(db)
			value.setupMock(mock)

			result, err := up.GetOneById(user.Id)

			switch value.name {
			case core.TestREPGetOneByIdSuccess:
				assert.Nil(t, err)
				assert.NotEmpty(t, result)
				assert.Equal(t, result, value.expectedRes)
				assert.Equal(t, err, value.expectedErr)
			case core.TestUserNotFound:
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

func TestUserRepository_GETONEBYEMAIL(t *testing.T) {}

func TestUserRepository_UPDATE(t *testing.T) {}

func TestUserRepository_DELETE(t *testing.T) {}
