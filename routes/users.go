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
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		newUser, err := r.core.RegisterUser(context, user)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, newUser)
	}
}

func (r *RoutesService) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var context, cancel = context.WithTimeout(context.Background(), time.Second * 100)
		defer cancel()

		var user models.Users
		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		foundUser, err := r.core.Login(context, user.Email, user.Password)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, foundUser)
	}
}

func (r *RoutesService) ListUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var context, cancel = context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()

		var filters models.UserFilter
		if err := ctx.BindJSON(&filters); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		users, err := r.core.ListUsers(context, &filters)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, user)
	}
}


func (r *RoutesService) CreateUserFollows() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var context, cancel = context.WithTimeout(context.Background(), time.Second * 30)
		defer cancel()

		var follows models.Follows
		if err := ctx.BindJSON(&follows); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		follows.UserId = ctx.GetString("id")
		createFollow, err := r.core.CreateUserFollows(context, &follows)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, createFollow)
	}
}

func (r *RoutesService) GetUserFollowers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var context, cancel = context.WithTimeout(context.Background(), time.Second * 30)
		defer cancel()

		var follow models.Follows

		follow.UserId = ctx.GetString("id")

		followers, err := r.core.GetUserFollowers(context, follow.UserId)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, followers)

	}
}

func (r *RoutesService) GetUserFollowings() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var context, cancel = context.WithTimeout(context.Background(), time.Second * 30)
		defer cancel()

		var follow models.Follows

		follow.FollowingId = ctx.GetString("id")

		followings, err := r.core.GetUserFollowings(context, follow.FollowingId)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, followings)

	}
}
