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
}

type UserInteractor struct {
	IUserRepository   repository.IUserRepository
	IFollowRepository repository.IFollowRepository
	IUtils            utils.IUtils[model.User]
}

func MockUserInteractor(userMock *mock.UserInteractorMock, utilsMock *mock.UtilsMock[model.User]) *UserInteractor {
	return &UserInteractor{
		IUserRepository: userMock,
		IUtils:          utilsMock,
	}
}

// -------------------------- CRUD ------------------------------

func (ui *UserInteractor) Create(ctx *gin.Context) (model.User, *core.Error) {
	data, err := ui.IUtils.InjectBodyInModel(ctx)
	if err != nil {
		return model.User{}, err
	}

	exist := ui.IUserRepository.IsExist(data.Email, "EMAIL")
	if exist {
		return model.User{}, core.NewError(
			http.StatusBadRequest,
			fmt.Sprintf(core.ErrAppINTUserExist, data.Email))
	}

	data.Id = ui.IUtils.GenerateUUID()

	data.Password, _ = ui.IUtils.HashPassword(data.Password)
	user, err := ui.IUserRepository.Create(data)
	return user, err
}

func (ui *UserInteractor) GetAll() (model.Users, *core.Error) {
	users, err := ui.IUserRepository.GetAll()
	return users, err
}

func (ui *UserInteractor) GetOne(c *gin.Context) (model.User, *core.Error) {
	uid, _ := c.Get("uid")

	user, err := ui.IUserRepository.GetOneById(uid.(string))
	if err != nil {
		return model.User{}, err
	}

	// Followers part
	followers, err := ui.IFollowRepository.GetAllByUserId(uid.(string))
	if err != nil {
		return model.User{}, err
	}

	user.Followers = followers.Followers
	user.Following = followers.Followings
	return user, nil
}

func (ui *UserInteractor) GetOneByEmail(ctx *gin.Context) (model.User, *core.Error) {
	data, err := ui.IUtils.InjectBodyInModel(ctx)
	if err != nil {
		return model.User{}, err
	}

	user, err := ui.IUserRepository.GetOneByEmail(data.Email)
	if err != nil {
		return model.User{}, err
	}

	followers, err := ui.IFollowRepository.GetAllByUserId(data.Id)
	if err != nil {
		return model.User{}, err
	}

	user.Password = data.Password
	user.Followers = followers.Followers
	user.Following = followers.Followings
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
		return core.NewError(
			http.StatusBadRequest,
			fmt.Sprintf(core.ErrAppINTUserNotExist, patch.Id))
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
	uid, _ := ctx.Get("uid")

	exist := ui.IUserRepository.IsExist(*uid.(*string), "ID")
	if !exist {
		return core.NewError(
			http.StatusBadRequest,
			fmt.Sprintf(core.ErrAppINTUserNotExist, *uid.(*string)))
	}

	err := ui.IUserRepository.Delete(*uid.(*string))
	return err
}
