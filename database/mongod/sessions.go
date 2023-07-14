package mongod

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func (d dbStore) sessionColl() *mongo.Collection {
	return d.client.Database(d.dbName).Collection("session")
}

func (d dbStore) SetSession(payload interface{}) error {
	var ctx context.Context
	_, err := d.sessionColl().InsertOne(ctx, payload)
	if err != nil {
		return err
	}

	return nil
}

func (d dbStore) ClearSession(key string) error {
	var ctx context.Context
	_, err := d.sessionColl().DeleteOne(ctx, key)
	if err != nil {
		return err
	}

	return nil
}

func (d dbStore) GetSession(key string) ([]byte, error) {
	var ctx context.Context
	var data []byte

	if err := d.sessionColl().FindOne(ctx, key).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}
