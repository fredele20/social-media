package core

import (
	"context"
	"errors"

	"github.com/fredele20/social-media/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrCreatePostFailed          = errors.New("failed to create a post, try again")
	ErrFailedToGetPostWithId     = errors.New("failed to get post with the given Id")
	ErrFailedToGetPostsByContent = errors.New("failed to get post(s) by the given content")
	ErrFailedToListPosts         = errors.New("failed to list posts")
	ErrFailedToAddLikeToPost     = errors.New("failed to add like to post")
	ErrFailedToAddCommentToPost  = errors.New("failed to add comment to post")
)

func (c *CoreService) CreatePost(ctx context.Context, payload *models.Posts) (*models.Posts, error) {
	generateId := primitive.NewObjectID()
	payload.Id = generateId.Hex()
	post, err := c.db.CreatePost(ctx, payload)
	if err != nil {
		c.logger.WithError(err).Error(ErrCreatePostFailed)
		return nil, err
	}

	return post, nil
}

func (c *CoreService) GetPostById(ctx context.Context, id string) (*models.Posts, error) {
	post, err := c.db.GetPostById(ctx, id)
	if err != nil {
		c.logger.WithError(err).Error(ErrFailedToGetPostWithId)
		return nil, err
	}

	return post, nil
}

func (c *CoreService) GetPostByContent(ctx context.Context, content string) (*models.ListPosts, error) {
	posts, err := c.db.GetPostByContent(ctx, content)
	if err != nil {
		c.logger.WithError(err).Error(ErrFailedToGetPostsByContent.Error())
		return nil, err
	}

	return posts, nil
}

func (c *CoreService) ListPosts(ctx context.Context, filter models.PostFilters) (*models.ListPosts, error) {
	posts, err := c.db.ListPosts(ctx, &filter)
	if err != nil {
		c.logger.WithError(err).Error(ErrFailedToListPosts)
		return nil, err
	}

	return posts, nil
}

// func (c *CoreService) ListFollowingUsersPosts(ctx context.Context) (*models.ListPosts, error) {}

func (c *CoreService) AddLike(ctx context.Context, userId string) (*models.Posts, error) {
	post, err := c.db.AddLike(ctx, userId)
	if err != nil {
		c.logger.WithError(err).Error(ErrFailedToAddLikeToPost)
		return nil, err
	}

	return post, err
}

func (c *CoreService) AddComment(ctx context.Context, payload *models.Comments) (*models.Comments, error) {
	generateId := primitive.NewObjectID()
	payload.Id = generateId.Hex()
	post, err := c.db.AddComment(ctx, payload)
	if err != nil {
		c.logger.WithError(err).Error(ErrFailedToAddCommentToPost)
		return nil, err
	}

	return post, nil
}
