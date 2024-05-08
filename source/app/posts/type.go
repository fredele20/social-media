package posts

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type postRespository struct {
	DB *mongo.Database
}

type RespositoryInterface interface {
	CreatePost(ctx context.Context, payload *Post) (Post, error)
	LikeOrUnlikePost(ctx context.Context, postId, userId string) error
	Comment(ctx context.Context, payload *Comment) error
	DeleteComment(ctx context.Context, commentId, postId, userId string) error
}

type postService struct {
	postRepo RespositoryInterface
}

type ServiceInterface interface {
	CreatePost(ctx context.Context, payload CreatePostRequest, userId string) (Post, error)
	LikeOrUnlikePost(ctx context.Context, postId, userId string) error
	Comment(ctx context.Context, userId string, data CommentRequest) error
	DeleteComment(ctx context.Context, commentId, postId, userId string) error
}

type postController struct {
	postService ServiceInterface
}

type ControllerInterface interface {
	Create(ctx *gin.Context)
	LikeOrUnlike(ctx *gin.Context)
	Comment(ctx *gin.Context)
}
