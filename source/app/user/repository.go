package user

import (
	"context"
	"errors"

	"github.com/fredele20/social-media/source/app/user/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepository(db mongo.Database) RepositoryInterface {
	return &userRepository{
		DB: db,
	}
}

func (u *userRepository) CreateUser(ctx context.Context, user *models.User) (models.User, error) {
	filters := bson.M{
		"$or": []bson.M{
			{
				"email": user.Email,
			},
			{
				"phone": user.Phone,
			},
		},
	}

	if err := u.userCollection().FindOne(ctx, filters).Decode(&user); err == nil {
		return *user, errors.New("duplicate user")
	}

	_, err := u.userCollection().InsertOne(ctx, user)
	if err != nil {
		return *user, err
	}

	return *user, nil
}

func (u *userRepository) UpdateUser(ctx context.Context, user *models.User) error {
	filter := bson.M{"id": user.Id}
	update := bson.M{"$set": user}

	if err := u.userCollection().FindOneAndUpdate(ctx, filter, update).Decode(&user); err != nil {
		return err
	}

	return nil
}

func (u *userRepository) FindUserByEmailOrPhone(ctx context.Context, loginId string) (models.User, error) {
	var user models.User

	filters := bson.M{
		"$or": []bson.M{
			{
				"email": loginId,
			},
			{
				"phone": loginId,
			},
		},
	}

	if err := u.userCollection().FindOne(ctx, filters).Decode(&user); err != nil {
		return user, errors.New("user not found")
	}

	return user, nil
}

func (u *userRepository) FindUserByField(ctx context.Context, field, value string) (models.User, error) {
	var user models.User
	if err := u.userCollection().FindOne(ctx, bson.M{field: value}).Decode(&user); err != nil {
		return user, err
	}
	return user, nil
}

func (u *userRepository) FindUserByEmail(ctx context.Context, email string) (models.User, error) {
	return u.FindUserByField(ctx, "email", email)
}

func (u *userRepository) FindUserById(ctx context.Context, id string) (models.User, error) {
	return u.FindUserByField(ctx, "id", id)
}

func (u *userRepository) userCollection() *mongo.Collection {
	return u.DB.Collection("user")
}

// func (u *userRepository) ListUsers(ctx context.Context, filter UserFilter) (ListUsers, error) {}
