package user

import "time"

type UserResponse struct {
	ID                    string     `json:"id"`
	FirstName             string     `json:"first_name"`
	LastName              string     `json:"last_name"`
	Email                 string     `json:"email"`
	Phone                 string     `json:"phone"`
	CountryCode           string     `json:"country_code"`
	PinSet                bool       `json:"pin_set"`
	ProfilePictureUrl     string     `json:"profile_picture_url"`
	EmailVerifiedAt       *time.Time `json:"email_verified_at"`
	PhoneNumberVerifiedAt *time.Time `json:"phone_number_verified_at"`
	CreatedAt             time.Time  `json:"created_at"`
}

type FollowUserRequest struct {
	UserId string `json:"user_id" binding:"required"`
}

type UserFollowsDetails struct {
	UserId    string      `json:"user_id"`
	Followers UserFollows `json:"followers"`
	Following UserFollows `json:"following"`
}

type UserFollows struct {
	Count int      `json:"count"`
	Data  []string `json:"data"`
}
