package interactor

import (
    "errors"
    "fmt"
    "github.com/gin-gonic/gin"
    "gitlab.com/gym-partner1/api/gym-partner-api/domain/model"
    "gitlab.com/gym-partner1/api/gym-partner-api/usecases/repository"
)

type UserInteractor struct {
    IUserRepository repository.IUserRepository
}

// -------------------------- CRUD ------------------------------

func (ui *UserInteractor) Create(c *gin.Context) (model.User, error) {
    var data model.User

    if err := c.ShouldBind(&data); err != nil {
        return model.User{}, errors.New("Error to parse body " + err.Error())
    }

    exist := ui.IUserRepository.IsExist(data.Email, "EMAIL")

    if exist {
        return model.User{}, errors.New(fmt.Sprintf("User [%s] already exist in the database", data.Email))
    }

    user, err := ui.IUserRepository.Create(data)
    return user, err
}