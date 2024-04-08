package service

import (
	"log"
	"owl-backend/internal/database"
	"owl-backend/internal/model"
)

func GetUserInfo(wallet string) (response interface{}, code model.ResponseCode, msg string) {

	// Get Owldinal Nft
	var nftTokens []model.OwldinalNftToken
	result := database.DB.Where("owner = ?", wallet).Find(&nftTokens)
	if result.Error != nil {
		log.Printf("Error fetching NFT tokens for wallet %s: %v", wallet, result.Error)
		return nil, model.ServerInternalError, "Error fetching NFT tokens"
	}
	owlInfo := model.UserBoxInfo{
		Total: len(nftTokens),
	}
	for _, owldinal := range nftTokens {
		if owldinal.IsStaking {
			owlInfo.Staked += 1
			owlInfo.StakedIdList = append(owlInfo.StakedIdList, owldinal.ID)
		} else {
			owlInfo.UnstakedIdList = append(owlInfo.UnstakedIdList, owldinal.ID)
		}
	}

	response = &model.GetUserInfoResponse{
		Wallet:  wallet,
		OwlInfo: owlInfo,
	}
	return response, model.Success, ""
}
