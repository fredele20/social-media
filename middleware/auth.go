package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/fredele20/social-media/libs/sessions"
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/gin-gonic/gin"
)

var (
	ErrTokenInvalid            = errors.New("invalid token string provided")
	ErrNoAuthHeaderProvided    = errors.New("no authentication header provided")
	ErrTokenVerificationFailed = errors.New("could not verify the token provided")
)

var JwtSecretKey = os.Getenv("JWT_SECRETKEY")

func VerifyAuthToken(token string) (*sessions.TokenPayload, error) {
	secret := jwt.NewHS256([]byte(JwtSecretKey))
	var tokenPayload sessions.TokenPayload
	_, err := jwt.Verify([]byte(token), secret, &tokenPayload)
	if err != nil {
		return nil, ErrTokenInvalid
	}

	return &tokenPayload, nil
}

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": ErrNoAuthHeaderProvided.Error()})
			ctx.Abort()
			return
		}

		payload, err := VerifyAuthToken(token)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": ErrTokenVerificationFailed.Error()})
			ctx.Abort()
			return
		}

		ctx.Set("id", payload.Id)
		ctx.Set("email", payload.Email)
		ctx.Next()
	}
}
