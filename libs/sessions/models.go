package sessions

import (
	"encoding/json"
	"time"

	"github.com/gbrlsnchs/jwt/v3"
)

type Session struct {
	Token          string         `json:"token"`
	UserID         string         `json:"userId"`
	Email          string         `json:"email"`
	TimeCreated    time.Time      `json:"timeCreated"`
	Validity       time.Duration  `json:"validity"`
	UnitOfValidity UnitOfValidity `json:"unitOfValidity"`
}

type TokenPayload struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	jwt.Payload
}

type UnitOfValidity string

const (
	UnitOfValidityMinute UnitOfValidity = "MINUTE"
	UnitOfValidityHour   UnitOfValidity = "HOUR"
)

func (u UnitOfValidity) IsValid() bool {
	switch u {
	case UnitOfValidityHour, UnitOfValidityMinute:
		return true
	}

	return false
}

func (u UnitOfValidity) String() string {
	return string(u)
}

func (sm *Session) ToByte() []byte {
	if sm == nil {
		return nil
	}

	b, _ := json.Marshal(sm)

	return b
}
