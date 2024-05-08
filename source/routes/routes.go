package routes

import (
	"fmt"

	"github.com/fredele20/social-media/source/app/appbase"
	"github.com/fredele20/social-media/source/app/middleware"
	"github.com/fredele20/social-media/source/setup"
	"github.com/gin-gonic/gin"
)

func RouteHandlers(r *gin.Engine) {

	app := appbase.New(*setup.Conf, setup.DB).LoadControllers()

	fmt.Println("App: ", app)

	v1 := r.Group("api/v1")

	authRequired := v1.Group("")

	authRequired.Use(middleware.AuthMiddleware)

	v1.POST("/user/signup", app.AuthC.SignUp)
	v1.POST("/user/login", app.AuthC.Login)
	v1.PUT("/user/update-status", app.UserC.UpdatedUserStatus)
	authRequired.POST("/user/follow", app.UserC.FollowUser)
	authRequired.GET("/user/follow-details", app.UserC.GetUserFollowsDetails)

	// authRequired.POST("/posts", app.PostC.Create)
	// authRequired.POST("/posts/like-unlike", app.PostC.LikeOrUnlike)
	// authRequired.POST("/posts/comment", app.PostC.Comment)
}
