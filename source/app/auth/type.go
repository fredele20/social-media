package auth

import (
	"context"

	"github.com/fredele20/social-media/source/app/user"
	"github.com/fredele20/social-media/source/config"
	"github.com/gin-gonic/gin"
)

type authController struct {
	authService AuthServiceInterface
}

type ControllerInterface interface {
	SignUp(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type authService struct {
	userRepo        user.RepositoryInterface
	userFollowsRepo user.UserFollowsRepositoryInterface
	tokenService    TokenServiceInterface
}

type AuthServiceInterface interface {
	SignUp(ctx context.Context, data *SignupRequest) (SignupResponse, error)
	Login(ctx context.Context, data *LoginRequest) (LoginResponse, error)
}

type tokenService struct {
	conf config.Config
}

type TokenServiceInterface interface {
	GenerateToken(userId string) (TokenResponse, error)
	ParseToken(token string) (JwtPayload, error)
}
