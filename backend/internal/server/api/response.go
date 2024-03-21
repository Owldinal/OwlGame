package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"owl-backend/internal/model"
)

func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, model.ResponseData{
		Code:    model.Success,
		Success: true,
		Message: "",
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, code model.ResponseCode, message string) {
	c.JSON(http.StatusOK, model.ResponseData{
		Code:    code,
		Success: false,
		Message: message,
		Data:    nil,
	})
}
