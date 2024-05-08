package controllers

import (
	"github.com/fredele20/social-media/source/app/user"
	"github.com/fredele20/social-media/source/utils"
	"github.com/gin-gonic/gin"
)

func (u *controller) FollowUser(ctx *gin.Context) {
	var param user.FollowUserRequest
	if !utils.ValidParams(ctx, &param) {
		return
	}

	userId := ctx.GetString("userId")

	if err := u.userService.FollowUser(ctx, userId, param.UserId); err != nil {
		utils.JsonErrorResponse(ctx, err.Error())
		return
	}

	utils.JsonSuccessResponse(ctx, nil, "successfully followed user")
}

func (u *controller) GetUserFollowsDetails(ctx *gin.Context) {
	userId := ctx.GetString("userId")

	data, err := u.userService.GetUserFollowsDetails(ctx, userId)
	if err != nil {
		utils.JsonErrorResponse(ctx, err.Error())
		return
	}

	utils.JsonSuccessResponse(ctx, data, "successfully followed user")
}
