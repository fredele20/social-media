package auth

import (
	"net/http"

	"github.com/fredele20/social-media/source/utils"
	"github.com/gin-gonic/gin"
)

func NewAuthContoller(authService AuthServiceInterface) ControllerInterface {
	return &authController{
		authService: authService,
	}
}

func (a *authController) SignUp(ctx *gin.Context) {
	var params SignupRequest
	if !utils.ValidParams(ctx, &params) {
		return
	}

	data, err := a.authService.SignUp(ctx, &params)
	if err != nil {
		status := http.StatusBadRequest
		if err == ErrSignUpFailed {
			status = http.StatusUnprocessableEntity
		}

		utils.JsonStatusErrorResponse(ctx, status, err.Error())
		return
	}

	utils.JsonCreatedResponse(ctx, data, "successully created user")

}

func (a *authController) Login(ctx *gin.Context) {
	var param LoginRequest

	if !utils.ValidParams(ctx, &param) {
		return
	}

	data, err := a.authService.Login(ctx, &param)
	if err != nil {
		utils.JsonStatusErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	utils.JsonSuccessResponse(ctx, data, "successfully login")
}
