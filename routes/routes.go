package routes

import (
	"errors"

	"github.com/fredele20/social-media/core"
	"github.com/sirupsen/logrus"
)

var (
	ErrBindingToStruct     = errors.New("failed to bind data to struct")
	ErrCreatingUser        = errors.New("failed to create user")
	ErrListingData         = errors.New("failed to list data")
	ErrFailedToGetUserById = errors.New("failed to get user by Id")
	ErrFollowUserFailed    = errors.New("failed to follow user")
	ErrLoginFailed         = errors.New("failed to login user from route")
)

type RoutesService struct {
	core   *core.CoreService
	logger logrus.Logger
}

func NewRoutesService(core *core.CoreService) *RoutesService {
	return &RoutesService{
		core: core,
	}
}
