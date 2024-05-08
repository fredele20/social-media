package posts

import (
	"context"
	"fmt"

	"github.com/fredele20/social-media/source/utils"
	"github.com/jinzhu/copier"
)

func New(postRepo RespositoryInterface) ServiceInterface {
	return &postService{
		postRepo: postRepo,
	}
}

func (p *postService) CreatePost(ctx context.Context, payload CreatePostRequest, userId string) (Post, error) {
	var response Post
	fmt.Println("User ID: ", userId)

	createdAt, _ := utils.ConvertUtcToNigerianTime(utils.CurrentTime())
	updatedAt, _ := utils.ConvertUtcToNigerianTime(utils.CurrentTime())

	post := Post{
		Id:        utils.GenerateId(),
		UserId:    userId,
		Content:   payload.Content,
		ImageURL:  payload.ImageURL,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	post, err := p.postRepo.CreatePost(ctx, &post)
	if err != nil {
		return post, err
	}

	copier.Copy(&response, post)

	return response, nil
}

func (p *postService) LikeOrUnlikePost(ctx context.Context, postId, userId string) error {
	if err := p.postRepo.LikeOrUnlikePost(ctx, postId, userId); err != nil {
		return err
	}
	return nil
}

func (p *postService) Comment(ctx context.Context, userId string, data CommentRequest) error {
	createdAt, _ := utils.ConvertUtcToNigerianTime(utils.CurrentTime())
	updatedAt, _ := utils.ConvertUtcToNigerianTime(utils.CurrentTime())

	comment := Comment{
		Id:        utils.GenerateId(),
		PostId:    data.PostId,
		UserId:    userId,
		Content:   data.Content,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
	if err := p.postRepo.Comment(ctx, &comment); err != nil {
		return err
	}

	return nil
}

func (p *postService) DeleteComment(ctx context.Context, commentId, postId, userId string) error {
	return nil
}
