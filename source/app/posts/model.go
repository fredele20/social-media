package posts

import "time"

type Post struct {
	Id        string    `json:"id"`
	UserId    string    `json:"user_id"`
	Content   string    `json:"content"`
	ImageURL  string    `json:"image_url"`
	Likes     []string  `json:"likes" bson:"likes"`
	Comments  []Comment `json:"comments"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Comment struct {
	Id        string    `json:"id"`
	PostId    string    `json:"postId"`
	UserId    string    `json:"userId"`
	Content   string    `json:"content"`
	ImageURL  string    `json:"imageUrl"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
