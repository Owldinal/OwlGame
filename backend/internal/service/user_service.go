package service

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"owl-backend/abigen"
	"owl-backend/internal/config"
	"owl-backend/internal/constant"
	"owl-backend/internal/database"
	"owl-backend/internal/eth"
	"owl-backend/internal/model"
	"owl-backend/pkg/log"
)

var owlTokenContract *abigen.OwlToken

func init() {
	var err error
	owlTokenContract, err = abigen.NewOwlToken(common.HexToAddress(config.C.OwlTokenAddr), eth.Client)
	if err != nil {
		log.Fatalf("Failed init user service: %v", err)
	}
}

func GetUserInfo(wallet string) (response interface{}, code model.ResponseCode, msg string) {
	user, notFound, err := getCurrentUser(wallet)
	if notFound {
		return nil, model.NotFound, fmt.Sprintf("No user with address %v", wallet)
	} else if err != nil {
		return nil, model.ServerInternalError, "Error fetching user info"
	}

	// load owl balance
	amount, err := loadOwlBalanceByWallet(common.HexToAddress(wallet))
	if err != nil {
		// TODO: hardhat 有问题，需要到线上链进行测试，之后修改
		//return nil, model.ServerInternalError, "Error fetching owl balance"
		amountDecimal := decimal.New(0, 0)
		amount = &amountDecimal
	}

	// Get Owldinal Nft
	var nftTokens []model.OwldinalNftToken
	result := database.DB.Where("owner = ?", wallet).Find(&nftTokens)
	if result.Error != nil {
		return nil, model.ServerInternalError, fmt.Sprintf("Error fetching NFT tokens (%v)", result.Error)
	}
	owlInfo := model.UserBoxInfo{
		Total: len(nftTokens),
	}
	for _, owldinal := range nftTokens {
		if owldinal.IsStaking {
			owlInfo.Staked += 1
			owlInfo.StakedIdList = append(owlInfo.StakedIdList, owldinal.TokenId)
		} else {
			owlInfo.UnstakedIdList = append(owlInfo.UnstakedIdList, owldinal.TokenId)
		}
	}

	// get mystery box info & total_earned
	var boxTokens []model.MysteryBoxToken
	if err := database.DB.Where("owner = ?", wallet).Find(&boxTokens).Error; err != nil {
		return nil, model.ServerInternalError, "Error fetching mystery boxes"
	}
	totalEarned := decimal.New(0, 0)
	fruitInfo := model.UserBoxInfo{}
	elfInfo := model.UserBoxInfo{}
	for _, token := range boxTokens {
		var boxInfo *model.UserBoxInfo
		if token.BoxType == constant.BoxTypeFruit {
			boxInfo = &fruitInfo
		} else {
			boxInfo = &elfInfo
		}
		boxInfo.Total += 1
		if token.IsStaking {
			boxInfo.Staked += 1
			boxInfo.StakedIdList = append(boxInfo.StakedIdList, token.TokenId)
		} else {
			boxInfo.UnstakedIdList = append(boxInfo.UnstakedIdList, token.TokenId)
		}
		// TODO: 需要确认这里 total earned 的含义，是不是所有 token 历史总收益的加总。
		// 也有可能是：1. 仅当前 stake 的收益，claim过的不算
		// 2. 仅 claim 的，还没提取的不算
		// 3. 也包括 invite referral 的，下个问题是否仅 cliam 还是 locked 的也算
		totalEarned = totalEarned.Add(token.TotalRewards)
	}

	response = &model.GetUserInfoResponse{
		Wallet:         wallet,
		OwlBalance:     amount.IntPart(),
		InvitationCode: user.InviteCode,
		InviteCount:    user.InviteCount,
		BuffLevel:      user.BuffLevel,
		OwlInfo:        owlInfo,
		ElfInfo:        elfInfo,
		FruitInfo:      fruitInfo,
		ReferralRewards: model.UserReferralRewards{
			Total:     user.ClaimedReferral.Add(user.LockedReferral).Add(user.AvailableReferral).IntPart(),
			Claimed:   user.ClaimedReferral.IntPart(),
			Available: user.AvailableReferral.IntPart(),
			Locked:    user.LockedReferral.IntPart(),
		},
	}
	return response, model.Success, ""
}

func GetUserOwldinalList(wallet string, pagination model.PaginationRequest) (response interface{}, code model.ResponseCode, msg string) {
	result := &model.PaginationResponse[model.OwldinalNftToken]{
		Page:    pagination.Page,
		PerPage: pagination.PerPage,
	}

	db := database.DB.Model(&model.OwldinalNftToken{}).Where("owner = ?", wallet)

	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return nil, model.ServerInternalError, fmt.Sprintf("failed to count OwldinalNftToken records: %v", err)
	}

	result.Total = int(total)
	result.PageCount = int((total + int64(pagination.PerPage) - 1) / int64(pagination.PerPage))

	var owldinalList []model.OwldinalNftToken
	err = db.Offset((pagination.Page - 1) * pagination.PerPage).
		Limit(pagination.PerPage).
		Find(&owldinalList).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result.List = []model.OwldinalNftToken{}
		} else {
			return nil, model.ServerInternalError, fmt.Sprintf("failed to find OwldinalNftToken records: %v", err)
		}
	} else {
		result.List = owldinalList
	}

	return result, model.Success, ""
}

func GetUserMysteryBoxList(wallet string, pagination model.PaginationRequest) (response interface{}, code model.ResponseCode, msg string) {
	result := &model.PaginationResponse[model.MysteryBoxToken]{
		Page:    pagination.Page,
		PerPage: pagination.PerPage,
	}

	db := database.DB.Model(&model.MysteryBoxToken{}).Where("owner = ?", wallet)

	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return nil, model.ServerInternalError, fmt.Sprintf("failed to count MysteryBoxToken records: %v", err)
	}

	result.Total = int(total)
	result.PageCount = int((total + int64(pagination.PerPage) - 1) / int64(pagination.PerPage))

	var list []model.MysteryBoxToken
	err = db.Offset((pagination.Page - 1) * pagination.PerPage).
		Limit(pagination.PerPage).
		Find(&list).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result.List = []model.MysteryBoxToken{}
		} else {
			return nil, model.ServerInternalError, fmt.Sprintf("failed to find MysteryBoxToken records: %v", err)
		}
	} else {
		result.List = list
	}

	return result, model.Success, ""
}

func GetUserInviteList(wallet string, pagination model.PaginationRequest) (response interface{}, code model.ResponseCode, msg string) {
	result := &model.PaginationResponse[model.InviteRelation]{
		Page:    pagination.Page,
		PerPage: pagination.PerPage,
	}

	db := database.DB.Model(&model.InviteRelation{}).Where("inviter = ?", wallet)

	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return nil, model.ServerInternalError, fmt.Sprintf("failed to count InviteRelation records: %v", err)
	}

	result.Total = int(total)
	result.PageCount = int((total + int64(pagination.PerPage) - 1) / int64(pagination.PerPage))

	var list []model.InviteRelation
	err = db.Offset((pagination.Page - 1) * pagination.PerPage).
		Limit(pagination.PerPage).
		Find(&list).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result.List = []model.InviteRelation{}
		} else {
			return nil, model.ServerInternalError, fmt.Sprintf("failed to find InviteRelation records: %v", err)
		}
	} else {
		result.List = list
	}

	return result, model.Success, ""
}

func getCurrentUser(wallet string) (userInfo *model.UserInfo, notFound bool, err error) {
	userInfoItem := model.UserInfo{
		Address: wallet,
	}
	if err := database.DB.Where(&userInfoItem).First(&userInfoItem).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, true, err
		} else {
			return nil, false, err
		}
	}

	return &userInfoItem, false, nil
}

func loadOwlBalanceByWallet(wallet common.Address) (*decimal.Decimal, error) {
	balance, err := owlTokenContract.BalanceOf(&bind.CallOpts{}, wallet)
	if err != nil {
		return nil, err
	}

	amount := decimal.NewFromBigInt(balance, -18)
	return &amount, nil
}
