package interactor

import "gitlab.com/gym-partner1/api/gym-partner-api/usecases/repository"

type IFollowInteractor interface{}

type FollowInteractor struct {
	IFollowRepository repository.IFollowRepository
}
