package services

import (
	"context"

	"github.com/fredele20/social-media/source/app/user"
)

func (f *userService) FollowUser(ctx context.Context, userId, followId string) error {
	if err := f.followRepo.FollowUser(ctx, userId, followId); err != nil {
		return err
	}
	return nil
}

func (f *userService) GetUserFollowsDetails(ctx context.Context, userId string) (user.UserFollowsDetails, error) {
	var response user.UserFollowsDetails

	data, err := f.followRepo.GetUserFollowsDetails(ctx, userId)
	if err != nil {
		return response, err
	}

	return data, nil
}
