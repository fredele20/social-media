package middleware

import (
	"net/http"

	"github.com/fredele20/social-media/source/app/auth"
	"github.com/fredele20/social-media/source/setup"
	"github.com/fredele20/social-media/source/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")

	claims, err := auth.NewTokenService(*setup.Conf).ParseToken(token)
	if err != nil {
		utils.JsonStatusErrorResponse(ctx, http.StatusForbidden, err.Error())
		ctx.Abort()
		return
	}

	ctx.Set("userId", claims.UserId)
	ctx.Next()
}
