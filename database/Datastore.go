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

	// Post implementation
	CreatePost(ctx context.Context, payload *models.Posts) (*models.Posts, error)
	GetPostByField(ctx context.Context, field, value string) (*models.Posts, error)
	GetPostById(ctx context.Context, id string) (*models.Posts, error)
	GetPostByContent(ctx context.Context, content string) (*models.ListPosts, error)
	ListPosts(ctx context.Context, filter *models.PostFilters) (*models.ListPosts, error)
	AddLike(ctx context.Context, userId string) (*models.Posts, error)
	AddComment(ctx context.Context, content string) (*models.Posts, error)
}
