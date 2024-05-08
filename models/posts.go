package models

type Posts struct {
	Id       string   `json:"id"`
	UserId   string   `json:"userId"`
	Content  string   `json:"content"`
	ImageURL string   `json:"imageUrl"`
	Likes    []string `json:"likes" bson:"likes"`
}

type ListPosts struct {
	Count int64    `json:"count"`
	Posts []*Posts `json:"posts"`
}

// type ListFollowingUsersPosts struct {
// 	UserId string `json:"userid"`
// }

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
	Id       string `json:"id"`
	PostId   string `json:"postId"`
	UserId   string `json:"userId"`
	Content  string `json:"content"`
	ImageURL string `json:"imageUrl"`
}
