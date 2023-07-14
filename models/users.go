package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/sirupsen/logrus"
)

type Users struct {
	Id         string    `json:"id"`
	Firstname  string    `json:"firstname"`
	Lastname   string    `json:"lastname"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	PictureURL string    `json:"pictureUrl"`
	Password   string    `json:"password"`
	Status     Status    `json:"status"`
	Token      *string   `json:"token"`
	Country    string    `json:"country"`
	Iso2       string    `json:"iso2"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type UserFilter struct {
	Limit int64 `json:"limit"`
}

type ListUsers struct {
	Users []*Users `json:"users"`
	Count int64    `json:"count"`
}

func (u Users) Validate() error {
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

type Status string

const (
	Active    Status = "active"
	NotActive Status = "notActive"
)
