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
	return u.client.Database(u.dbName).Collection("follows")
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

	_, err := u.CreateOne(ctx, filters, user, payload)
	if err != nil {
		if err != nil {
			return nil, err
		}
	}

	// if err := u.user().FindOne(ctx, filters).Decode(&user); err == nil {
	// 	return nil, ErrDuplicate
	// }
	// _, err := u.user().InsertOne(ctx, payload)

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
	return u.GetUserByField(ctx, "id", id)
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

	if err := cursor.All(ctx, &users); err != nil {
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

	filters := bson.M{
		"$and": []bson.M{
			{
				"followingid": payload.FollowingId,
			},
			{
				"userid": payload.UserId,
			},
		},
	}

	var user models.Users

	if err := d.followers().FindOne(ctx, filters).Decode(&user); err == nil {
		return nil, ErrDuplicateFollower
	}

	// var user models.Users
	_, err := d.followers().InsertOne(ctx, payload)
	if err != nil {
		return nil, err
	}

	return payload, nil
}


// TODO: Coming back to this function
func (d dbStore) CreateUserFollows(ctx context.Context, payload *models.Follows) (*models.NewFollows, error) {

	var follow models.NewFollows
	userId := payload.UserId
	update := bson.M{
		"$push": bson.M{
			"following": payload.FollowingId,
		},
	}
	
	if err := d.followers().FindOne(ctx, bson.M{"userid": userId}).Decode(&follow); err == nil {
		if err := d.followers().FindOne(ctx, bson.M{"following": payload.FollowingId}).Decode(&follow); err == nil {
			return nil, ErrDuplicateFollower
		}
		_, err := d.followers().UpdateOne(ctx, bson.M{"userid": userId}, update)
		if err != nil {
			return nil, err
		}

		update = bson.M{
			"$push": bson.M{
				"followers": userId,
			},
		}

		if err := d.followers().FindOne(ctx, bson.M{"userid": payload.FollowingId}).Decode(&follow); err == nil {
			if err := d.followers().FindOne(ctx, bson.M{"following": payload.FollowingId}).Decode(&follow); err == nil {
				return nil, ErrDuplicateFollower
			}
			_, err := d.followers().UpdateOne(ctx, bson.M{"userid": payload.FollowingId}, update)
			if err != nil {
				return nil, err
			}
			return &follow, nil
		}

		follow.UserId = payload.FollowingId
		follow.Following = make([]string, 0)
		follow.Followers = make([]string, 0)
		follow.Followers = append(follow.Followers, payload.UserId)
		_, err = d.followers().InsertOne(ctx, follow)
		if err != nil {
			return nil, err
		}

		return &follow, nil
	}

	follow.Following = make([]string, 0)
	follow.Followers = make([]string, 0)
	follow.Following = append(follow.Following, payload.FollowingId)
	follow.UserId = userId
	_, err := d.followers().InsertOne(ctx, follow)
	if err != nil {
		return nil, err
	}

	update = bson.M{
		"$push": bson.M{
			"followers": userId,
		},
	}

	if err := d.followers().FindOne(ctx, bson.M{"userid": payload.FollowingId}).Decode(&follow); err == nil {
		if err := d.followers().FindOne(ctx, bson.M{"following": payload.FollowingId}).Decode(&follow); err == nil {
			return nil, ErrDuplicateFollower
		}
		_, err := d.followers().UpdateOne(ctx, bson.M{"userid": payload.FollowingId}, update)
		if err != nil {
			return nil, err
		}
		return &follow, nil
	}

	follow.UserId = payload.FollowingId
	follow.Following = make([]string, 0)
	follow.Followers = make([]string, 0)
	follow.Followers = append(follow.Followers, payload.UserId)
	_, err = d.followers().InsertOne(ctx, follow)
	if err != nil {
		return nil, err
	}

	return &follow, nil
}

func (d dbStore) GetUserFollowers(ctx context.Context, userId string) (*models.ListFollowers, error) {

	filter := bson.M{"userid": userId}

	var follows *models.NewFollows
	err := d.followers().FindOne(ctx, filter).Decode(&follows)
	if err != nil {
		return nil, err
	}

	count := len(follows.Followers)

	return &models.ListFollowers{
		Followers: follows.Followers,
		Count: int64(count),
	}, nil

}

func (d dbStore)	GetUserFollowings(ctx context.Context, userId string) (*models.ListFollowings, error) {

	filter := bson.M{"userid": userId}

	var follow *models.NewFollows
	err := d.followers().FindOne(ctx, filter).Decode(&follow)
	if err != nil {
		return nil, err
	}

	count := len(follow.Following)

	return &models.ListFollowings{
		Followings: follow.Following,
		Count: int64(count),
	}, nil

}

var (
	ErrDuplicate         = errors.New("error, duplicate user")
	ErrDuplicateFollower = errors.New("error, you can not follow a user more than once")
)
