package repository_test

import (
	"context"
	"testing"

	"github.com/fredele20/social-media/source/app/user"
	"github.com/fredele20/social-media/source/app/user/models"
	"github.com/fredele20/social-media/source/utils"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestUser(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	collection := mt.Coll

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}, {Key: "phone", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err := collection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		mt.Fatal()
	}

	mt.Run("Test Create Success", func(mt *mtest.T) {

		userRepo := user.NewUserRepository(*mt.DB)
		mt.AddMockResponses(mtest.CreateSuccessResponse())
		user, _ := userRepo.CreateUser(context.TODO(), &models.User{
			Id:        utils.GenerateId(),
			Firstname: "Victor",
			Lastname:  "Ishola",
			Email:     "test@gmail.com",
			Phone:     "09031599595",
		})

		// assert.Nil(t, err)
		assert.NotEmpty(t, user)
	})

	mt.Run("Test Create Failure - Duplicate", func(mt *mtest.T) {

		userRepo := user.NewUserRepository(*mt.DB)
		mt.AddMockResponses(mtest.CreateSuccessResponse())
		user, _ := userRepo.CreateUser(context.TODO(), &models.User{
			Id:        utils.GenerateId(),
			Firstname: "Victor",
			Lastname:  "Ishola",
			Email:     "test@gmail.com",
			Phone:     "09031599595",
		})

		// assert.Nil(t, err)
		assert.NotEmpty(t, user)
	})

	mt.Run("Test Create Failure", func(mt *mtest.T) {
		userRepo := user.NewUserRepository(*mt.DB)
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: -1},
		})

		_, err := userRepo.CreateUser(context.TODO(), &models.User{
			Id:        utils.GenerateId(),
			Firstname: "Victor",
			Lastname:  "Ishola",
			Email:     "test@gmail.com",
			Phone:     "09031599595",
		})

		assert.NotNil(t, err)
	})
}
