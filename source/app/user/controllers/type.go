package controllers

import (
	"github.com/fredele20/social-media/source/app/user/services"
	"github.com/gin-gonic/gin"
)

type controller struct {
	userService services.ServiceInterface
}

type ControllerInterface interface {
	UpdatedUserStatus(ctx *gin.Context)

	FollowUser(ctx *gin.Context)
	GetUserFollowsDetails(ctx *gin.Context)
}
