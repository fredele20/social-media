package auth

import (
	"strings"
	"time"

	"github.com/fredele20/social-media/source/config"
	"github.com/fredele20/social-media/source/utils"
	"github.com/gbrlsnchs/jwt/v3"
)

func NewTokenService(conf config.Config) TokenServiceInterface {
	return &tokenService{conf: conf}
}

type JwtPayload struct {
	UserId string `json:"user_id"`
	jwt.Payload
}

func (t *tokenService) GenerateToken(userId string) (TokenResponse, error) {
	var tokenData TokenResponse
	payload := &JwtPayload{
		UserId: userId,
		Payload: jwt.Payload{
			ExpirationTime: jwt.NumericDate(expirationTime()),
		},
	}

	token, tokenErr := jwt.Sign(payload, jwt.NewHS256(t.jwtKey()))
	if tokenErr != nil {
		return tokenData, tokenErr
	}

	tokenData = TokenResponse{AccessToken: string(token), ExpiresAt: expirationTime()}
	return tokenData, nil
}

func (t *tokenService) RefereshToken(header string) (TokenResponse, error) {
	var tokenData TokenResponse

	return tokenData, nil
}

func (t *tokenService) ParseToken(header string) (JwtPayload, error) {
	var claims JwtPayload

	tokenString, headerErr := extractBearerToken(header)
	if headerErr != nil {
		return claims, headerErr
	}

	secret := jwt.NewHS256(t.jwtKey())
	_, err := jwt.Verify([]byte(tokenString), secret, &claims)
	if err != nil {
		return claims, ErrInvalidToken
	}

	return claims, nil
}

func expirationTime() time.Time {
	return utils.CurrentTime().Add(JWT_EXPIRED_DURATION * time.Hour)
}

func (t *tokenService) jwtKey() []uint8 {
	return []byte(t.conf.JwtSecretKey)
}

func extractBearerToken(header string) (string, error) {
	if header == "" {
		return "", ErrAuthRequired
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", ErrWrongHeader
	}

	return jwtToken[1], nil
}
