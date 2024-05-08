package posts

import (
	"github.com/fredele20/social-media/source/utils"
	"github.com/gin-gonic/gin"
)

func NewController(postService ServiceInterface) ControllerInterface {
	return &postController{
		postService: postService,
	}
}

func (p *postController) Create(ctx *gin.Context) {
	var param CreatePostRequest
	if !utils.ValidParams(ctx, &param) {
		return
	}

	userId := ctx.GetString("userId")
	post, err := p.postService.CreatePost(ctx, param, userId)
	if err != nil {
		utils.JsonErrorResponse(ctx, err.Error())
		return
	}

	utils.JsonCreatedResponse(ctx, post, "Successfully created post")
}

func (p *postController) LikeOrUnlike(ctx *gin.Context) {
	var param LikePostRequest
	if !utils.ValidParams(ctx, &param) {
		return
	}

	userId := ctx.GetString("userId")

	if err := p.postService.LikeOrUnlikePost(ctx, param.PostId, userId); err != nil {
		utils.JsonErrorResponse(ctx, err.Error())
		return
	}

	utils.JsonSuccessResponse(ctx, nil, "successfully liked a post")
}

func (p *postController) Comment(ctx *gin.Context) {
	var param CommentRequest
	if !utils.ValidParams(ctx, &param) {
		return
	}

	userId := ctx.GetString("userId")

	if err := p.postService.Comment(ctx, userId, param); err != nil {
		utils.JsonErrorResponse(ctx, err.Error())
		return
	}

	utils.JsonSuccessResponse(ctx, nil, "comment added succussfully")
}
