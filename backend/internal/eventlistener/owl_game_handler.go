package eventlistener

import (
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"math/big"
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

	err = updateDailyPoolSnapshot(DailyPoolUpdater{Increase: eventItem.Amount})
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

	err = updateDailyPoolSnapshot(DailyPoolUpdater{Decrease: eventItem.Amount})
	if err != nil {
		return err
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
		Event:    model.NewEvent(&event.Raw),
		User:     event.User.Hex(),
		Count:    event.Count.Int64(),
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

	transactionDetail := &model.RewardPoolTransactionRecord{
		User:        eventItem.User,
		Operation:   "Mint",
		Description: "Gen1 Mystery Box",
		Count:       eventItem.Count,
		Amount:      decimal.NewFromInt(eventItem.Count).Mul(PrizeIncreasedForEachMint),
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

	updateResult := database.DB.Model(&model.MysteryBoxToken{}).
		Where("token_id IN ?", tokenIds).
		Update("is_staking", true).
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
	tokenItem.IsStaking = false
	tokenItem.StakingTime = nil
	if err := database.DB.Save(&tokenItem).Error; err != nil {
		log.Warnf("Error updating inviter user infoo: %v", err)
		return err
	}

	return nil
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

	err = updateDailyPoolSnapshot(DailyPoolUpdater{Burn: eventItem.Amount})
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
			tx.Rollback()
			return err
		}

		currentRewards := decimal.NewFromBigInt(tokenInfo.Reward, -18)
		if currentRewards.Cmp(token.CurrentRewards) > 0 {
			addRewards := currentRewards.Sub(token.CurrentRewards)
			token.CurrentRewards = currentRewards
			token.TotalRewards.Add(addRewards)

			if err := tx.Save(&token).Error; err != nil {
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

type DailyPoolUpdater struct {
	Increase  decimal.Decimal
	Decrease  decimal.Decimal
	Burn      decimal.Decimal
	MarketCap decimal.Decimal
}

func updateDailyPoolSnapshot(updater DailyPoolUpdater) error {
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
