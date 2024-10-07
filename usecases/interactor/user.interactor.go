package interactor

import (
    "errors"
    "fmt"
    "gitlab.com/Titouan-Esc/api_common/utils"
    "gitlab.com/gym-partner1/api/gym-partner-api/domain/model"
    "gitlab.com/gym-partner1/api/gym-partner-api/usecases/repository"
)

type UserInteractor struct {
    IUserRepository repository.IUserRepository
}

// -------------------------- CRUD ------------------------------

func (ui *UserInteractor) GetAll() (model.Users, error) {
    users, err := ui.IUserRepository.GetAll()
    return users, err
}

func (ui *UserInteractor) GetOneById(uid string) (model.User, error) {
    user, err := ui.IUserRepository.GetOneById(uid)
    return user, err
}

func (ui *UserInteractor) Create(data model.User) error {
    exist := ui.IUserRepository.IsExist(data.Email, "EMAIL")

    if exist {
        return errors.New(fmt.Sprintf("User [%s] already exist in database", data.Email))
    }

    err := ui.IUserRepository.Create(data)

    return err
}

func (ui *UserInteractor) Update(patch model.User) error {
    exist := ui.IUserRepository.IsExist(patch.Id, "ID")

    if !exist {
        return errors.New(fmt.Sprintf("User [%s] not found, or not exist in database", patch.Email))
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

func (ui *UserInteractor) Delete(uid string) error {
    exist := ui.IUserRepository.IsExist(uid, "ID")

    if !exist {
        return errors.New(fmt.Sprintf("User [%s] not found or not exist in database", uid))
    }

    err := ui.IUserRepository.Delete(uid)

    return err
}