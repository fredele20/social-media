package handlers

import (
	"github.com/fredele20/social-media/middleware"
	"github.com/fredele20/social-media/routes"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	routes *routes.RoutesService
}

func NewHandler(r *routes.RoutesService) *Handler {
	return &Handler{
		routes: r,
	}
}

func UserHandler(incomingRoutes *gin.Engine, h Handler) {
	incomingRoutes.POST("/users/register/", h.routes.RegisterUser())
	incomingRoutes.POST("/users/login/", h.routes.Login())
	incomingRoutes.GET("/users/", h.routes.ListUsers())
	incomingRoutes.GET("/users/:id/", h.routes.GetUserById())

	incomingRoutes.Use(middleware.Authenticate())
	// incomingRoutes.POST("/users/follow/", h.routes.CreateUserFollower())
	incomingRoutes.POST("/users/follow/", h.routes.CreateNewUserFollower())
	incomingRoutes.GET("/users/followers/", h.routes.GetUserFollowers())
	incomingRoutes.GET("/users/followings/", h.routes.GetUserFollowings())
	incomingRoutes.POST("/posts/", h.routes.CreatePost())
	incomingRoutes.POST("/posts/comments/", h.routes.AddComment())
	incomingRoutes.GET("/posts/", h.routes.ListPosts())
}
