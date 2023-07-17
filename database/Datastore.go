package database

import (
	"context"

	"github.com/fredele20/social-media/models"
)

type Datastore interface {
	// User implementation
	CreateUser(ctx context.Context, payload *models.Users) (*models.Users, error)
	ListUsers(ctx context.Context, filters *models.UserFilter) (*models.ListUsers, error)
	GetUserByField(ctx context.Context, field, value string) (*models.Users, error)
	GetUserByEmail(ctx context.Context, email string) (*models.Users, error)
	GetUserById(ctx context.Context, id string) (*models.Users, error)
	CreateUserFollower(ctx context.Context, payload *models.Follows) (*models.Follows, error)
	CreateNewUserFollower(ctx context.Context, payload *models.Follows) (*models.NewFollows, error)
	GetUserFollowers(ctx context.Context, userId string) (*models.ListFollowers, error)
	GetUserFollowings(ctx context.Context, followingId string) (*models.ListFollowings, error)
	ListFollowingUsersPosts(ctx context.Context, userId string) (*models.ListPosts, error)

	// Post implementation
	CreatePost(ctx context.Context, payload *models.Posts) (*models.Posts, error)
	GetPostByField(ctx context.Context, field, value string) (*models.Posts, error)
	GetPostById(ctx context.Context, id string) (*models.Posts, error)
	GetPostByContent(ctx context.Context, content string) (*models.ListPosts, error)
	GetPostsByUserId(ctx context.Context, userId string) (*models.Posts, error)
	ListPosts(ctx context.Context, filter *models.PostFilters) (*models.ListPosts, error)
	AddLike(ctx context.Context, userId string) (*models.Posts, error)
	AddComment(ctx context.Context, comment *models.Comments) (*models.Comments, error)

	// User sessions
	SetSession(payload interface{}) error
	ClearSession(key string) error
	GetSession(key string) ([]byte, error)
}
