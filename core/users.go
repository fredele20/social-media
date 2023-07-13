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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrUserValidationFailed     = errors.New("failed to validate user before persisting")
	ErrCreateUserDuplicate      = errors.New("failed to create user, duplicate record found")
	ErrCreateUserFailed         = errors.New("user creation failed")
	ErrPhoneNumberNotValid      = errors.New("sorry, phone number cannot be used")
	ErrListUserFailed           = errors.New("failed to list users")
	ErrFailedToGetUserByEmail   = errors.New("failed to get user with provided email")
	ErrFailedToGetUserById      = errors.New("failed to get user with the provided id")
	ErrCreateUserFollowerFailed = errors.New("failed to create a follower for the user")
)

type CoreService struct {
	db     database.Datastore
	logger *logrus.Logger
}

func NewCoreService(db database.Datastore, l *logrus.Logger) *CoreService {
	return &CoreService{
		db:     db,
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

func (c *CoreService) RegisterUser(ctx context.Context, payload *models.Users) (*models.Users, error) {
	if err := payload.Validate(); err != nil {
		c.logger.WithError(err).Error(ErrUserValidationFailed.Error())
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
	generateId := primitive.NewObjectID()
	payload.Id = generateId.Hex()

	user, err := c.db.CreateUser(ctx, payload)
	if err != nil {
		fmt.Println(err.Error())
		if err == mongod.ErrDuplicate {
			c.logger.WithError(err).Error("create user failed, attempted duplicate record")
			return nil, ErrCreateUserDuplicate
		}
		c.logger.WithError(err).Error(err.Error())
		return nil, ErrCreateUserFailed
	}

	return user, nil
}

func (c *CoreService) ListUsers(ctx context.Context, filters *models.UserFilter) (*models.ListUsers, error) {
	users, err := c.db.ListUsers(ctx, filters)
	if err != nil {
		c.logger.WithError(err).Error(ErrListUserFailed)
		return nil, err
	}

	return users, nil
}

func (c *CoreService) GetUserByEmail(ctx context.Context, email string) (*models.Users, error) {
	user, err := c.db.GetUserByEmail(ctx, email)
	if err != nil {
		c.logger.WithError(err).Error(ErrFailedToGetUserByEmail.Error())
		return nil, err
	}

	return user, nil
}

func (c *CoreService) GetUserById(ctx context.Context, id string) (*models.Users, error) {
	user, err := c.db.GetUserById(ctx, id)
	if err != nil {
		c.logger.WithError(err).Error(ErrFailedToGetUserById)
		return nil, err
	}
	return user, nil
}

func (c *CoreService) CreateUserFollower(ctx context.Context, payload *models.Follows) (*models.Follows, error) {
	follow, err := c.db.CreateUserFollower(ctx, payload)
	if err != nil {
		c.logger.WithError(err).Error(ErrCreateUserFollowerFailed)
		return nil, err
	}

	return follow, nil
}
