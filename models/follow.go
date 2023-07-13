package models

type Follows struct {
	UserId      string   `json:"userId"`
	FollowersId []string `json:"followerId"`
}
