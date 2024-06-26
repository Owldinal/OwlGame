package api

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"math/big"
	"owl-backend/internal/config"
	"owl-backend/internal/eventlistener"
	"owl-backend/internal/model"
	"owl-backend/internal/service"
	"owl-backend/internal/util"
	"owl-backend/pkg/log"
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

func ClaimBoxes(c *gin.Context) {
	var req model.ClaimBoxRequest
	if err := c.ShouldBind(&req); err != nil {
		ErrorResponse(c, model.WrongParam, "Missing Param")
		return
	}
	if !common.IsHexAddress(req.Address) {
		ErrorResponse(c, model.WrongParam, "Wrong Param 'address'")
		return
	}

	err := util.ValidSignature(req.Address, req.Message, req.Signature)
	if err != nil {
		ErrorResponse(c, model.InvalidSignature, err.Error())
		return
	}
	data, code, msg := service.ClaimToken(req.Address, req.TokenIds)

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
	var req model.RewardsHistoryRequest
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

func GetRequestMintTx(c *gin.Context) {
	var req model.GetMintTxRequest
	if err := c.ShouldBind(&req); err != nil {
		ErrorResponse(c, model.WrongParam, "Missing Param")
		return
	}

	data, code, msg := service.GetMintTx(req.Tx)

	if code == model.Success {
		SuccessResponse(c, data)
	} else {
		ErrorResponse(c, code, msg)
	}
}

func UpdateRewards(c *gin.Context) {
	var req model.UpdateFruitRequest
	if err := c.ShouldBind(&req); err != nil {
		ErrorResponse(c, model.WrongParam, "Missing Param")
		return
	}

	if req.Secret != config.C.Secret {
		ErrorResponse(c, model.WrongParam, "Wrong Param")
		return
	}

	log.Infof("Start update fruit rewards: %v", req.CompareTime)
	eventlistener.UpdateFruitReward(req.CompareTime)

	SuccessResponse(c, nil)
	//if code == model.Success {
	//	SuccessResponse(c, data)
	//} else {
	//	ErrorResponse(c, code, msg)
	//}
}

func RetryAllJobs(c *gin.Context) {
	var req model.AdminSecret
	if err := c.ShouldBind(&req); err != nil {
		ErrorResponse(c, model.WrongParam, "Missing Param")
		return
	}

	if req.Secret != config.C.Secret {
		ErrorResponse(c, model.WrongParam, "Wrong Param")
		return
	}

	data, code, msg := service.RetryAllJobs()

	if code == model.Success {
		SuccessResponse(c, data)
	} else {
		ErrorResponse(c, code, msg)
	}
}

func RetryTransferMultiple(c *gin.Context) {
	var req model.RetryTransferMultipleRequest
	if err := c.ShouldBind(&req); err != nil {
		ErrorResponse(c, model.WrongParam, "Missing Param")
		return
	}

	if req.Secret != config.C.Secret {
		ErrorResponse(c, model.WrongParam, "Wrong Param")
		return
	}

	data, code, msg := service.RetryTransferMultiple(req.TaskId, req.Transfer, req.Burn)

	if code == model.Success {
		SuccessResponse(c, data)
	} else {
		ErrorResponse(c, code, msg)
	}
}

func GetRequestJob(c *gin.Context) {
	var req model.RequestJobRequest
	if err := c.ShouldBind(&req); err != nil {
		ErrorResponse(c, model.WrongParam, "Missing Param")
		return
	}
	data, code, msg := service.GetRequestJobStatus(req.Tx)

	if code == model.Success {
		SuccessResponse(c, data)
	} else {
		ErrorResponse(c, code, msg)
	}
}

func ReloadLog(c *gin.Context) {
	var req model.ReloadLogRequest
	if err := c.ShouldBind(&req); err != nil {
		ErrorResponse(c, model.WrongParam, "Missing Param")
		return
	}

	block := big.NewInt(req.Block)
	eventlistener.ProcessLogs(block, block)
	SuccessResponse(c, nil)
}

func CheckSignature(c *gin.Context) {
	var req model.SignedRequest
	if err := c.ShouldBind(&req); err != nil {
		ErrorResponse(c, model.WrongParam, "Missing Param")
		return
	}

	if !common.IsHexAddress(req.Address) {
		ErrorResponse(c, model.WrongParam, "Wrong Param 'wallet'")
		return
	}

	err := util.ValidSignature(req.Address, req.Message, req.Signature)
	if err != nil {
		ErrorResponse(c, model.InvalidSignature, err.Error())
		return
	}

	SuccessResponse(c, true)
}
