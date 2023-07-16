package routes

import (
	"context"
	"net/http"
	"time"

	"github.com/fredele20/social-media/models"
	"github.com/gin-gonic/gin"
)


func (r *RoutesService) CreatePost() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var context, cancel = context.WithTimeout(context.Background(), time.Second * 30)
		defer cancel()

		userId := ctx.GetString("id")

		var post models.Posts
		if err := ctx.BindJSON(&post); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		post.UserId = userId

		newPost, err := r.core.CreatePost(context, &post)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, newPost)
	}
}

func (r *RoutesService) ListPosts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var context, cancel = context.WithTimeout(context.Background(), time.Second * 30)
		defer cancel()

		var filter models.PostFilters
		if err := ctx.BindJSON(&filter); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		posts, err := r.core.ListPosts(context, filter)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, posts)
	}
}

func (r *RoutesService) AddComment() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var context, cancel = context.WithTimeout(context.Background(), time.Second * 30)
		defer cancel()

		// postId := ctx.Param("id")

		var comment models.Comments
		if err := ctx.BindJSON(&comment); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		comment.UserId = ctx.GetString("id")

		post, err := r.core.AddComment(context, &comment)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, post)
	}
}