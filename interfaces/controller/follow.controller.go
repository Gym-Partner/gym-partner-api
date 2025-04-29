package controller

import (
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/interfaces/repository"
	"gitlab.com/gym-partner1/api/gym-partner-api/usecases/interactor"
)

type FollowController struct {
	IFollowInteractor interactor.IFollowInteractor
	Log               *core.Log
}

func NewFollowController(db *core.Database) *FollowController {
	return &FollowController{
		IFollowInteractor: &interactor.FollowInteractor{
			IFollowRepository: repository.FollowRepository{
				DB:  db.Handler,
				Log: db.Logger,
			},
		},
		Log: db.Logger,
	}
}
