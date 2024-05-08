package controllers

import (
	"github.com/fredele20/social-media/source/app/user/services"
	"github.com/fredele20/social-media/source/utils"
	"github.com/gin-gonic/gin"
)

func NewController(userService services.ServiceInterface) ControllerInterface {
	return &controller{
		userService: userService,
	}
}

func (u *controller) UpdatedUserStatus(ctx *gin.Context) {
	userId := "65e6efb7a82ed8f77c471afa"

	if updateErr := u.userService.UpdatedUserStatus(ctx, userId); updateErr != nil {
		utils.JsonErrorResponse(ctx, updateErr.Error())
	}

	utils.JsonSuccessResponse(ctx, nil, "user status updated successfully")
}
