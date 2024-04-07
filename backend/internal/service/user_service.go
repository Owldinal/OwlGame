package service

import "owl-backend/internal/model"

func GetUserInfo(wallet string) (response interface{}, code model.ResponseCode, msg string) {

	response = &model.GetUserInfoResponse{
		Wallet: wallet,
	}
	return response, model.Success, ""
}
