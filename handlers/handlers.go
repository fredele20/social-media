package handlers

import (
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
	incomingRoutes.POST("/users/login", h.routes.Login())
	incomingRoutes.GET("/users/", h.routes.ListUsers())
	incomingRoutes.GET("/users/:id", h.routes.GetUserById())
	incomingRoutes.POST("/users/follow", h.routes.CreateUserFollower())
}