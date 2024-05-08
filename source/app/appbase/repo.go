package appbase

import (
	"github.com/fredele20/social-media/source/app/user"
)

// func (b *base) WithDBCollection() setup.Collections {
// 	return setup.NewCollections(b.client)
// }

func (b *base) WithUserRepository() user.RepositoryInterface {
	return user.NewUserRepository(*b.db)
}

func (b *base) WithUserFollowsRepository() user.UserFollowsRepositoryInterface {
	return user.NewUserFollowsRepository(*b.db)
}

// func (b *base) WithPostRepository() posts.RespositoryInterface {
// 	return posts.NewRepositoryInterface(b.WithDBCollection())
// }
