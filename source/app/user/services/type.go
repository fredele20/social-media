package services

import (
	"context"

	userT "github.com/fredele20/social-media/source/app/user"
)

type userService struct {
	userRepo   userT.RepositoryInterface
	followRepo userT.UserFollowsRepositoryInterface
}

type ServiceInterface interface {
	UpdatedUserStatus(ctx context.Context, userId string) error

	FollowUser(ctx context.Context, userId, followId string) error
	GetUserFollowsDetails(ctx context.Context, userId string) (userT.UserFollowsDetails, error)
}
