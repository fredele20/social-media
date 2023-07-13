package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Follows struct {
	Id          primitive.ObjectID `json:"id"`
	UserId      string             `json:"userId"`
	FollowersId []string           `json:"followerId"`
}
