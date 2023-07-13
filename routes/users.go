package routes

import (
	"context"
	"net/http"
	"time"

	"github.com/fredele20/social-media/models"
	"github.com/gin-gonic/gin"
)

func reuseContext(duration time.Duration) context.Context {
	var context, cancel = context.WithTimeout(context.Background(), duration)
	defer cancel()

	return context
}

func (r *RoutesService) RegisterUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var context, cancel = context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		// reuseContext(time.Second * 30)

		var user *models.Users
		if err := ctx.BindJSON(&user); err != nil {
			// r.logger.WithError(err).Error(ErrBindingToStruct.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": ErrBindingToStruct.Error()})
			return
		}

		newUser, err := r.core.RegisterUser(context, user)
		if err != nil {
			// r.logger.WithError(err).Error(ErrCreatingUser.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": ErrCreatingUser.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, newUser)
	}
}

func (r *RoutesService) ListUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var context, cancel = context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()

		var filters models.UserFilter
		if err := ctx.BindJSON(&filters); err != nil {
			r.logger.WithError(err).Error(ErrBindingToStruct.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": ErrBindingToStruct.Error()})
			return
		}

		users, err := r.core.ListUsers(context, &filters)
		if err != nil {
			r.logger.WithError(err).Error(ErrListingData.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": ErrListingData.Error()})
			return
		}

		ctx.JSON(http.StatusOK, users)
	}
}

func (r *RoutesService) GetUserById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var context, cancel = context.WithTimeout(context.Background(), time.Second * 30)
		defer cancel()

		userId := ctx.Param("id")

		user, err := r.core.GetUserById(context, userId)
		if err != nil {
			r.logger.WithError(err).Error(ErrFailedToGetUserById)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": ErrFailedToGetUserById})
			return
		}

		ctx.JSON(http.StatusOK, user)
	}
}

func (r *RoutesService) CreateUserFollower() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var context, cancel = context.WithTimeout(context.Background(), time.Second * 30)
		defer cancel()

		var follower models.Follows

		if err := ctx.BindJSON(&follower); err != nil {
			r.logger.WithError(err).Error(ErrBindingToStruct.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": ErrBindingToStruct.Error()})
			return
		}

		// followerId := ctx.GetString("id")
		// follower.FollowersId = append(follower.FollowersId, followerId)
		follow, err := r.core.CreateUserFollower(context, &follower)
		if err != nil {
			r.logger.WithError(err).Error(ErrFollowUserFailed)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": ErrFollowUserFailed})
			return
		}

		ctx.JSON(http.StatusOK, follow)
	}
}
