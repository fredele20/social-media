package user

import (
	"context"

	"github.com/fredele20/social-media/source/app/user/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserFollowsRepository(db mongo.Database) UserFollowsRepositoryInterface {
	return &userFollowsRepository{
		DB: db,
	}
}

func (u *userFollowsRepository) CreateUserFollows(ctx context.Context, userId string) error {

	follow := models.Follows{
		UserId:    userId,
		Followers: make([]string, 0),
		Following: make([]string, 0),
	}
	_, insertErr := u.collection().InsertOne(ctx, follow)
	if insertErr != nil {
		return insertErr
	}

	return nil
}

func (u *userFollowsRepository) FollowUser(ctx context.Context, userId, followId string) error {
	var follows models.Follows

	findFilter := bson.M{"user_id": userId, "following": bson.M{"$in": []string{followId}}}
	err := u.collection().FindOne(context.Background(), findFilter).Decode(&follows)
	if err == nil {
		return ErrDuplicateFollows
	}

	filter := bson.D{primitive.E{Key: "user_id", Value: userId}}
	update := bson.D{{Key: "$push", Value: bson.D{primitive.E{Key: "following", Value: followId}}}}

	_, updateErr := u.collection().UpdateOne(ctx, filter, update)
	if updateErr != nil {
		return updateErr
	}

	filter2 := bson.D{primitive.E{Key: "user_id", Value: followId}}
	update2 := bson.D{{Key: "$push", Value: bson.D{primitive.E{Key: "followers", Value: userId}}}}

	_, updateErr2 := u.collection().UpdateOne(ctx, filter2, update2)
	if updateErr2 != nil {
		return updateErr2
	}

	return nil
}

func (u *userFollowsRepository) GetUserFollowsDetails(ctx context.Context, userId string) (UserFollowsDetails, error) {
	var userFollowsDetails UserFollowsDetails
	var follows models.Follows

	if err := u.collection().FindOne(ctx, bson.M{"user_id": userId}).Decode(&follows); err != nil {
		return userFollowsDetails, err
	}

	userFollowsDetails.UserId = follows.UserId
	userFollowsDetails.Followers.Count = len(follows.Followers)
	userFollowsDetails.Following.Count = len(follows.Following)
	userFollowsDetails.Followers.Data = follows.Followers
	userFollowsDetails.Following.Data = follows.Following

	return userFollowsDetails, nil
}

func (u *userFollowsRepository) collection() *mongo.Collection {
	return u.DB.Collection("follows")
}
