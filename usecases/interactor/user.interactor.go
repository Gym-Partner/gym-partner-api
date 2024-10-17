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

func (ui *UserInteractor) Create(ctx *gin.Context) (model.User, *core.Error) {
    aws, _ := ctx.Get("aws")
    data, err := utils.InjectBodyInModel[model.User](ctx)
    if err != nil {
        return model.User{}, err
    }

    exist := ui.IUserRepository.IsExist(data.Email, "EMAIL")
    if exist {
        return model.User{}, core.NewError(500, fmt.Sprintf("User [%s] already exist in the database", data.Email))
    }

    user, err := ui.IUserRepository.Create(data)

    cognito, err := aws.(*core.AWS).NewCognito()
    if err != nil {
        return user, core.NewError(500, "Failed initial AWS Service", err)
    }

    data.Id = user.Id
    if err = cognito.SignUp(data); err != nil {
        return user, core.NewError(500, "Failed create user to Cognito", err)
    }

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

func (ui *UserInteractor) GetOneByEmail(ctx *gin.Context) (model.User, *core.Error) {
    data, err := utils.InjectBodyInModel[model.User](ctx)
    if err != nil {
        return model.User{}, err
    }

    user, err := ui.IUserRepository.GetOneByEmail(data.Email)
    user.Password = data.Password
    return user, err
}

func (ui *UserInteractor) Update(ctx *gin.Context) *core.Error {
    patch, err := utils.InjectBodyInModel[model.User](ctx)
    if err != nil {
        return err
    }

    exist := ui.IUserRepository.IsExist(patch.Id, "ID")
    if !exist {
        return core.NewError(500, fmt.Sprintf("User [%s] not found, or not exist in the database", patch.Id))
    }

    target, err := ui.IUserRepository.GetOneById(patch.Id)
    if err != nil {
        return err
    }

    if err = utils.Bind(&target, patch); err != nil {
        return err
    }

    err = ui.IUserRepository.Update(target)
    return err
}

func (ui *UserInteractor) Delete(ctx *gin.Context) *core.Error {
    data, err := utils.InjectBodyInModel[model.User](ctx)
    if err != nil {
        return err
    }

    exist := ui.IUserRepository.IsExist(data.Id, "ID")
    if !exist {
        return core.NewError(500, "User not found in the database")
    }

    if err = ui.IUserRepository.Delete(data.Id); err != nil {
        return err
    }

    return nil
}

func (ui *UserInteractor) Login(ctx *gin.Context, user model.User) (string, *core.Error) {
    aws, _ := ctx.Get("aws")
    cognito, err := aws.(*core.AWS).NewCognito()
    if err != nil {
        return "", err
    }

    token, err := cognito.SignIn(user)
    if err != nil {
        return "", err
    }

    return token, nil
}