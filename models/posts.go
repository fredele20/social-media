package models

type Posts struct {
	Id       string     `json:"id"`
	UserId   string     `json:"userId"`
	Content  string     `json:"content"`
	ImageURL string     `json:"imageUrl"`
	Likes    []string   `json:"likes" bson:"likes"`
	Comments []Comments `json:"comments" bson:"comments"`
}

type ListPosts struct {
	Count int64    `json:"count"`
	Posts []*Posts `json:"posts"`
}

type PostFilters struct {
	Limit int64 `json:"limit"`
}

//	type Likes struct {
//		PostId string             `json:"postId"`
//		UserId string             `json:"userId"`
//	}
type Like struct {
	Liker string `json:"liker"`
}

type CommentList struct {
	Comments []*Comments `json:"comments"`
	Count    int64       `json:"count"`
}

type Comments struct {
	CommenterId string `json:"commenter"`
	Content     string `json:"content"`
}
