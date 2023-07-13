package mongod

import (
	"context"
	"errors"
	"fmt"

	"github.com/fredele20/social-media/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (u dbStore) user() *mongo.Collection {
	return u.client.Database(u.dbName).Collection("users")
}

func (u dbStore) followers() *mongo.Collection {
	return u.client.Database(u.dbName).Collection("followers")
}

func (u dbStore) CreateUser(ctx context.Context, payload *models.Users) (*models.Users, error) {
	filters := bson.M{
		"$or": []bson.M{
			{
				"email": payload.Email,
			},
			{
				"phone": payload.Phone,
			},
		},
	}

	var user models.Users

	if err := u.user().FindOne(ctx, filters).Decode(&user); err != nil {
		return nil, ErrDuplicate
	}
	_, err := u.user().InsertOne(ctx, payload)
	if err != nil {
		return nil, err
	}

	return payload, nil
}

func (u dbStore) GetUserByField(ctx context.Context, field, value string) (*models.Users, error) {
	var user models.Users
	if err := u.user().FindOne(ctx, bson.M{field: value}).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (u dbStore) GetUserByEmail(ctx context.Context, email string) (*models.Users, error) {
	return u.GetUserByField(ctx, "email", email)
}

func (u dbStore) GetUserById(ctx context.Context, id string) (*models.Users, error) {
	return u.GetUserByField(ctx, "userId", id)
}

func (u dbStore) ListUsers(ctx context.Context, filters *models.UserFilter) (*models.ListUsers, error) {
	opts := options.Find()
	opts.SetProjection(bson.M{
		"password": false,
		"token":    false,
	})

	if filters.Limit != 0 {
		opts.SetLimit(filters.Limit)
	}

	filter := bson.M{}

	var users []*models.Users

	cursor, err := u.user().Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	if cursor.All(ctx, &users); err != nil {
		fmt.Println(err)
		return nil, err
	}

	count, err := u.user().CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &models.ListUsers{
		Users: users,
		Count: count,
	}, nil
}

func (d dbStore) CreateUserFollower(ctx context.Context, payload *models.Follows) (*models.Follows, error) {
	
	// var user models.Users
	_, err := d.followers().InsertOne(ctx, payload)
	if err != nil {
		return nil, err
	}

	return payload, nil
}

var ErrDuplicate = errors.New("error, duplicate user")