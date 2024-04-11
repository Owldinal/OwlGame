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

func GetUserOwldinalList(c *gin.Context) {
	var req model.GetUserOwldinalsRequest
	if err := c.ShouldBind(&req); err != nil {
		ErrorResponse(c, model.WrongParam, "Missing Param")
		return
	}

	if !common.IsHexAddress(req.Wallet) {
		ErrorResponse(c, model.WrongParam, "Wrong Param 'wallet'")
		return
	}

	if req.PaginationRequest.Page < 1 {
		req.PaginationRequest.Page = 1
	}
	if req.PaginationRequest.PerPage < 1 {
		req.PaginationRequest.PerPage = 20
	}

	data, code, msg := service.GetUserOwldinalList(req.Wallet, req.PaginationRequest)

	if code == model.Success {
		SuccessResponse(c, data)
	} else {
		ErrorResponse(c, code, msg)
	}
}

func GetUserBoxList(c *gin.Context) {
	var req model.GetUserOwldinalsRequest
	if err := c.ShouldBind(&req); err != nil {
		ErrorResponse(c, model.WrongParam, "Missing Param")
		return
	}

	if !common.IsHexAddress(req.Wallet) {
		ErrorResponse(c, model.WrongParam, "Wrong Param 'wallet'")
		return
	}

	if req.PaginationRequest.Page < 1 {
		req.PaginationRequest.Page = 1
	}
	if req.PaginationRequest.PerPage < 1 {
		req.PaginationRequest.PerPage = 20
	}

	data, code, msg := service.GetUserMysteryBoxList(req.Wallet, req.PaginationRequest)

	if code == model.Success {
		SuccessResponse(c, data)
	} else {
		ErrorResponse(c, code, msg)
	}
}

func GetUserInviteList(c *gin.Context) {
	var req model.GetUserInviterRequest
	if err := c.ShouldBind(&req); err != nil {
		ErrorResponse(c, model.WrongParam, "Missing Param")
		return
	}

	if !common.IsHexAddress(req.Wallet) {
		ErrorResponse(c, model.WrongParam, "Wrong Param 'wallet'")
		return
	}

	if req.PaginationRequest.Page < 1 {
		req.PaginationRequest.Page = 1
	}
	if req.PaginationRequest.PerPage < 1 {
		req.PaginationRequest.PerPage = 20
	}

	data, code, msg := service.GetUserInviteList(req.Wallet, req.PaginationRequest)

	if code == model.Success {
		SuccessResponse(c, data)
	} else {
		ErrorResponse(c, code, msg)
	}
}

func GetGameInfo(c *gin.Context) {
	data, code, msg := service.GetGameInfo()

	if code == model.Success {
		SuccessResponse(c, data)
	} else {
		ErrorResponse(c, code, msg)
	}
}

func GetRewardTrend(c *gin.Context) {
	data, code, msg := service.GetRewardsTrending()

	if code == model.Success {
		SuccessResponse(c, data)
	} else {
		ErrorResponse(c, code, msg)
	}
}

func GetRewardHistory(c *gin.Context) {
	var req model.CursorRequest
	if err := c.ShouldBind(&req); err != nil {
		ErrorResponse(c, model.WrongParam, "Missing Param")
		return
	}

	if req.Limit < 1 {
		req.Limit = 20
	}

	data, code, msg := service.GetTreasuryRevenueHistory(req)

	if code == model.Success {
		SuccessResponse(c, data)
	} else {
		ErrorResponse(c, code, msg)
	}
}
