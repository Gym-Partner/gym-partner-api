package interactor

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "gitlab.com/gym-partner1/api/gym-partner-api/core"
    "gitlab.com/gym-partner1/api/gym-partner-api/domain/model"
    "gitlab.com/gym-partner1/api/gym-partner-api/usecases/repository"
    "gitlab.com/gym-partner1/api/gym-partner-api/utils"
)

type UserInteractor struct {
    IUserRepository repository.IUserRepository
}

// -------------------------- CRUD ------------------------------

func (ui *UserInteractor) Create(c *gin.Context) (model.User, *core.Error) {
    data, err := utils.InjectBodyInModel[model.User](c)
    exist := ui.IUserRepository.IsExist(data.Email, "EMAIL")

    if exist {
        return model.User{}, core.NewError(500, fmt.Sprintf("User [%s] already exist in the database", data.Email))
    }

    user, err := ui.IUserRepository.Create(data)
    return user, err
}

func (ui *UserInteractor) GetAll() (model.Users, *core.Error) {
    users, err := ui.IUserRepository.GetAll()
    return users, err
}

func (ui *UserInteractor) GetOne(c *gin.Context) (model.User, *core.Error) {
    data, err := utils.InjectBodyInModel[model.User](c)
    if err != nil {
        return model.User{}, err
    }

    user, err := ui.IUserRepository.GetOneById(data.Id)
    return user, err
}