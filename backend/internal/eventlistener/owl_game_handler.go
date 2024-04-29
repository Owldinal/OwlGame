package eventlistener

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"math"
	"math/big"
	"owl-backend/internal/config"
	"owl-backend/internal/constant"
	"owl-backend/internal/database"
	"owl-backend/internal/model"
	"owl-backend/internal/util"
	"owl-backend/pkg/log"
	"time"
)

type OwlGameJoinGameHandler struct{}

var PrizeIncreasedForEachMint = decimal.NewFromInt(50000)

func (h *OwlGameJoinGameHandler) Handle(vlog types.Log) error {
	event, err := owlGameContract.ParseJoinGame(vlog)
	if err != nil {
		return err
	}
	// save event to database
	//log.Infof("[%v-%v] Mint box: user = %v, boxId = %v", event.Raw.TxHash, event.Raw.Index, event.User, event.TokenId.Uint64())
	eventItem := model.OwlGameJoinGameEvent{
		Event:      model.NewEvent(&event.Raw),
		User:       event.User.Hex(),
		InviteCode: event.InviteCode,
	}
	eventResult := database.DB.Clauses().Create(&eventItem)
	if eventResult.Error != nil {
		if errors.Is(eventResult.Error, gorm.ErrDuplicatedKey) {
			return nil
		} else {
			log.Warnf("Error Is: %v", eventResult.Error)
			return eventResult.Error
		}
	}

	item := model.UserInfo{
		Address:    eventItem.User,
		InviteCode: util.DecodeInviteCode(eventItem.InviteCode),
	}
	itemResult := database.DB.Clauses().Create(&item)
	if itemResult.Error != nil {
		return itemResult.Error
	}

	return nil
}

type OwlGameBindInvitationHandler struct{}

func (h *OwlGameBindInvitationHandler) Handle(vlog types.Log) error {
	event, err := owlGameContract.ParseBindInvitation(vlog)
	if err != nil {
		return err
	}
	// save event to database
	//log.Infof("[%v-%v] Mint box: user = %v, boxId = %v", event.Raw.TxHash, event.Raw.Index, event.User, event.TokenId.Uint64())
	eventItem := model.OwlGameBindInvitationEvent{
		Event:   model.NewEvent(&event.Raw),
		Inviter: event.Inviter.Hex(),
		Invitee: event.Invitee.Hex(),
	}
	eventResult := database.DB.Clauses().Create(&eventItem)
	if eventResult.Error != nil {
		if errors.Is(eventResult.Error, gorm.ErrDuplicatedKey) {
			return nil
		} else {
			log.Warnf("Error Is: %v", eventResult.Error)
			return eventResult.Error
		}
	}

	inviteItem := model.InviteRelation{
		Event:   eventItem.Event,
		Inviter: eventItem.Inviter,
		Invitee: eventItem.Invitee,
	}
	itemResult := database.DB.Clauses().Create(&inviteItem)
	if itemResult.Error != nil {
		return itemResult.Error
	}

	var userInfoItem model.UserInfo
	if err := database.DB.Where("address = ?", eventItem.Inviter).First(&userInfoItem).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Warnf("Inviter not found with address: %v", eventItem.Inviter)
			return err
		} else {
			log.Warnf("Error Is: %v", err)
			return err
		}
	}
	userInfoItem.InviteCount++
	if err := database.DB.Save(&userInfoItem).Error; err != nil {
		log.Warnf("Error updating inviter user infoo: %v", err)
		return err
	}

	return nil
}

type OwlGamePrizePoolIncreasedHandler struct{}

func (h *OwlGamePrizePoolIncreasedHandler) Handle(vlog types.Log) error {
	event, err := owlGameContract.ParsePrizePoolIncreased(vlog)
	if err != nil {
		return err
	}
	// save event to database
	// log.Infof("[%v-%v] Pool Increase: amount = %v\n", event.Raw.TxHash, event.Raw.Index, event.Amount)
	eventItem := model.OwlGamePrizePoolIncreasedEvent{
		Event:  model.NewEvent(&event.Raw),
		Amount: decimal.NewFromBigInt(event.Amount, -18),
	}
	eventResult := database.DB.Clauses().Create(&eventItem)
	if eventResult.Error != nil {
		if errors.Is(eventResult.Error, gorm.ErrDuplicatedKey) {
			return nil
		} else {
			log.Warnf("Error Is: %v", eventResult.Error)
			return eventResult.Error
		}
	}

	err = UpdateDailyPoolSnapshot(DailyPoolUpdater{Increase: eventItem.Amount})
	if err != nil {
		return err
	}

	return nil
}

type OwlGamePrizePoolDecreasedHandler struct{}

func (h *OwlGamePrizePoolDecreasedHandler) Handle(vlog types.Log) error {
	event, err := owlGameContract.ParsePrizePoolDecreased(vlog)
	if err != nil {
		return err
	}
	// save event to database
	//log.Infof("[%v-%v] Mint box: user = %v, boxId = %v", event.Raw.TxHash, event.Raw.Index, event.User, event.TokenId.Uint64())
	eventItem := model.OwlGamePrizePoolDecreasedEvent{
		Event:  model.NewEvent(&event.Raw),
		Amount: decimal.NewFromBigInt(event.Amount, -18),
	}
	eventResult := database.DB.Clauses().Create(&eventItem)
	if eventResult.Error != nil {
		if errors.Is(eventResult.Error, gorm.ErrDuplicatedKey) {
			return nil
		} else {
			log.Warnf("Error Is: %v", eventResult.Error)
			return eventResult.Error
		}
	}

	// 改成中心化计算收益之后，奖池的逻辑移动到了 UpdateFruitReward 方法里面，这里就不需要处理了（实际上应该也不会变）
	//err = UpdateDailyPoolSnapshot(DailyPoolUpdater{Decrease: eventItem.Amount})
	//if err != nil {
	//	return err
	//}

	return nil
}

type OwlGameRequestMintHandler struct{}

func (h *OwlGameRequestMintHandler) Handle(vlog types.Log) error {
	event, err := owlGameContract.ParseRequestMint(vlog)
	if err != nil {
		return err
	}

	eventItem := model.OwlGameRequestMintEvent{
		Event:     model.NewEvent(&event.Raw),
		User:      event.User.Hex(),
		Count:     event.Count.Uint64(),
		RequestId: event.RequestId.Uint64(),
	}
	eventResult := database.DB.Clauses().Create(&eventItem)
	if eventResult.Error != nil {
		if errors.Is(eventResult.Error, gorm.ErrDuplicatedKey) {
			return nil
		} else {
			log.Warnf("Error Is: %v", eventResult.Error)
			return eventResult.Error
		}
	}

	job := &model.RequestMintJob{
		RequestTxHash:      eventItem.TxHash,
		RequestLogIndex:    eventItem.LogIndex,
		RequestBlockNumber: eventItem.BlockNumber,
		RequestBlockHash:   eventItem.BlockHash,

		RequestId: eventItem.RequestId,
		User:      eventItem.User,
		Count:     eventItem.Count,
		Status:    constant.MintJobStatusNotStart,
	}

	itemResult := database.DB.Clauses().Create(&job)
	if itemResult.Error != nil {
		return itemResult.Error
	}

	return nil
}

type OwlGameMintMysteryBoxHandler struct{}

func (h *OwlGameMintMysteryBoxHandler) Handle(vlog types.Log) error {
	event, err := owlGameContract.ParseMintMysteryBox(vlog)
	if err != nil {
		return err
	}
	// save event to database
	//log.Infof("[%v-%v] Mint box: user = %v, boxId = %v", event.Raw.TxHash, event.Raw.Index, event.User, event.TokenId.Uint64())
	tokenIds := make([]uint64, len(event.TokenId))
	for i, bigInt := range event.TokenId {
		tokenIds[i] = bigInt.Uint64()
	}
	eventItem := model.OwlGameMintMysteryBoxEvent{
		Event:     model.NewEvent(&event.Raw),
		User:      event.User.Hex(),
		Count:     event.Count.Uint64(),
		RequestId: event.RequestId.Uint64(),
		TokenIds:  tokenIds,
	}
	eventResult := database.DB.Clauses().Create(&eventItem)
	if eventResult.Error != nil {
		if errors.Is(eventResult.Error, gorm.ErrDuplicatedKey) {
			return nil
		} else {
			log.Warnf("Error Is: %v", eventResult.Error)
			return eventResult.Error
		}
	}

	// update job
	job := model.RequestMintJob{
		RequestId: eventItem.RequestId,
	}
	if err := database.DB.Where(&job).First(&job).Error; err != nil {
		log.Warnf("Failed to load job %v: %v", eventItem.RequestId, err)
		//return err
	}
	job.HasConfirmed = true
	if err := database.DB.Model(&job).Select("HasConfirmed").Updates(job).Error; err != nil {
		log.Warnf("Failed to confirm job %v: %v", eventItem.RequestId, err)
		//return err
	}

	// insert poll transaction
	transactionDetail := &model.RewardPoolTransactionRecord{
		User:        eventItem.User,
		Operation:   "Mint",
		Description: "Gen1 Mystery Box",
		Count:       eventItem.Count,
		Amount:      decimal.NewFromInt(int64(eventItem.Count)).Mul(PrizeIncreasedForEachMint),
		Event:       eventItem.Event,
	}

	itemResult := database.DB.Clauses().Create(&transactionDetail)
	if itemResult.Error != nil {
		return itemResult.Error
	}

	return nil
}

type OwlGameStakeOwldinalNftHandler struct{}

func (h *OwlGameStakeOwldinalNftHandler) Handle(vlog types.Log) error {
	event, err := owlGameContract.ParseStakeOwldinalNft(vlog)
	if err != nil {
		return err
	}
	// save event to database
	//log.Infof("[%v-%v] Mint box: user = %v, boxId = %v", event.Raw.TxHash, event.Raw.Index, event.User, event.TokenId.Uint64())
	tokenIds := make([]uint64, len(event.TokenId))
	for i, bigInt := range event.TokenId {
		tokenIds[i] = bigInt.Uint64()
	}
	eventItem := model.OwlGameStakeOwldinalNftEvent{
		Event:    model.NewEvent(&event.Raw),
		User:     event.User.Hex(),
		TokenIds: tokenIds,
	}
	eventResult := database.DB.Clauses().Create(&eventItem)
	if eventResult.Error != nil {
		if errors.Is(eventResult.Error, gorm.ErrDuplicatedKey) {
			return nil
		} else {
			log.Warnf("Error Is: %v", eventResult.Error)
			return eventResult.Error
		}
	}

	updateResult := database.DB.Model(&model.OwldinalNftToken{}).
		Where("token_id IN ?", tokenIds).
		Update("is_staking", true).
		Update("staking_time", time.Now())
	if updateResult.Error != nil {
		return updateResult.Error
	}

	userInfo := model.UserInfo{
		Address: eventItem.User,
	}
	if err := database.DB.Where(&userInfo).First(&userInfo).Error; err != nil {
		log.Warnf("Error Is: %v", err)
		return err
	}

	userInfo.BuffLevel += len(tokenIds)
	if err := database.DB.Save(&userInfo).Error; err != nil {
		log.Warnf("Error updating inviter user infoo: %v", err)
		return err
	}

	return nil
}

type OwlGameUnstakeOwldinalNftHandler struct{}

func (h *OwlGameUnstakeOwldinalNftHandler) Handle(vlog types.Log) error {
	event, err := owlGameContract.ParseUnstakeOwldinalNft(vlog)
	if err != nil {
		return err
	}
	// save event to database
	//log.Infof("[%v-%v] Mint box: user = %v, boxId = %v", event.Raw.TxHash, event.Raw.Index, event.User, event.TokenId.Uint64())
	tokenIds := make([]uint64, len(event.TokenId))
	for i, bigInt := range event.TokenId {
		tokenIds[i] = bigInt.Uint64()
	}
	eventItem := model.OwlGameUnstakeOwldinalNftEvent{
		Event:    model.NewEvent(&event.Raw),
		User:     event.User.Hex(),
		TokenIds: tokenIds,
	}
	eventResult := database.DB.Clauses().Create(&eventItem)
	if eventResult.Error != nil {
		if errors.Is(eventResult.Error, gorm.ErrDuplicatedKey) {
			return nil
		} else {
			log.Warnf("Error Is: %v", eventResult.Error)
			return eventResult.Error
		}
	}

	updateResult := database.DB.Model(&model.OwldinalNftToken{}).
		Where("token_id IN ?", tokenIds).
		Update("is_staking", false).
		Update("staking_time", nil)

	if updateResult.Error != nil {
		return updateResult.Error
	}

	userInfo := model.UserInfo{
		Address: eventItem.User,
	}
	if err := database.DB.Where(&userInfo).First(&userInfo).Error; err != nil {
		log.Warnf("Error Is: %v", err)
		return err
	}

	userInfo.BuffLevel -= len(tokenIds)
	if err := database.DB.Save(&userInfo).Error; err != nil {
		log.Warnf("Error updating inviter user infoo: %v", err)
		return err
	}

	return nil
}

type OwlGameStakeMysteryBoxHandler struct{}

func (h *OwlGameStakeMysteryBoxHandler) Handle(vlog types.Log) error {
	event, err := owlGameContract.ParseStakeMysteryBox(vlog)
	if err != nil {
		return err
	}
	// save event to database
	//log.Infof("[%v-%v] Mint box: user = %v, boxId = %v", event.Raw.TxHash, event.Raw.Index, event.User, event.TokenId.Uint64())
	tokenIds := make([]uint64, len(event.TokenId))
	for i, bigInt := range event.TokenId {
		tokenIds[i] = bigInt.Uint64()
	}
	eventItem := model.OwlGameStakeMysteryBoxEvent{
		Event:     model.NewEvent(&event.Raw),
		User:      event.User.Hex(),
		TokenIds:  tokenIds,
		BuffLevel: uint8(event.BuffLevel),
	}
	eventResult := database.DB.Clauses().Create(&eventItem)
	if eventResult.Error != nil {
		if errors.Is(eventResult.Error, gorm.ErrDuplicatedKey) {
			return nil
		} else {
			log.Warnf("Error Is: %v", eventResult.Error)
			return eventResult.Error
		}
	}

	updateResult := database.DB.Model(&model.MysteryBoxToken{}).
		Where("token_id IN ?", tokenIds).
		Update("is_staking", true).
		Update("buff_level", eventItem.BuffLevel).
		Update("staking_time", time.Now())
	if updateResult.Error != nil {
		return updateResult.Error
	}

	return nil
}

type OwlGameUnstakeMysteryBoxHandler struct{}

func (h *OwlGameUnstakeMysteryBoxHandler) Handle(vlog types.Log) error {
	event, err := owlGameContract.ParseUnstakeMysteryBox(vlog)
	if err != nil {
		return err
	}
	// save event to database
	//log.Infof("[%v-%v] Mint box: user = %v, boxId = %v", event.Raw.TxHash, event.Raw.Index, event.User, event.TokenId.Uint64())
	eventItem := model.OwlGameUnstakeMysteryBoxEvent{
		Event:   model.NewEvent(&event.Raw),
		User:    event.User.Hex(),
		TokenId: event.TokenId.Uint64(),
		Rewards: decimal.NewFromBigInt(event.Rewards, -18),
		BoxType: event.BoxType,
	}
	eventResult := database.DB.Clauses().Create(&eventItem)
	if eventResult.Error != nil {
		if errors.Is(eventResult.Error, gorm.ErrDuplicatedKey) {
			return nil
		} else {
			log.Warnf("Error Is: %v", eventResult.Error)
			return eventResult.Error
		}
	}

	tokenItem := model.MysteryBoxToken{
		TokenId: eventItem.TokenId,
	}
	if err := database.DB.Where(&tokenItem).First(&tokenItem).Error; err != nil {
		log.Warnf("Error Is: %v", err)
		return err
	}
	//if !eventItem.Rewards.Equal(tokenItem.CurrentRewards) {
	// 这里不需要了，因为合约的奖励总是0
	//log.Warnf("Cliam: rewards not equal as db. event=%v, db=%v", eventItem.Rewards, tokenItem.CurrentRewards)
	//}

	tokenItem.IsStaking = false
	tokenItem.StakingTime = nil

	if tokenItem.CurrentRewards.IsZero() {
		if err := database.DB.Save(&tokenItem).Error; err != nil {
			log.Warnf("Error updating unstake box: %v", err)
			return err
		}

		return nil
	}

	// Has rewards. do transfer
	isMoonBoost, buffLevel, err := getBuffStatus(tokenItem.Owner)
	if err != nil {
		return err
	}
	rewardProportion := getTokenRewardProportion(tokenItem.BoxType, buffLevel, isMoonBoost)
	actualRewards, burnRewards, err := calculateRewards(tokenItem.CurrentRewards, rewardProportion)
	if err != nil {
		log.Warnf("Failed to calculateRewards, token=%+v, err=%v", tokenItem, err)
	}

	tokenItem.ClaimedRewards = tokenItem.ClaimedRewards.Add(actualRewards)
	tokenItem.CurrentRewards = decimal.Zero

	if err := database.DB.Save(&tokenItem).Error; err != nil {
		log.Warnf("Error updating unstake box: %v", err)
		return err
	}

	err = claimSingleTokenRewards(&tokenItem, &eventItem, actualRewards, burnRewards, buffLevel, isMoonBoost)
	if err != nil {
		return err
	}

	return nil
}

func claimSingleTokenRewards(
	tokenItem *model.MysteryBoxToken,
	eventItem *model.OwlGameUnstakeMysteryBoxEvent,
	actualRewards decimal.Decimal,
	burnRewards decimal.Decimal,
	buffLevel uint8,
	isMoonBoost bool,
) error {
	userInfo := model.UserInfo{
		Address: eventItem.User,
	}
	if err := database.DB.Where(&userInfo).First(&userInfo).Error; err != nil {
		return err
	}
	userInfo.TotalEarned = userInfo.TotalEarned.Add(actualRewards)
	if err := database.DB.Save(&userInfo).Error; err != nil {
		return err
	}

	transferRecord := model.TransferRewards{
		User:          tokenItem.Owner,
		TokenId:       tokenItem.TokenId,
		BoxType:       tokenItem.BoxType,
		Rewards:       actualRewards,
		BurnedRewards: burnRewards,
		BuffLevel:     buffLevel,
		MoonBoost:     isMoonBoost,
		//IsConfirmed:   false,

		ClaimTxHash:      eventItem.TxHash,
		ClaimLogIndex:    eventItem.LogIndex,
		ClaimBlockHash:   eventItem.BlockHash,
		ClaimBlockNumber: eventItem.BlockNumber,
	}

	if err := database.DB.Create(&transferRecord).Error; err != nil {
		log.Warnf("Error Create transferRecordr: %v", err)
		return err
	}

	txHash, blockHash, blockNumber, err := transferRewardsToUser(tokenItem, actualRewards)
	if err != nil {
		transferRecord.Result = transferRecord.Result + fmt.Sprintf("Error : %v;", err)
		// 不需要 return，这个可以之后再重试，先进行后续的处理
	} else {
		transferRecord.TransferTxHash = txHash
		transferRecord.TransferBlockNumber = blockNumber
		transferRecord.TransferBlockHash = blockHash
		transferRecord.Result = transferRecord.Result + fmt.Sprintf("Success;")
	}

	if err := database.DB.Save(&transferRecord).Error; err != nil {
		log.Warnf("Error update transferRecord for transfer: %v", err)
		return err
	}

	if tokenItem.BoxType == constant.BoxTypeFruit {
		// burnRewards 是需要发给所有妖精的奖励

		// 这里 unstake 多个的时候，会有多个事件同时发出来，所以这里加上事务以防修改的时候冲突
		tx := database.DB.Begin()
		var elfTokens []model.MysteryBoxToken
		if err := tx.Set("gorm:query_option", "FOR UPDATE").
			Where("is_staking = ? AND box_type = ?", true, constant.BoxTypeElf).
			Find(&elfTokens).
			Error; err != nil {
			tx.Rollback()
			log.Warnf("Error finding MysteryBoxToken with staking: %v", err)
			return err
		}

		eachElfRewards := burnRewards.Div(decimal.NewFromInt(int64(len(elfTokens))))
		for _, elf := range elfTokens {
			elf.CurrentRewards = elf.CurrentRewards.Add(eachElfRewards)
			elf.TotalRewards = elf.TotalRewards.Add(eachElfRewards)
			if err := tx.Save(&elf).Error; err != nil {
				tx.Rollback()
				log.Warnf("Error saving updated MysteryBoxToken: %v", err)
				return err
			}
		}

		if err := tx.Commit().Error; err != nil {
			log.Warnf("Error committing transaction: %v", err)
			return err
		}

	} else if tokenItem.BoxType == constant.BoxTypeElf {
		burnTxHash, _, _, err := burnOwlToken(burnRewards)
		if err != nil {
			transferRecord.Result = transferRecord.Result + fmt.Sprintf("BurnError : %v;", err)
		} else {
			transferRecord.Result = transferRecord.Result + fmt.Sprintf("BurnSuccess;")
		}
		transferRecord.BurnTxHash = burnTxHash
		if err := database.DB.Save(&transferRecord).Error; err != nil {
			log.Warnf("Error update transferRecord for burna: %v", err)
			return err
		}

		err = UpdateDailyPoolSnapshot(DailyPoolUpdater{Burn: burnRewards})
		if err != nil {
			return err
		}
	}

	return nil
}

func getBuffStatus(user string) (hasMoonBoost bool, owlBuffLevel uint8, err error) {
	// buff level not correct, should check owldinal nft
	var nftTokens []model.OwldinalNftToken
	result := database.DB.Where("owner = ?", user).Find(&nftTokens)
	if result.Error != nil {
		err = result.Error
		return
	}
	if config.C.NeedCheckMoonBoost {
		hasMoonBoost = len(nftTokens) > 0 || constant.MoonBoostAddress[user]
	}
	for _, nft := range nftTokens {
		if nft.IsStaking {
			owlBuffLevel++
		}
	}

	return
}

func getTokenRewardProportion(boxType constant.BoxType, buffLevel uint8, isMoonBoost bool) (rewardProportion int64) {
	if boxType == constant.BoxTypeFruit {
		rewardProportion = 75
		if buffLevel >= 3 {
			rewardProportion = 85
		}
		if isMoonBoost {
			rewardProportion += 7
		}
	} else if boxType == constant.BoxTypeElf {
		rewardProportion = 85
		if buffLevel >= 2 {
			rewardProportion = 90
		}
		if isMoonBoost {
			rewardProportion += 7
		}
	}

	return
}

func calculateRewards(originReward decimal.Decimal, rewardProportion int64) (
	actualRewards decimal.Decimal,
	burnRewards decimal.Decimal,
	err error,
) {
	actualRewards = originReward.Mul(decimal.NewFromInt(rewardProportion)).Div(decimal.NewFromInt(100))
	burnRewards = originReward.Sub(actualRewards)

	return
}

func transferRewardsToUser(
	token *model.MysteryBoxToken, reward decimal.Decimal,
) (txHash string, blockHash string, blockNumber uint64, err error) {
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(config.C.ChainId))
	if err != nil {
		return
	}

	targetAddress := common.HexToAddress(token.Owner)
	targetCount := new(big.Int)
	targetCount.SetString(reward.Mul(decimal.NewFromFloat(1e18)).StringFixed(0), 10)

	if config.C.GasPrice > 0 {
		auth.GasPrice = big.NewInt(config.C.GasPrice)
	}

	var tx *types.Transaction
	tx, err = owlTokenContract.Transfer(auth, targetAddress, targetCount)
	if err != nil {
		return
	}
	txHash = tx.Hash().Hex()
	receipt, err := bind.WaitMined(context.Background(), ethClient, tx)
	blockHash = receipt.BlockHash.Hex()
	blockNumber = receipt.BlockNumber.Uint64()

	return
}

func burnOwlToken(
	burnAmount decimal.Decimal,
) (txHash string, blockHash string, blockNumber uint64, err error) {
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(config.C.ChainId))
	if err != nil {
		return
	}

	targetAddress := common.HexToAddress(constant.BurnAddr)
	targetCount := new(big.Int)
	targetCount.SetString(burnAmount.Mul(decimal.NewFromFloat(1e18)).StringFixed(0), 10)

	if config.C.GasPrice > 0 {
		auth.GasPrice = big.NewInt(config.C.GasPrice)
	}

	var tx *types.Transaction
	tx, err = owlTokenContract.Transfer(auth, targetAddress, targetCount)
	if err != nil {
		return
	}
	txHash = tx.Hash().Hex()
	receipt, err := bind.WaitMined(context.Background(), ethClient, tx)
	blockHash = receipt.BlockHash.Hex()
	blockNumber = receipt.BlockNumber.Uint64()

	return
}

type OwlGameOwlTokenBurnedHandler struct{}

func (h *OwlGameOwlTokenBurnedHandler) Handle(vlog types.Log) error {
	event, err := owlGameContract.ParseOwlTokenBurned(vlog)
	if err != nil {
		return err
	}
	// save event to database
	//log.Infof("[%v-%v] Mint box: user = %v, boxId = %v", event.Raw.TxHash, event.Raw.Index, event.User, event.TokenId.Uint64())
	eventItem := model.OwlGameOwlTokenBurnedEvent{
		Event:     model.NewEvent(&event.Raw),
		User:      event.User.Hex(),
		MintCount: event.MintCount.Uint64(),
		Amount:    decimal.NewFromBigInt(event.Amount, -18),
	}
	eventResult := database.DB.Clauses().Create(&eventItem)
	if eventResult.Error != nil {
		if errors.Is(eventResult.Error, gorm.ErrDuplicatedKey) {
			return nil
		} else {
			log.Warnf("Error Is: %v", eventResult.Error)
			return eventResult.Error
		}
	}

	err = UpdateDailyPoolSnapshot(DailyPoolUpdater{Burn: eventItem.Amount})
	if err != nil {
		return err
	}

	return nil
}

type OwlGameFruitRewardUpdatedHandler struct{}
type OwlGameElfRewardUpdatedHandler struct{}

func (h *OwlGameFruitRewardUpdatedHandler) Handle(vlog types.Log) error {
	event, err := owlGameContract.ParseFruitRewardUpdated(vlog)
	if err != nil {
		return err
	}
	// save event to database
	//log.Infof("[%v-%v] Mint box: user = %v, boxId = %v", event.Raw.TxHash, event.Raw.Index, event.User, event.TokenId.Uint64())
	eventItem := model.OwlGameFruitRewardUpdateEvent{
		Event:           model.NewEvent(&event.Raw),
		Count:           event.Count.Uint64(),
		TotalFruitCount: event.TotalFruitCount.Uint64(),
		TotalElfCount:   event.TotalElfCount.Uint64(),
		Amount:          decimal.NewFromBigInt(event.Amount, -18),
	}
	eventResult := database.DB.Clauses().Create(&eventItem)
	if eventResult.Error != nil {
		if errors.Is(eventResult.Error, gorm.ErrDuplicatedKey) {
			return nil
		} else {
			log.Warnf("Error Is: %v", eventResult.Error)
			return eventResult.Error
		}
	}

	UpdateAprSnapshot(&eventItem)
	return globalUpdateRewards(constant.BoxTypeFruit, eventItem.Count, eventItem.Amount)
}

func (h *OwlGameElfRewardUpdatedHandler) Handle(vlog types.Log) error {
	event, err := owlGameContract.ParseElfRewardUpdated(vlog)
	if err != nil {
		return err
	}
	// save event to database
	//log.Infof("[%v-%v] Mint box: user = %v, boxId = %v", event.Raw.TxHash, event.Raw.Index, event.User, event.TokenId.Uint64())
	eventItem := model.OwlGameElfRewardUpdateEvent{
		Event:  model.NewEvent(&event.Raw),
		Count:  event.Count.Uint64(),
		Amount: decimal.NewFromBigInt(event.Amount, -18),
	}
	eventResult := database.DB.Clauses().Create(&eventItem)
	if eventResult.Error != nil {
		if errors.Is(eventResult.Error, gorm.ErrDuplicatedKey) {
			return nil
		} else {
			log.Warnf("Error Is: %v", eventResult.Error)
			return eventResult.Error
		}
	}

	return globalUpdateRewards(constant.BoxTypeElf, eventItem.Count, eventItem.Amount)
}

func globalUpdateRewards(boxType constant.BoxType, count uint64, amount decimal.Decimal) error {
	var tokenList []model.MysteryBoxToken
	err := database.DB.Where("box_type = ?", boxType).Where("is_staking = ?", true).Find(&tokenList).Error

	if err != nil {
		return err
	}

	tx := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, token := range tokenList {
		tokenId := big.NewInt(int64(token.TokenId))
		tokenInfo, err := owlGameContract.GetTokenInfo(&bind.CallOpts{}, tokenId)
		if err != nil {
			log.Warnf("error when load token info for token id %v", tokenId)
			tx.Rollback()
			return err
		}

		currentRewards := decimal.NewFromBigInt(tokenInfo.Reward, -18)
		if currentRewards.Cmp(token.CurrentRewards) > 0 {
			addRewards := currentRewards.Sub(token.CurrentRewards)
			token.CurrentRewards = currentRewards
			token.TotalRewards = token.TotalRewards.Add(addRewards)

			if err := tx.Save(&token).Error; err != nil {
				log.Warnf("error when update rewards for token id %v", tokenId)
				tx.Rollback()
				return err
			}
		}
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func UpdateAprSnapshot(fruitRewardEvent *model.OwlGameFruitRewardUpdateEvent) {
	snapshot := model.AprSnapshot{
		Event:           fruitRewardEvent.Event,
		RewardCount:     fruitRewardEvent.Count,
		TotalFruitCount: fruitRewardEvent.TotalFruitCount,
		TotalElfCount:   fruitRewardEvent.TotalElfCount,
		FruitRewards:    fruitRewardEvent.Amount,
	}

	// Fruit APR 计算公式： 1- 上次4h给每个水果发了多少奖励X ; 2- 水果成本 10w
	// APR= 6X / 100000 * 365
	// APY = (1 + APR/365) ^ 365 - 1
	fruitRewards := fruitRewardEvent.Amount
	fruitAprDecimal := fruitRewards.
		Mul(decimal.NewFromInt(6)).
		Div(decimal.NewFromInt(100000)).
		Mul(decimal.NewFromInt(365))
	fruitApr := fruitAprDecimal.InexactFloat64()
	fruitApy := math.Pow(1+(fruitApr/365), 365) - 1

	snapshot.FruitApr = fruitApr
	snapshot.FruitApy = fruitApy

	// Elf APR 计算公式：
	// 1- 上次4h给所有水果发了多少奖励Z
	// 2- 妖精成本10w
	// 3- 上次4h总质押的妖精数量Y
	// APR=（16%*Z*6）/Y/100000 *365
	// APY = (1 + APR/365) ^ 365 - 1
	if int64(fruitRewardEvent.TotalElfCount) > 0 {
		totalFruitRewards := fruitRewardEvent.Amount.Mul(decimal.NewFromInt(int64(fruitRewardEvent.Count)))
		elfAprDecimal := totalFruitRewards.
			Mul(decimal.NewFromInt(6)).
			Mul(decimal.NewFromFloat32(0.16)).
			Div(decimal.NewFromInt(int64(fruitRewardEvent.TotalElfCount))).
			Div(decimal.NewFromInt(100000)).
			Mul(decimal.NewFromInt(365))
		elfApr := elfAprDecimal.InexactFloat64()
		elfApy := math.Pow(1+(elfApr/365), 365) - 1

		snapshot.ElfApr = elfApr
		snapshot.ElfApy = elfApy
	}

	if math.IsInf(snapshot.FruitApy, 1) {
		log.Warnf("apr_snapshots: fruit apy is inf. %+v", snapshot)
		snapshot.FruitApy = math.MaxFloat64
	}

	if math.IsInf(snapshot.ElfApy, 1) {
		log.Warnf("apr_snapshots: elf apy is inf. %+v", snapshot)
		snapshot.ElfApy = math.MaxFloat64
	}

	if err := database.DB.Create(&snapshot).Error; err != nil {
		log.Warnf("Error saving updateAprSnapshot: %v", err)
	}
}

type DailyPoolUpdater struct {
	Increase  decimal.Decimal
	Decrease  decimal.Decimal
	Burn      decimal.Decimal
	MarketCap decimal.Decimal
}

func UpdateDailyPoolSnapshot(updater DailyPoolUpdater) error {
	currentDate := time.Now().UTC().Truncate(24 * time.Hour)
	var snapshot model.DailyPoolSnapshot
	// Get newest data
	if err := database.DB.Order("id desc").First(&snapshot).Error; err != nil {
		log.Warnf("Error retrieving the latest DailyPoolSnapshot: %v", err)
		return err
	}

	// check if is today's snapshot
	if snapshot.Date != currentDate {
		// if not today's snapshot, create new one
		snapshot.ID = 0
		snapshot.Date = currentDate
	}

	if !updater.Increase.IsZero() {
		snapshot.TotalPoolAmount = snapshot.TotalPoolAmount.Add(updater.Increase)
	}

	if !updater.Decrease.IsZero() {
		snapshot.TotalPoolAmount = snapshot.TotalPoolAmount.Sub(updater.Decrease)
		snapshot.AllocatedRewards = snapshot.AllocatedRewards.Add(updater.Decrease)
	}

	if !updater.Burn.IsZero() {
		snapshot.TotalBurn = snapshot.TotalBurn.Add(updater.Burn)
	}

	if !updater.MarketCap.IsZero() {
		snapshot.TotalMarketCap = snapshot.TotalMarketCap.Add(updater.MarketCap)
	}

	if err := database.DB.Save(&snapshot).Error; err != nil {
		log.Warnf("Error saving DailyPoolSnapshot: %v", err)
		return err
	}

	return nil
}
