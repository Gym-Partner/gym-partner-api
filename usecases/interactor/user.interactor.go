package interactor

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "gitlab.com/gym-partner1/api/gym-partner-api/core"
    "gitlab.com/gym-partner1/api/gym-partner-api/domain/model"
    "gitlab.com/gym-partner1/api/gym-partner-api/usecases/repository"
)

type UserInteractor struct {
    IUserRepository repository.IUserRepository
}

// -------------------------- CRUD ------------------------------

func (ui *UserInteractor) Create(c *gin.Context) (model.User, *core.Error) {
    var data model.User

    if err := c.ShouldBind(&data); err != nil {
        return model.User{}, core.NewError(500, "Error to bind request body to model.User", err)
    }

    exist := ui.IUserRepository.IsExist(data.Email, "EMAIL")

    if exist {
        return model.User{}, core.NewError(500, fmt.Sprintf("User [%s] already exist in the database", data.Email))
    }

    user, err := ui.IUserRepository.Create(data)
    return user, err
}