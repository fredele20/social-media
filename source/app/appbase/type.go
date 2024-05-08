package appbase

import (
	"github.com/fredele20/social-media/source/app/auth"
	"github.com/fredele20/social-media/source/app/posts"
	userC "github.com/fredele20/social-media/source/app/user/controllers"
	"github.com/fredele20/social-media/source/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type base struct {
	conf config.Config
	db   *mongo.Database
}

type baseControllers struct {
	AuthC auth.ControllerInterface
	UserC userC.ControllerInterface
	PostC posts.ControllerInterface
}
