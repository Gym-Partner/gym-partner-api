package controller

import (
	"encoding/json"
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"gitlab.com/gym-partner1/api/gym-partner-api/mock"

	"github.com/gin-gonic/gin"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/interfaces/repository"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
	"gitlab.com/gym-partner1/api/gym-partner-api/services/interactor"
	"gitlab.com/gym-partner1/api/gym-partner-api/utils"
)

type UserController struct {
	IUserInteractor interactor.IUserInteractor
	Log             *core.Log
	Rabbit          *core.RabbitMQ
}

var mailStatusMap = sync.Map{}

// ------------------------------ Constructor ------------------------------

func NewUserController(db *core.Database, rabbit *core.RabbitMQ) *UserController {
	return &UserController{
		IUserInteractor: &interactor.UserInteractor{
			IUserRepository: repository.UserRepository{
				DB:  db.Handler,
				Log: db.Logger,
			},
			IFollowRepository: repository.FollowRepository{
				DB:  db.Handler,
				Log: db.Logger,
			},
			IUtils: utils.Utils[model.User]{},
		},
		Log:    db.Logger,
		Rabbit: rabbit,
	}
}

// ------------------------------ Mock Constructor ------------------------------

func MockUserController(UCMock *mock.UserControllerMock) *UserController {
	log := core.NewLog("/Users/oscar/Documents/gym-partner-env", true)
	log.ChargeLog()

	return &UserController{
		IUserInteractor: UCMock,
		Log:             log,
	}
}

// ------------------------------ CRUD ------------------------------

// Create godoc
// @Summary Create a new user
// @Description Create new user in database and return the created user without the password.
// @Tags User
// @Accept json
// @Produce json
// @Param user body model.User{} true "User data"
// @Success 201 {object} model.User "User successfully created"
// @Failure 400 {object} core.Error "User already exist in database."
// @Failure 500 {object} core.Error "Internal server error"
// @Router /user/create [post]
func (uc *UserController) Create(ctx *gin.Context) {
	user, err := uc.IUserInteractor.Create(ctx)
	if err != nil {
		ctx.JSON(err.Code, err.Respons())
		return
	}

	//if err := uc.Rabbit.PublishMessage("User created successfully"); err != nil {
	//	ctx.JSON(http.StatusInternalServerError, err.Error())
	//	return
	//}

	ctx.JSON(http.StatusCreated, user.Response())
}

// GetAll godoc
// @Summary Retrieve all user
// @Description Retrieve all user in database and return this without password.
// @Tags User
// @Produce json
// @Param Authorization header string true "User's Token"
// @Success 200 {object} model.Users "Users successfully retrieves"
// @Failure 500 {object} core.Error "Internal server error"
// @Router /user/getAll [get]
func (uc *UserController) GetAll(ctx *gin.Context) {
	users, err := uc.IUserInteractor.GetAll()
	if err != nil {
		ctx.JSON(err.Code, err.Respons())
		return
	}

	ctx.JSON(http.StatusOK, users.Response())
}

// GetOne godoc
// @Summary Retrieve one user
// @Description Retrieve one user with id in token and return this without password.
// @Tags User
// @Produce json
// @Param Authorization header string true "User Token"
// @Success 200 {object} model.User "User successfully retrieve"
// @Failure 500 {object} core.Error "Internal server error"
// @Router /user/getOne [get]
func (uc *UserController) GetOne(ctx *gin.Context) {
	user, err := uc.IUserInteractor.GetOne(ctx)
	if err != nil {
		ctx.JSON(err.Code, err.Respons())
		return
	}

	ctx.JSON(http.StatusOK, user.Response())
}

// GetUsers godoc
// @Summary Retrieve a list of users
// @Description Accepts a raw array of user IDs (JSON array of strings).
// @Tags User
// @Accept json
// @Produce json
// @Param Authorization header string true "User's token"
// @Param request body model.GetUsersRequestBody true "Array of user IDs (as raw JSON)"
// @Success 200 {object} model.Users "Users successfully retrieves"
// @Failure 500 {object} core.Error "Internal server error"
// @Router /users/get_users [post]
func (uc *UserController) GetUsers(ctx *gin.Context) {
	users, err := uc.IUserInteractor.GetUsers(ctx)
	if err != nil {
		ctx.JSON(err.Code, err.Respons())
		return
	}

	ctx.JSON(http.StatusOK, users.Response())
}

// Search godoc
// @Summary Retrieves users in search bar
// @Description Retrieve users in search bar, with the first letter of his first_name / username / email.
// @Tags User
// @Produce json
// @Param Authorization header string true "User's token"
// @Param Params query string true "Search query (min 3 characters)"
// @Param limit query int false "Max number of users to return (default: 10, max: 50)"
// @Param offset query int false "Number of users to skip (default: 0)"
// @Success 200 {object} model.Users "Matching users"
// @Failure 400 {object} core.Error "Query too short or invalid"
// @Failure 500 {object} core.Error "Internal server error
// @Router /users/search [get]
func (uc *UserController) Search(ctx *gin.Context) {
	query := strings.ToLower(ctx.Query("query"))
	limitStr := ctx.DefaultQuery("limit", "10")
	offsetStr := ctx.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit > 50 {
		limit = 10
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = 0
	}

	if len(query) < 3 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "query too short"})
		return
	}

	users, searchErr := uc.IUserInteractor.Search(query, limit, offset)
	if searchErr != nil {
		ctx.JSON(searchErr.Code, searchErr.Respons())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  users,
		"count": len(users),
	})
}

// UploadImage godoc
// @Summary Upload user's image
// @Description Upload user's image in S3 aws's service and his url in database.
// @Tags User
// @Accept multipart/form-data
// @Produce json
// @Param Authorization header string true "User's token"
// @Param image formData file true "User's profile image (JPEG, PNG, max TMB)"
// @Success 201 {object} model.UsersImage "User's image uploaded successfully"
// @Failure 406 {object} core.Error "User's image not available in request body"
// @Failure 500 {object} core.Error "Internal server error"
// @Router /user/upload_image [post]
func (uc *UserController) UploadImage(ctx *gin.Context) {
	userImage, err := uc.IUserInteractor.UploadImage(ctx)
	if err != nil {
		ctx.JSON(err.Code, err.Respons())
		return
	}

	ctx.JSON(http.StatusCreated, userImage)
}

// Update godoc
// @Summary Update one user
// @Description Update on user and return nil
// @Tags User
// @Accept json
// @Param Authorization header string true "User Token"
// @Param user_update body model.User{} true "User data update"
// @Success 200 {object} nil "User successfully updated"
// @Failure 400 {object} core.Error{} "User not exist in database"
// @Failure 500 {object} core.Error{} "Internal server error"
// @Router /user/update [patch]
func (uc *UserController) Update(ctx *gin.Context) {
	if err := uc.IUserInteractor.Update(ctx); err != nil {
		ctx.JSON(err.Code, err.Respons())
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

// Delete godoc
// @Summary Delete one user
// @Description Delete one user and return nil
// @Tags User
// @Param Authorization header string true "User Token"
// @Success 200 {object} nil "User successfully deleted"
// @Failure 400 {object} core.Error{} "User not exist in database"
// @Failure 500 {object} core.Error{} "Internal server error"
// @Router /user/delete [delete]
func (uc *UserController) Delete(ctx *gin.Context) {
	if err := uc.IUserInteractor.Delete(ctx); err != nil {
		ctx.JSON(err.Code, err.Respons())
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func (uc *UserController) RabbitMQTest(ctx *gin.Context) {
	var user model.User
	user.GenerateTestStruct()

	corrID := uuid.New().String()
	mailStatusMap.Store(corrID, "pending")

	body, _ := json.Marshal(user)

	if err := uc.Rabbit.Publish(core.QueueAPI, body, amqp.Publishing{
		ReplyTo:       string(core.QueueSMTP),
		Type:          "application/json",
		CorrelationId: corrID,
	}); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"correlation_id": corrID,
		"status":         "pending",
	})
}

func (uc *UserController) StartReplyConsumer() {
	msgs, err := uc.Rabbit.Consume(core.QueueSMTP)
	if err != nil {
		panic(err)
	}

	go func() {
		for msg := range msgs {
			var res map[string]string

			_ = json.Unmarshal(msg.Body, &res)
			status := res["status"]
			corrID := msg.CorrelationId

			mailStatusMap.Store(corrID, status)
			log.Println("Response from service B: ", res)
		}
	}()
}
