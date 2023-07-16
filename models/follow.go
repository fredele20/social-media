package models

type Follows struct {
	FollowingId string `json:"followingId"`
	UserId      string `json:"userId"`
}

type ListFollowers struct {
	Followers []*Follows `json:"followers"`
	Count int64 `json:"count"`
}

type ListFollowings struct {
	Followings []*Follows `json:"following"`
	Count int64 `json:"count"`
}
