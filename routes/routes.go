package routes

import (
	"github.com/fredele20/social-media/core"
	"github.com/sirupsen/logrus"
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
