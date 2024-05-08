package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  int
	Success bool
	Message string
	Data    interface{}
}

func (r *Response) DefaultResponse() {
	if r.Status == 0 {
		r.Status = http.StatusBadRequest
	}

	if r.Message == "" && r.Status == 0 {
		r.Message = "Bad request"
	}

	if r.Data == nil {
		r.Data = map[string]interface{}{}
	}
}

func JsonResponse(ctx *gin.Context, r *Response) {
	r.DefaultResponse()
	ctx.JSON(r.Status, gin.H{"success": r.Success, "message": r.Message, "data": r.Data})
}

func JsonSuccessResponse(ctx *gin.Context, params interface{}, message string) {
	response := Response{Status: http.StatusOK, Success: true, Message: message, Data: params}
	JsonResponse(ctx, &response)
}

func JsonCreatedResponse(ctx *gin.Context, params interface{}, message string) {
	response := Response{Status: http.StatusCreated, Success: true, Message: message, Data: params}
	JsonResponse(ctx, &response)
}

func JsonErrorResponse(ctx *gin.Context, message string) {
	response := Response{Message: message}
	JsonResponse(ctx, &response)
}

func JsonStatusErrorResponse(ctx *gin.Context, status int, message string) {
	response := Response{Status: status, Message: message}
	JsonResponse(ctx, &response)
}
