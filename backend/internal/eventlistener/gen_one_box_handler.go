package eventlistener

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"gorm.io/gorm"
	"owl-backend/internal/constant"
	"owl-backend/internal/database"
	"owl-backend/internal/model"
	"owl-backend/pkg/log"
)

type GenOneBoxMintBoxHandler struct{}

func (h *GenOneBoxMintBoxHandler) Handle(vlog types.Log) error {
	event, err := genOneBoxContract.ParseMintBox(vlog)
	if err != nil {
		return err
	}
	// save event to database
	//log.Infof("[%v-%v] Mint MysteryBox: user = %v, boxId = %v", event.Raw.TxHash, event.Raw.Index, event.User, event.TokenId.Uint64())
	eventItem := model.GenOneBoxMintBoxEvent{
		Event:   model.NewEvent(&event.Raw),
		User:    event.User.Hex(),
		TokenId: event.TokenId.Uint64(),
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

	// Save into db
	item := model.MysteryBoxToken{
		TokenId: eventItem.TokenId,
		Owner:   eventItem.User,
		BoxType: constant.BoxType(eventItem.BoxType),
	}

	if item.BoxType == constant.BoxTypeBurned {
		// burned token, ignore it.
		return nil
	}

	itemResult := database.DB.Clauses().Create(&item)
	if itemResult.Error != nil {
		return itemResult.Error
	}

	return nil
}

type GenOneBoxTransferHandler struct{}

func (h *GenOneBoxTransferHandler) Handle(vlog types.Log) error {
	event, err := genOneBoxContract.ParseTransfer(vlog)
	if err != nil {
		return err
	}
	//log.Infof("[%v-%v] Transfer MysteryBox from = %v , to = %v , box= %v", event.Raw.TxHash, event.Raw.Index, event.From, event.To, event.TokenId)

	// save event to database
	eventItem := model.GenOneBoxTransferEvent{
		Event:    model.NewEvent(&event.Raw),
		FromUser: event.From.Hex(),
		ToUser:   event.To.Hex(),
		TokenId:  event.TokenId.Uint64(),
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

	// When FromUser=0x0， means this is a mint. don't update db (mint will do this.)
	if eventItem.FromUser == constant.NoneAddr {
		return nil
	}

	// When ToUser=0x0， means this is a burn. don't update db
	if eventItem.ToUser == constant.NoneAddr {
		return nil
	}

	// is staking or unstaking, handle it in owl_game_handler
	if eventItem.FromUser == OwlGameAddr || eventItem.ToUser == OwlGameAddr {
		return nil
	}

	var token model.MysteryBoxToken
	if err := database.DB.Where("token_id = ?", eventItem.TokenId).First(&token).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Warnf("MysteryBox not found with id: %v", eventItem.TokenId)
			return err
		} else {
			log.Warnf("Error Is: %v", err)
			return err
		}
	}

	// wrong owner
	if token.Owner != eventItem.FromUser {
		err := fmt.Errorf("MysteryBox owner mismatch: token owner is %v, event from user is %v", token.Owner, eventItem.FromUser)
		log.Warnf("Error Is: %v", err)
		return err
	}

	token.Owner = eventItem.ToUser
	if err := database.DB.Save(&token).Error; err != nil {
		log.Warnf("Error updating MysteryBox owner: %v", err)
		return err
	}

	return nil
}
