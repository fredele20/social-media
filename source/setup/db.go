package setup

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// var Client *mongo.Client
var DB *mongo.Database

// type dbCollection struct {
// 	client *mongo.Client
// }

func ConnectDB() {
	var ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(Conf.DatabaseURL))
	if err != nil {
		log.Fatalf("failed to connect to mongodb: %v", err)
	}

	db := client.Database(Conf.DatabaseName)
	fmt.Println("Connection to mongodb successful...")

	DB = db
	// Client = client
}

// type Collections interface {
// 	Collection(collectionName string) *mongo.Collection
// }

// func NewCollections(client *mongo.Client) Collections {
// 	return &dbCollection{
// 		client: client,
// 	}
// }

// func (d *dbCollection) Collection(collectionName string) *mongo.Collection {
// 	return d.client.Database(Conf.DatabaseName).Collection(collectionName)
// }
