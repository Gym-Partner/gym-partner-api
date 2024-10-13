package interactor

import (
    "errors"
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

    user, err := ui.IUserRepository.Create(data)
    return user, err
}