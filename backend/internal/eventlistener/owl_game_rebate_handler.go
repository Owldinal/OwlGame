package eventlistener

import (
	"errors"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"owl-backend/internal/database"
	"owl-backend/internal/model"
	"owl-backend/pkg/log"
)

type OwlGameRebateRewardsIncreasedHandler struct{}

func (h *OwlGameRebateRewardsIncreasedHandler) Handle(vlog types.Log) error {
	event, err := owlGameContract.ParseRebateRewardsIncreased(vlog)
	if err != nil {
		return err
	}
	// save event to database
	//log.Infof("[%v-%v] Mint box: user = %v, boxId = %v", event.Raw.TxHash, event.Raw.Index, event.User, event.TokenId.Uint64())
	eventItem := model.OwlGameRebateRewardsIncreasedEvent{
		Event:     model.NewEvent(&event.Raw),
		User:      event.User.Hex(),
		Invitee:   event.Invitee.Hex(),
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

	// Update user's unclaimedReferral
	var userInfoItem model.UserInfo
	if err := database.DB.Where("address = ?", eventItem.User).First(&userInfoItem).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Warnf("Inviter not found with address: %v", eventItem.User)
			return err
		} else {
			log.Warnf("Error Is: %v", err)
			return err
		}
	}
	userInfoItem.UnclaimedReferral = userInfoItem.UnclaimedReferral.Add(eventItem.Amount)
	if err := database.DB.Save(&userInfoItem).Error; err != nil {
		log.Warnf("Error updating inviter user infoo: %v", err)
		return err
	}

	inviteRelationItem := model.InviteRelation{
		Inviter: eventItem.User,
		Invitee: eventItem.Invitee,
	}
	if err := database.DB.Where(&inviteRelationItem).First(&inviteRelationItem).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Warnf("Inviter not found with address: %v", eventItem.User)
			return err
		} else {
			log.Warnf("Error Is: %v", err)
			return err
		}
	}
	inviteRelationItem.ReferralRewards = inviteRelationItem.ReferralRewards.Add(eventItem.Amount)
	if err := database.DB.Save(&inviteRelationItem).Error; err != nil {
		log.Warnf("Error updating inviter user infoo: %v", err)
		return err
	}

	return nil
}

type OwlGameUnlockableRebateIncreasedHandler struct{}

func (h *OwlGameUnlockableRebateIncreasedHandler) Handle(vlog types.Log) error {
	event, err := owlGameContract.ParseUnlockableRebateIncreased(vlog)
	if err != nil {
		return err
	}
	// save event to database
	//log.Infof("[%v-%v] Mint box: user = %v, boxId = %v", event.Raw.TxHash, event.Raw.Index, event.User, event.TokenId.Uint64())

	// 这里犯了个错误，上线的时候对于这条事件写错表了。但因为后续对 user_info 表的修改是正常的，所以目前不影响使用。
	// 如果在不提前修复数据库事件的前提下，将这里改对，会导致数据重复入库，导致用户解锁量加倍
	// 为了避免引入新问题，这里暂不做修改，将错就错。
	// 如果需要修复，那么需要分两步，首先只索引历史事件而不进行后续处理，然后对新的事件才能正常进行处理。
	eventItem := model.OwlGameRebateRewardsIncreasedEvent{
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

	var userInfoItem model.UserInfo
	if err := database.DB.Where("address = ?", eventItem.User).First(&userInfoItem).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Warnf("Inviter not found with address: %v", eventItem.User)
			return err
		} else {
			log.Warnf("Error Is: %v", err)
			return err
		}
	}

	userInfoItem.UnlockableReferral = userInfoItem.UnlockableReferral.Add(eventItem.Amount)
	if err := database.DB.Save(&userInfoItem).Error; err != nil {
		log.Warnf("Error updating inviter user infoo: %v", err)
		return err
	}

	return nil
}

type OwlGameClaimRebateClaimedHandler struct{}

func (h *OwlGameClaimRebateClaimedHandler) Handle(vlog types.Log) error {
	event, err := owlGameContract.ParseRebateClaimed(vlog)
	if err != nil {
		return err
	}
	// save event to database
	//log.Infof("[%v-%v] Mint box: user = %v, boxId = %v", event.Raw.TxHash, event.Raw.Index, event.User, event.TokenId.Uint64())
	eventItem := model.OwlGameRebateClaimedEvent{
		Event:  model.NewEvent(&event.Raw),
		User:   event.User.Hex(),
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

	userInfo := model.UserInfo{
		Address: eventItem.User,
	}
	if err := database.DB.Where(&userInfo).First(&userInfo).Error; err != nil {
		return err
	}

	userInfo.TotalEarned = userInfo.TotalEarned.Add(eventItem.Amount)
	userInfo.UnlockableReferral = userInfo.UnlockableReferral.Sub(eventItem.Amount)
	userInfo.UnclaimedReferral = userInfo.UnclaimedReferral.Sub(eventItem.Amount)
	userInfo.ClaimedReferral = userInfo.ClaimedReferral.Add(eventItem.Amount)

	if err := database.DB.Save(&userInfo).Error; err != nil {
		return err
	}

	return nil
}
