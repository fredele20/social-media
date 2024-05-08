package posts

type CreatePostRequest struct {
	Content  string `json:"content" binding:"required"`
	ImageURL string `json:"image_url"`
}

type LikePostRequest struct {
	PostId string `json:"post_id" binding:"required"`
}

type CommentRequest struct {
	PostId   string `json:"post_id" binding:"required"`
	Content  string `json:"content" binding:"required"`
	ImageURL string `json:"imageUrl"`
}
