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
	"time"
)

var (
	owlTokenContract *abigen.OwlToken
	owlGameContract  *abigen.OwlGame
	moonBoostAddress map[string]bool
)

func init() {
	var err error
	owlTokenContract, err = abigen.NewOwlToken(common.HexToAddress(config.C.OwlTokenAddr), eth.Client)
	if err != nil {
		log.Fatalf("Failed init user service: %v", err)
	}

	owlGameContract, err = abigen.NewOwlGame(common.HexToAddress(config.C.OwlGameAddr), eth.Client)

	// init moon boost list.
	moonBoostAddress = make(map[string]bool)
	addressList := []string{"0xAABB"}
	for _, address := range addressList {
		moonBoostAddress[address] = true
	}
}

func GetUserInfo(wallet string) (response *model.GetUserInfoResponse, code model.ResponseCode, msg string) {
	user, notFound, err := getCurrentUser(wallet)
	if notFound {
		return nil, model.NotFound, fmt.Sprintf("No user with address %v", wallet)
	} else if err != nil {
		return nil, model.ServerInternalError, "Error fetching user info"
	}

	// load owl balance
	amount, err := loadOwlBalanceByWallet(common.HexToAddress(wallet))
	if err != nil {
		return nil, model.ServerInternalError, "Error fetching owl balance"
		//amountDecimal := decimal.New(0, 0)
		//amount = &amountDecimal
	}

	// check moon boost
	isMoonBoost := false
	if config.C.NeedCheckMoonBoost {
		globalEnable, err := owlGameContract.IsMoonBoostEnable(&bind.CallOpts{})
		if err != nil {
			return nil, model.ServerInternalError, fmt.Sprintf("Failed to check moon boost (%v)", err)
		}
		if globalEnable {
			isMoonBoost = moonBoostAddress[wallet]
		}
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

			// calculate APR: 单个ELF的APR = （单个ELF的日平均Earning/100000）*365
			stakingRewards := token.CurrentRewards
			stakingDays := int64(time.Now().Sub(*token.StakingTime).Hours()) / 24
			apr := stakingRewards.Div(decimal.NewFromInt(stakingDays)).Div(constant.MysteryBoxMintPrice).Mul(decimal.NewFromInt32(365))
			boxInfo.Apr += apr.InexactFloat64()

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
		IsMoonBoost:    isMoonBoost,
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

func GetUserOwldinalList(wallet string, pagination model.PaginationRequest) (response *model.PaginationResponse[model.UserOwldinal], code model.ResponseCode, msg string) {
	result := &model.PaginationResponse[model.UserOwldinal]{
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

	var list []model.OwldinalNftToken
	err = db.Offset((pagination.Page - 1) * pagination.PerPage).
		Limit(pagination.PerPage).
		Find(&list).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result.List = []model.UserOwldinal{}
		} else {
			return nil, model.ServerInternalError, fmt.Sprintf("failed to find OwldinalNftToken records: %v", err)
		}
	} else {
		result.List = make([]model.UserOwldinal, 0, len(list))

		for _, token := range list {
			data := &model.UserOwldinal{
				TokenId:   token.TokenId,
				TokenUri:  token.TokenUri,
				IsStaking: token.IsStaking,
			}
			result.List = append(result.List, *data)
		}
	}

	return result, model.Success, ""
}

func GetUserMysteryBoxList(wallet string, pagination model.PaginationRequest) (response *model.PaginationResponse[model.UserMysteryBox], code model.ResponseCode, msg string) {
	response = &model.PaginationResponse[model.UserMysteryBox]{
		Page:    pagination.Page,
		PerPage: pagination.PerPage,
	}

	db := database.DB.Model(&model.MysteryBoxToken{}).Where("owner = ?", wallet)

	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return nil, model.ServerInternalError, fmt.Sprintf("failed to count MysteryBoxToken records: %v", err)
	}

	response.Total = int(total)
	response.PageCount = int((total + int64(pagination.PerPage) - 1) / int64(pagination.PerPage))

	var list []model.MysteryBoxToken
	err = db.Offset((pagination.Page - 1) * pagination.PerPage).
		Limit(pagination.PerPage).
		Find(&list).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.List = []model.UserMysteryBox{}
		} else {
			return nil, model.ServerInternalError, fmt.Sprintf("failed to find MysteryBoxToken records: %v", err)
		}
	} else {

		resultList := make([]model.UserMysteryBox, 0, len(list))
		for _, token := range list {
			data := &model.UserMysteryBox{
				TokenId:   token.TokenId,
				BoxType:   token.BoxType,
				Earning:   token.CurrentRewards, // TODO: Earning 是什么意思?
				IsStaking: token.IsStaking,
			}
			if token.IsStaking {
				stakingRewards := token.CurrentRewards
				stakingDays := int64(time.Now().Sub(*token.StakingTime).Hours()) / 24
				if stakingDays == 0 {
					stakingDays = 1
				}
				apr := stakingRewards.Div(decimal.NewFromInt(stakingDays)).Div(constant.MysteryBoxMintPrice).Mul(decimal.NewFromInt32(365))
				data.Apr = apr.InexactFloat64()
			}
			resultList = append(resultList, *data)
		}

		response.List = resultList
	}

	return response, model.Success, ""
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
