package core

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/fredele20/social-media/database"
	"github.com/fredele20/social-media/database/mongod"
	"github.com/fredele20/social-media/models"
	"github.com/fredele20/social-media/utils"
	"github.com/nyaruka/phonenumbers"
	"github.com/sirupsen/logrus"
)

var (
	ErrUserValidationFailed = errors.New("failed to validate user before persisting")
	ErrCreateUserDuplicate  = errors.New("failed to create user, duplicate record found")
	ErrCreateUserFailed     = errors.New("user creation failed")
	ErrPhoneNumberNotValid = errors.New("sorry, phone number cannot be used")
)

type UserService struct {
	mongo  database.UserDatastore
	logger *logrus.Logger
}

func NewUserService(mongo database.UserDatastore, l *logrus.Logger) *UserService {
	return &UserService{
		mongo:  mongo,
		logger: l,
	}
}

func buildPictureFromName(name string) string {
	return fmt.Sprintf("https://ui-avatars.com/api/?name=%s", strings.ReplaceAll(name, " ", "+"))
}

func parsePhone(phone, iso2 string) (string, error) {
	num, err := phonenumbers.Parse(phone, iso2)
	if err != nil {
		return "", err
	}

	switch phonenumbers.GetNumberType(num) {
	case phonenumbers.VOIP, phonenumbers.VOICEMAIL:
		return "", ErrPhoneNumberNotValid
	}

	return phonenumbers.Format(num, phonenumbers.E164), nil
}

func (u *UserService) RegisterUser(ctx context.Context, payload *models.Users) (*models.Users, error) {
	if err := payload.Validate(); err != nil {
		u.logger.WithError(err).Error(ErrUserValidationFailed.Error())
		return nil, err
	}

	phone, _ := parsePhone(payload.Phone, payload.Iso2)
	password := utils.HashPassword(payload.Password)

	payload.PictureURL = buildPictureFromName(fmt.Sprintf("%s+%s", payload.Firstname, payload.Lastname))
	payload.Password = password
	payload.Phone = phone
	payload.Status = models.NotActive
	payload.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	payload.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	user, err := u.mongo.CreateUser(ctx, payload)
	if err != nil {
		fmt.Println(err.Error())
		if err == mongod.ErrDuplicate {
			u.logger.WithError(err).Error("create user failed, attempted duplicate record")
			return nil, ErrCreateUserDuplicate
		}
		u.logger.WithError(err).Error(err.Error())
		return nil, ErrCreateUserFailed
	}

	return user, nil
}
