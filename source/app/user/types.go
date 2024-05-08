package user

import (
	"context"

	"github.com/fredele20/social-media/source/app/user/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	DB mongo.Database
}

//go:generate mockgen -source=type.go -destination=../..tests/user/repository/repo.go -package=mocks
type RepositoryInterface interface {
	// ListUsers(ctx context.Context, filters UserFilter) (ListUsers, error)
	CreateUser(ctx context.Context, user *models.User) (models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
	FindUserByEmailOrPhone(ctx context.Context, loginId string) (models.User, error)
	FindUserByField(ctx context.Context, field, value string) (models.User, error)
	FindUserByEmail(ctx context.Context, email string) (models.User, error)
	FindUserById(ctx context.Context, id string) (models.User, error)
}

type userFollowsRepository struct {
	DB mongo.Database
}

type UserFollowsRepositoryInterface interface {
	CreateUserFollows(ctx context.Context, userId string) error
	FollowUser(ctx context.Context, userId, followId string) error
	GetUserFollowsDetails(ctx context.Context, userId string) (UserFollowsDetails, error)
}
