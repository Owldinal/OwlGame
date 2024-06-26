package api

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"owl-backend/internal/config"
	"owl-backend/internal/model"
	"owl-backend/internal/service"
)

func GetMintSignature(c *gin.Context) {
	var req model.GetMintSignatureRequest
	if err := c.ShouldBind(&req); err != nil {
		ErrorResponse(c, model.WrongParam, "Missing Param")
		return
	}

	if !common.IsHexAddress(req.Wallet) {
		ErrorResponse(c, model.WrongParam, "Wrong Param 'wallet'")
		return
	}

	if !config.C.DisableHCaptcha {
		_, err := service.VerifyCaptcha(req.HCaptcha)
		if err != nil {
			ErrorResponse(c, model.HCaptchaFailed, err.Error())
			return
		}
	}

	data, code, msg := service.GenerateHashAndSignForMint(req.Wallet)
	if code == model.Success {
		SuccessResponse(c, data)
	} else {
		ErrorResponse(c, code, msg)
	}
}
