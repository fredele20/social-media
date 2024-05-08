package user

import "errors"

var (
	ErrDuplicateFollows = errors.New("you are already following this user")
)
