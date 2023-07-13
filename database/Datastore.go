package database

import (
	"context"

	"github.com/fredele20/social-media/models"
)

type UserDatastore interface {
	CreateUser(ctx context.Context, payload *models.Users) (*models.Users, error)
	ListUsers(ctx context.Context, filters *models.UserFilter) (*models.ListUsers, error)
	GetUserByField(ctx context.Context, field, value string) (*models.Users, error)
	GetUserByEmail(ctx context.Context, email string) (*models.Users, error)
}
