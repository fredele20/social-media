package mongod

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/fredele20/social-media/database"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type dbStore struct {
	dbName string
	client *mongo.Client
}

func ConnectDB(connectionUri, databaseName string) (database.Datastore, error) {
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
func (u dbStore) CreateOne(ctx context.Context, filter, data, payload interface{}) (interface{}, error) {
	
	if err := u.user().FindOne(ctx, filter).Decode(&data); err == nil {
		return nil, ErrDuplicate
	}
	_, err := u.user().InsertOne(ctx, payload)
	if err != nil {
		return nil, err
	}

	return payload, nil
}
