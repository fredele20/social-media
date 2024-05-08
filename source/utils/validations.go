package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

const (
	DESCRIPTION_LENGTH = 100
	INVALID_INPUT = "bad request"
)

func ValidParams(ctx *gin.Context, params interface{}) bool {
	if err := ctx.ShouldBindBodyWith(&params, binding.JSON); err != nil {
		JsonErrorResponse(ctx, INVALID_INPUT)
		return false
	}

	paramsMap := ConvertDataToMap(params)
	limitDataKeys := [2]string{"naration", "service_description"}
	for _, k := range limitDataKeys {
		if paramsMap[k] != nil && !validParamLength(ctx, paramsMap[k].(string), k) {
			return false
		}
	}

	return true
}

func validParamLength(ctx *gin.Context, value, key string) bool {
	if len(value) > DESCRIPTION_LENGTH {
		errMsg := fmt.Sprintf("%s should not be more than %d characters", key, DESCRIPTION_LENGTH)
		JsonErrorResponse(ctx, errMsg)
		return false
	}

	return true
}
