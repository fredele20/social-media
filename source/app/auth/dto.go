package auth

import (
	"time"

	"github.com/fredele20/social-media/source/app/user"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/sirupsen/logrus"
)

type SignupRequest struct {
	Firstname   string `json:"firstname" binding:"required"`
	Lastname    string `json:"lastname" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Phone       string `json:"phone" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Country     string `json:"country" binding:"required"`
	CountryCode string `json:"country_code" binding:"required"`
}

type SignupResponse struct {
	Token TokenResponse     `json:"token,omitempty"`
	User  user.UserResponse `json:"user,omitempty"`
}

type LoginRequest struct {
	LoginId  string `json:"login_id"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token TokenResponse     `json:"token,omitempty"`
	User  user.UserResponse `json:"user,omitempty"`
}

type TokenResponse struct {
	AccessToken string    `json:"access_token"`
	ExpiresAt   time.Time `json:"expires_at"`
}

func (u SignupRequest) Validate() error {
	if err := validation.ValidateStruct(&u,
		validation.Field(&u.Firstname, validation.Required),
		validation.Field(&u.Lastname, validation.Required),
		// validation.Field(&u.Email, validation.Required, is.Email),
		// validation.Field(&u.Phone, validation.Required, is.E164),
		validation.Field(&u.Password, validation.Required),
		validation.Field(&u.Country, validation.Required),
	); err != nil {
		logrus.WithError(err).Error(err.Error())
		return err
	}
	return nil
}
