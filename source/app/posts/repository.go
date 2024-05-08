package posts

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRepositoryInterface(db *mongo.Database) RespositoryInterface {
	return &postRespository{
		DB: db,
	}
}

func (p *postRespository) CreatePost(ctx context.Context, post *Post) (Post, error) {
	post.Likes = make([]string, 0)
	post.Comments = make([]Comment, 0)
	_, err := p.collection().InsertOne(ctx, post)
	if err != nil {
		return *post, err
	}

	return *post, nil
}

func (p *postRespository) LikeOrUnlikePost(ctx context.Context, postId, userId string) error {
	var post Post

	findFilter := bson.M{"id": postId, "likes": bson.M{"$in": []string{userId}}}
	err := p.collection().FindOne(context.Background(), findFilter).Decode(&post)
	if err == nil {
		return p.update(ctx, "$pull", userId, postId)
	}

	return p.update(ctx, "$push", userId, postId)
}

func (p *postRespository) Comment(ctx context.Context, payload *Comment) error {
	filter := bson.D{primitive.E{Key: "id", Value: payload.PostId}}
	update := bson.D{{Key: "$push", Value: bson.D{primitive.E{Key: "comments", Value: payload}}}}
	_, updateErr := p.collection().UpdateOne(ctx, filter, update)
	if updateErr != nil {
		return updateErr
	}

	return nil
}

func (p *postRespository) DeleteComment(ctx context.Context, commentId, postId, userId string) error {
	filter := bson.D{primitive.E{Key: "id", Value: postId}}
	update := bson.M{"$pull": bson.M{"comments": bson.M{"id": commentId}}}
	_, updateErr := p.collection().UpdateOne(ctx, filter, update)
	if updateErr != nil {
		return updateErr
	}

	return nil
}

func (p *postRespository) update(ctx context.Context, aggregateFunc, userId, postId string) error {
	filter := bson.D{primitive.E{Key: "id", Value: postId}}
	update := bson.D{{Key: aggregateFunc, Value: bson.D{primitive.E{Key: "likes", Value: userId}}}}

	_, updatedErr := p.collection().UpdateOne(ctx, filter, update)
	if updatedErr != nil {
		return updatedErr
	}
	return nil
}

func (p *postRespository) collection() *mongo.Collection {
	return p.DB.Collection("posts")
}
