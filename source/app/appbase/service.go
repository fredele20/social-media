package appbase

import (
	"github.com/fredele20/social-media/source/app/auth"
	userS "github.com/fredele20/social-media/source/app/user/services"
)

func (b *base) WithTokenService() auth.TokenServiceInterface {
	return auth.NewTokenService(b.conf)
}

func (b *base) WithAuthenticationService() auth.AuthServiceInterface {
	return auth.NewService(
		b.WithUserRepository(),
		b.WithUserFollowsRepository(),
		b.WithTokenService(),
	)
}

func (b *base) WithUserService() userS.ServiceInterface {
	return userS.NewService(b.WithUserRepository(), b.WithUserFollowsRepository())
}

// func (b *base) WithPostService() posts.ServiceInterface {
// 	return posts.New(b.WithPostRepository())
// }
