package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Posts struct {
	Id       primitive.ObjectID `json:"id"`
	UserId   string             `json:"userId"`
	Content  string             `json:"content"`
	ImageURL string             `json:"imageUrl"`
	Likes    []string           `json:"likes"`
	Comments []*Comments           `json:"comments"`
}

type ListPosts struct {
	Count int64    `json:"count"`
	Posts []*Posts `json:"posts"`
}

type PostFilters struct {
	Limit int64 `json:"limit"`
}

// type Likes struct {
// 	PostId string             `json:"postId"`
// 	UserId string             `json:"userId"`
// }

type Comments struct {
	Commenter string `json:"commenter"`
	Content   string `json:"content"`
}
