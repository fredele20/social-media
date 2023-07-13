package mongod

import (
	"context"

	"github.com/fredele20/social-media/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (d dbStore) postColl() *mongo.Collection {
	return d.client.Database(d.dbName).Collection("posts")
}


func (d dbStore) CreatePost(ctx context.Context, payload *models.Posts) (*models.Posts, error) {
	_, err := d.postColl().InsertOne(ctx, payload)
	if err != nil {
		return nil, err
	}

	return payload, nil
}

func (d dbStore) GetPostByField(ctx context.Context, field, value string) (*models.Posts, error) {
	var post models.Posts
	if err := d.postColl().FindOne(ctx, bson.M{field: value}).Decode(&post); err != nil {
		return nil, err
	}

	return &post, nil
}

func (d dbStore) GetPostById(ctx context.Context, id string) (*models.Posts, error) {
	return d.GetPostByField(ctx, "id", id)
}

func (d dbStore) GetPostByContent(ctx context.Context, content string) (*models.ListPosts, error) {
	cursor, err := d.postColl().Find(ctx, content)
	if err != nil {
		return nil, err
	}

	var posts []*models.Posts

	filter := bson.M{}

	if err := cursor.All(ctx, &posts); err != nil {
		return nil, err
	}

	count, err := d.postColl().CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &models.ListPosts{
		Count: count,
		Posts: posts,
	}, nil
}

func (d dbStore) ListPosts(ctx context.Context, filters *models.PostFilters) (*models.ListPosts, error) {

	opts := options.Find()
	
	if filters.Limit != 0 {
		opts.SetLimit(filters.Limit)
	}

	filter := bson.M{}

	var posts []*models.Posts

	cursor, err := d.postColl().Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &posts); err != nil {
		return nil, err
	}

	count, err := d.postColl().CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &models.ListPosts{
		Posts: posts,
		Count: count,
	}, nil
}

func (d dbStore) AddLike(ctx context.Context, userId string) (*models.Posts, error) {
	likes := []string{}
	likes = append(likes, userId)
	update := bson.M{
		"$set": bson.M{
			"likes": likes,
		},
	}
	var post models.Posts
	if err := d.postColl().FindOneAndUpdate(ctx, "postId", update).Decode(&post); err != nil {
		return nil, err
	}

	return &post, nil
}

func (d dbStore) AddComment(ctx context.Context, comment string) (*models.Posts, error) {
	comments := []string{}
	comments = append(comments, comment)
	update := bson.M{
		"$set": bson.M{
			"comments": comments,
		},
	}

	var post models.Posts
	if err := d.postColl().FindOneAndUpdate(ctx, "postId", update).Decode(&post); err != nil {
		return nil, err
	}

	return &post, nil
}