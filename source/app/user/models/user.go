package models

import (
	"time"
)

type User struct {
	Id          string    `json:"id" bson:"id"`
	Firstname   string    `json:"firstname" bson:"firstname"`
	Lastname    string    `json:"lastname"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	PictureURL  string    `json:"pictureUrl"`
	Password    string    `json:"password"`
	Status      Status    `json:"status"`
	Token       *string   `json:"token"`
	Country     string    `json:"country"`
	CountryCode string    `json:"country_code"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type UserFilter struct {
	Limit int64 `json:"limit"`
}

type ListUsers struct {
	Users []*User `json:"users"`
	Count int64   `json:"count"`
}

type Status string

const (
	ACTIVATED   Status = "activated"
	DEACTIVATED Status = "deactivated"
)
