package appbase

import (
	"github.com/fredele20/social-media/source/config"
	"go.mongodb.org/mongo-driver/mongo"
)

func New(conf config.Config, db *mongo.Database) *base {
	return &base{
		conf: conf,
		db:   db,
	}
}

func (b *base) LoadControllers() baseControllers {
	var c baseControllers
	c.AuthC = b.WithAuthController()
	c.UserC = b.WithUserController()
	// c.PostC = b.WithPostController()

	return c
}
