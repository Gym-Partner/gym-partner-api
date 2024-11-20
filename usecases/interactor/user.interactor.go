package interactor

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/mock"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
	"gitlab.com/gym-partner1/api/gym-partner-api/usecases/repository"
	"gitlab.com/gym-partner1/api/gym-partner-api/utils"
)

type IUserInteractor interface {
	Create(ctx *gin.Context) (model.User, *core.Error)
	GetAll() (model.Users, *core.Error)
	GetOne(c *gin.Context) (model.User, *core.Error)
	GetOneByEmail(ctx *gin.Context) (model.User, *core.Error)
	Update(ctx *gin.Context) *core.Error
	Delete(ctx *gin.Context) *core.Error
	Login(user model.User) (string, *core.Error)
}

type UserInteractor struct {
	IUserRepository repository.IUserRepository
	IUtils          utils.IUtils[model.User]
	ICognito        core.ICognito
}

func MockUserInteractor(userMock *mock.UserMock, utilsMock *mock.UtilsMock[model.User], cognitoMock *mock.CognitoMock) *UserInteractor {
	return &UserInteractor{
		IUserRepository: userMock,
		IUtils:          utilsMock,
		ICognito:        cognitoMock,
	}
}

// -------------------------- CRUD ------------------------------

func (ui *UserInteractor) Create(ctx *gin.Context) (model.User, *core.Error) {
	var userCognito model.User

	data, err := ui.IUtils.InjectBodyInModel(ctx)
	if err != nil {
		return model.User{}, err
	}

	exist := ui.IUserRepository.IsExist(data.Email, "EMAIL")
	if exist {
		return model.User{}, core.NewError(http.StatusBadRequest, core.ErrIntUserExist)
	}

	data.Id = ui.IUtils.GenerateUUID()

	userCognito.UserToAnother(data)

	data.Password, _ = ui.IUtils.HashPassword(data.Password)
	user, err := ui.IUserRepository.Create(data)
	if err != nil {
		return user, core.NewError(http.StatusInternalServerError, core.ErrDBCreateUser, err)
	}

	if err = ui.ICognito.SignUp(userCognito); err != nil {
		return user, core.NewError(http.StatusBadRequest, core.ErrIntCreateUserAWS, err)
	}

	return user, err
}

func (ui *UserInteractor) GetAll() (model.Users, *core.Error) {
	users, err := ui.IUserRepository.GetAll()
	return users, err
}

func (ui *UserInteractor) GetOne(c *gin.Context) (model.User, *core.Error) {
	uid, _ := c.Get("uid")

	user, err := ui.IUserRepository.GetOneById(*uid.(*string))
	return user, err
}

func (ui *UserInteractor) GetOneByEmail(ctx *gin.Context) (model.User, *core.Error) {
	data, err := ui.IUtils.InjectBodyInModel(ctx)
	if err != nil {
		return model.User{}, err
	}

	user, err := ui.IUserRepository.GetOneByEmail(data.Email)
	user.Password = data.Password
	return user, err
}

func (ui *UserInteractor) Update(ctx *gin.Context) *core.Error {
	uid, _ := ctx.Get("uid")
	patch, err := ui.IUtils.InjectBodyInModel(ctx)
	if err != nil {
		return err
	}
	patch.Id = *uid.(*string)

	exist := ui.IUserRepository.IsExist(patch.Id, "ID")
	if !exist {
		return core.NewError(http.StatusBadRequest, fmt.Sprintf(core.ErrIntUserNotExist, patch.Id))
	}

	target, err := ui.IUserRepository.GetOneById(patch.Id)
	if err != nil {
		return err
	}

	if err = ui.IUtils.Bind(&target, patch); err != nil {
		return err
	}

	err = ui.IUserRepository.Update(target)
	return err
}

func (ui *UserInteractor) Delete(ctx *gin.Context) *core.Error {
	token, _ := ctx.Get("token")
	uid, _ := ctx.Get("uid")

	exist := ui.IUserRepository.IsExist(*uid.(*string), "ID")
	if !exist {
		return core.NewError(http.StatusBadRequest, fmt.Sprintf(core.ErrIntUserNotExist, *uid.(*string)))
	}

	if err := ui.IUserRepository.Delete(*uid.(*string)); err != nil {
		return err
	}

	if err := ui.ICognito.DeleteUser(token.(string)); err != nil {
		return err
	}

	return nil
}

func (ui *UserInteractor) Login(user model.User) (string, *core.Error) {
	token, err := ui.ICognito.SignIn(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
