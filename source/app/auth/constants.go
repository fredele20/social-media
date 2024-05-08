package auth

import "errors"

const (
	JWT_EXPIRED_DURATION = 24 // in hours
)

var (
	ErrAuthRequired        = errors.New("authorization is required")
	ErrWrongHeader         = errors.New("wrongly formatted authorization header")
	ErrInvalidToken        = errors.New("invalid authorization token")
	ErrInvalidLogin        = errors.New("invalid login credentials")
	ErrPhoneNumberNotValid = errors.New("sorry, phone number cannot be used")
	ErrSignUpFailed        = errors.New("an error occured during signup, pls try again")
)
