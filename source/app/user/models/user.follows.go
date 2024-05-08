package models

type Follows struct {
	UserId    string   `json:"user_id" bson:"user_id"`
	Followers []string `json:"followers" bson:"followers"`
	Following []string `json:"following" bson:"following"`
}
