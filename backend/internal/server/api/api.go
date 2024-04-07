package api

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"owl-backend/internal/model"
	"owl-backend/internal/service"
)

func GetUserInfo(c *gin.Context) {
	var req model.GetUserInfoRequest
	if err := c.ShouldBind(&req); err != nil {
		ErrorResponse(c, model.WrongParam, "Missing Param")
		return
	}

	if !common.IsHexAddress(req.Wallet) {
		ErrorResponse(c, model.WrongParam, "Wrong Param 'wallet'")
		return
	}

	data, code, msg := service.GetUserInfo(req.Wallet)
	if code == model.Success {
		SuccessResponse(c, data)
	} else {
		ErrorResponse(c, code, msg)
	}
}
