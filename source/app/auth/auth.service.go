package auth

import (
	"context"
	"fmt"
	"strings"

	"github.com/fredele20/social-media/source/app/user"
	"github.com/fredele20/social-media/source/app/user/models"
	"github.com/fredele20/social-media/source/utils"
	"github.com/jinzhu/copier"
	"github.com/nyaruka/phonenumbers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewService(
	userRepo user.RepositoryInterface,
	userFollowsRepo user.UserFollowsRepositoryInterface,
	tokenService TokenServiceInterface,
) AuthServiceInterface {
	return &authService{
		userRepo:        userRepo,
		tokenService:    tokenService,
		userFollowsRepo: userFollowsRepo,
	}
}

func (a *authService) SignUp(ctx context.Context, data *SignupRequest) (SignupResponse, error) {
	var response SignupResponse
	if payloadErr := data.Validate(); payloadErr != nil {
		return response, payloadErr
	}

	phone, _ := a.parsePhone(data.Phone, data.CountryCode)

	userData := models.User{
		Firstname:   data.Firstname,
		Lastname:    data.Lastname,
		Email:       data.Email,
		Password:    utils.HashPassword(data.Password),
		Phone:       phone,
		Country:     data.Country,
		CountryCode: data.CountryCode,
	}

	user, tokenData, createUserErr := a.createUserWithToken(ctx, userData)
	if createUserErr != nil {
		return response, createUserErr
	}

	copier.Copy(&response.User, user)
	response.Token = tokenData

	go a.userFollowsRepo.CreateUserFollows(ctx, user.Id)

	return response, nil
}

func (a *authService) Login(ctx context.Context, data *LoginRequest) (LoginResponse, error) {
	var response LoginResponse
	user, userErr := a.userRepo.FindUserByEmailOrPhone(ctx, data.LoginId)
	if userErr != nil || !utils.VerifyPassword(user.Password, data.Password) {
		return response, ErrInvalidLogin
	}

	tokenData, tokenErr := a.tokenService.GenerateToken(user.Id)
	if tokenErr != nil {
		return response, tokenErr
	}

	copier.Copy(&response.User, &user)
	response.Token = tokenData

	return response, nil
}

func (a *authService) createUserWithToken(ctx context.Context, userData models.User) (models.User, TokenResponse, error) {
	var tokenData TokenResponse

	userData.PictureURL = a.buildPictureFromName(fmt.Sprintf("%s %s", userData.Firstname, userData.Lastname))
	generateId := primitive.NewObjectID()
	userData.Id = generateId.Hex()
	userData.Status = models.DEACTIVATED
	userData.CreatedAt, _ = utils.ConvertUtcToNigerianTime(utils.CurrentTime())
	userData.UpdatedAt, _ = utils.ConvertUtcToNigerianTime(utils.CurrentTime())

	user, userErr := a.userRepo.CreateUser(ctx, &userData)
	if userErr != nil {
		return user, tokenData, userErr
	}

	tokenResult, tokenErr := a.tokenService.GenerateToken(user.Id)
	if tokenErr != nil {
		return user, tokenData, tokenErr
	}

	return user, tokenResult, nil
}

func (a *authService) buildPictureFromName(name string) string {
	return fmt.Sprintf("https://ui-avatars.com/api/?name=%s", strings.ReplaceAll(name, " ", "+"))
}

func (a *authService) parsePhone(phone, iso2 string) (string, error) {
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
