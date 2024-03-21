package api

import (
	"github.com/gin-gonic/gin"
)

// CheckHealth godoc
// @Summary      Show an account
// @Description  get string by ID
// @Router       /checkHealth [get]
func CheckHealth(c *gin.Context) {
	SuccessResponse(c, "hello")
}
