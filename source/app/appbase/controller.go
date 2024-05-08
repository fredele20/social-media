package appbase

import (
	"github.com/fredele20/social-media/source/app/auth"
	userC "github.com/fredele20/social-media/source/app/user/controllers"
)

func (b *base) WithAuthController() auth.ControllerInterface {
	return auth.NewAuthContoller(b.WithAuthenticationService())
}

func (b *base) WithUserController() userC.ControllerInterface {
	return userC.NewController(b.WithUserService())
}

// func (b *base) WithPostController() posts.ControllerInterface {
// 	return posts.NewController(b.WithPostService())
// }
