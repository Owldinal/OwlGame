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
	eventItem := model.OwlGameClaimRebateClaimedEvent{
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
