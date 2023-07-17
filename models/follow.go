package models

type Follows struct {
	UserId      string `json:"userId"`
	FollowingId string `json:"followingId"`
}

type ListFollowers struct {
	Followers []*Follows `json:"followers"`
	Count     int64      `json:"count"`
}

type ListFollowings struct {
	Followings []*Follows `json:"following"`
	Count      int64      `json:"count"`
}

type NewFollows struct {
	UserId    string   `json:"userId"`
	Followers []string `json:"followers"`
	Following []string `json:"following"`
}

type Followers struct {
	UserId    string   `json:"userId"`
	Followers []string `json:"followers"`
}

type Following struct {
	UserId    string   `json:"userId"`
	Following []string `json:"following"`
}
