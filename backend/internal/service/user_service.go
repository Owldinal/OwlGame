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
	"owl-backend/internal/eventlistener"
	"owl-backend/internal/model"
	"owl-backend/pkg/log"
)

var (
	owlTokenContract *abigen.OwlToken
	owlGameContract  *abigen.OwlGame
)

func init() {
	var err error
	owlTokenContract, err = abigen.NewOwlToken(common.HexToAddress(config.C.OwlTokenAddr), eth.Client)
	if err != nil {
		log.Fatalf("Failed init user service: %v", err)
	}

	owlGameContract, err = abigen.NewOwlGame(common.HexToAddress(config.C.OwlGameAddr), eth.Client)
}

func GetUserInfo(wallet string) (response *model.GetUserInfoResponse, code model.ResponseCode, msg string) {
	user, notFound, err := getCurrentUser(wallet)
	if notFound {
		//return nil, model.NotFound, fmt.Sprintf("No user with address %v", wallet)
		// if not found ,still return all the struct
		user = &model.UserInfo{}
	} else if err != nil {
		return nil, model.ServerInternalError, "Error fetching user info"
	}

	// load owl balance
	amount, err := loadOwlBalanceByWallet(common.HexToAddress(wallet))
	if err != nil {
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

	// check moon boost
	isMoonBoost := false
	if config.C.NeedCheckMoonBoost && !notFound {
		isMoonBoost = constant.MoonBoostAddress[wallet] || owlInfo.Total > 0
	}

	// get mystery box info & total_earned
	var boxTokens []model.MysteryBoxToken
	if err := database.DB.Where("owner = ?", wallet).Find(&boxTokens).Error; err != nil {
		return nil, model.ServerInternalError, "Error fetching mystery boxes"
	}
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
	}
	var lastApr model.AprSnapshot
	if err := database.DB.Order("id DESC").First(&lastApr).Error; err != nil {
		return nil, model.ServerInternalError, fmt.Sprintf("failed to load apr: %v", err)
	}
	fruitInfo.Apr = lastApr.FruitApr
	fruitInfo.Apy = lastApr.FruitApy
	elfInfo.Apr = lastApr.ElfApr
	elfInfo.Apy = lastApr.ElfApy

	// calculate referral
	totalReferral := user.ClaimedReferral.Add(user.UnclaimedReferral)
	lockedReferral := user.UnclaimedReferral.Sub(user.UnlockableReferral)
	if lockedReferral.IsNegative() {
		lockedReferral = decimal.Zero
	}
	availableReferral := user.UnclaimedReferral.Sub(lockedReferral)

	response = &model.GetUserInfoResponse{
		Wallet:         wallet,
		HasJoined:      !notFound,
		OwlBalance:     amount.IntPart(),
		TotalEarned:    user.TotalEarned,
		InvitationCode: user.InviteCode,
		InviteCount:    user.InviteCount,
		BuffLevel:      user.BuffLevel,
		IsMoonBoost:    isMoonBoost,
		OwlInfo:        owlInfo,
		ElfInfo:        elfInfo,
		FruitInfo:      fruitInfo,
		ReferralRewards: model.UserReferralRewards{
			Total:     totalReferral.IntPart(),
			Claimed:   user.ClaimedReferral.IntPart(),
			Available: availableReferral.IntPart(),
			Locked:    lockedReferral.IntPart(),
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

		var lastApr model.AprSnapshot
		if err := database.DB.Order("id DESC").First(&lastApr).Error; err != nil {
			return nil, model.ServerInternalError, fmt.Sprintf("failed to load apr: %v", err)
		}

		resultList := make([]model.UserMysteryBox, 0, len(list))
		for _, token := range list {
			data := &model.UserMysteryBox{
				TokenId:   token.TokenId,
				BoxType:   token.BoxType,
				Earning:   token.CurrentRewards,
				Claimed:   token.ClaimedRewards,
				IsStaking: token.IsStaking,
			}

			if token.BoxType == constant.BoxTypeFruit {
				data.Apr = lastApr.FruitApr
				data.Apy = lastApr.FruitApy
			} else if token.BoxType == constant.BoxTypeElf {
				data.Apr = lastApr.ElfApr
				data.Apy = lastApr.ElfApy
			}

			//if token.IsStaking {
			//	// 单个ELF的APR= （单个ELF的日平均Earning/100000）*365
			//	stakingRewards := token.CurrentRewards
			//	stakingDays := int64(time.Now().Sub(*token.StakingTime).Hours()) / 24
			//	if stakingDays == 0 {
			//		stakingDays = 1
			//	}
			//	apr := stakingRewards.
			//		Div(decimal.NewFromInt(stakingDays)).
			//		Div(constant.MysteryBoxMintPrice).
			//		Mul(decimal.NewFromInt32(365))
			//	data.Apr = apr.InexactFloat64()
			//	data.Apy = math.Pow(1+(data.Apr/365), 365) - 1
			//}
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

func GetMintTx(requestTx string) (response *model.GetMintTxResponse, code model.ResponseCode, msg string) {
	job := model.RequestMintJob{
		RequestTxHash: requestTx,
	}
	if err := database.DB.Where(&job).First(&job).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, model.NotFound, "Can not find this request tx"
		} else {
			return nil, model.ServerInternalError, fmt.Sprintf("failed to get mint tx: %v", err)
		}
	}

	response = &model.GetMintTxResponse{
		RequestTx: requestTx,
		MintTx:    job.JobTxHash,
		JobMsg:    job.Result,
	}

	if job.Status == constant.MintJobStatusFailed && job.RetryCount < 3 {
		response.JobStatus = constant.MintJobStatusProcessing
	} else {
		response.JobStatus = job.Status
	}

	return response, model.Success, ""
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

func ClaimToken(user string, tokenIds []uint64) (response *model.ClaimBoxResponse, code model.ResponseCode, msg string) {
	claimedIdList, totalRewards, burnedRewards, txHash, err := eventlistener.ClaimMultipleTokenRewards(user, tokenIds)
	if err != nil {
		return nil, model.ServerInternalError, fmt.Sprintf("Failed to claim. err: %v", err)
	}

	if txHash == "" {
		return nil, model.ServerInternalError, fmt.Sprintf("Your transaction is being processed, the system will handle it automatically.")
	}

	var tokenList []model.MysteryBoxToken
	if err = database.DB.
		Where("token_id IN ?", claimedIdList).
		Find(&tokenList).
		Error; err != nil {
		log.Warnf("Failed to load claimedTokenIds: %v, err: %v", claimedIdList, err)
		return nil, model.ServerInternalError, fmt.Sprintf("Your transaction is being processed, the system will handle it automatically.")
	}

	response = &model.ClaimBoxResponse{
		TotalClaimed:    totalRewards,
		TotalBurned:     burnedRewards,
		TransactionHash: txHash,
		ClaimedBoxes:    tokenList,
	}

	return nil, model.Success, ""
}
