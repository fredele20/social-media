package mongod

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/fredele20/social-media/database"
	"github.com/fredele20/social-media/models"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type dbStore struct {
	dbName string
	client *mongo.Client
}

func ConnectDB(connectionUri, databaseName string) (database.UserDatastore, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionUri))
	if err != nil {
		fmt.Println("something failed")
		logrus.WithError(err).Error(err.Error())
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		fmt.Println("failed to connect to mongodb")
		log.Fatal(err)
	}

	fmt.Println("connection to mongodb established...")
	return &dbStore{dbName: databaseName, client: client}, nil
}

func (u dbStore) userColl() *mongo.Collection {
	return u.client.Database(u.dbName).Collection("users")
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

	if err := u.userColl().FindOne(ctx, filters).Decode(&user); err != nil {
		return nil, ErrDuplicate
	}
	_, err := u.userColl().InsertOne(ctx, payload)
	if err != nil {
		return nil, err
	}

	return payload, nil
}

func (u dbStore) GetUserByField(ctx context.Context, field, value string) (*models.Users, error) {
	var user models.Users
	if err := u.userColl().FindOne(ctx, bson.M{field: value}).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (u dbStore) GetUserByEmail(ctx context.Context, email string) (*models.Users, error) {
	return u.GetUserByField(ctx, "email", email)
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

	cursor, err := u.userColl().Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	if cursor.All(ctx, &users); err != nil {
		fmt.Println(err)
		return nil, err
	}

	count, err := u.userColl().CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &models.ListUsers{
		Users: users,
		Count: count,
	}, nil
}

var ErrDuplicate = errors.New("error, duplicate user")
