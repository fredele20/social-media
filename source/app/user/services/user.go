package services

import (
	"context"

	userT "github.com/fredele20/social-media/source/app/user"
	"github.com/fredele20/social-media/source/app/user/models"
)

func NewService(
	userRepo userT.RepositoryInterface,
	followRepo userT.UserFollowsRepositoryInterface,
) ServiceInterface {
	return &userService{
		userRepo:   userRepo,
		followRepo: followRepo,
	}
}

func (u *userService) UpdatedUserStatus(ctx context.Context, userId string) error {
	user, userErr := u.userRepo.FindUserById(ctx, userId)
	if userErr != nil {
		return userErr
	}

	var status string
	if user.Status == models.ACTIVATED {
		status = string(models.DEACTIVATED)
	} else if user.Status == models.DEACTIVATED {
		status = string(models.ACTIVATED)
	}

	user.Status = models.Status(status)

	if updateErr := u.userRepo.UpdateUser(ctx, &user); updateErr != nil {
		return updateErr
	}

	return nil
}

