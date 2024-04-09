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

type OwldinalNftMintBoxHandler struct{}

func (h *OwldinalNftMintBoxHandler) Handle(vlog types.Log) error {
	event, err := owldinalNftContract.ParseMintBox(vlog)
	if err != nil {
		return err
	}
	// save event to database
	//log.Infof("[%v-%v] Mint box: user = %v, boxId = %v", event.Raw.TxHash, event.Raw.Index, event.User, event.BoxId.Uint64())
	eventItem := model.OwldinalNftMintBoxEvent{
		Event:    model.NewEvent(&event.Raw),
		User:     event.User.Hex(),
		BoxId:    event.BoxId.Uint64(),
		MintType: event.MintType.Uint64(),
		TokenUri: event.Url,
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

	// update nft token
	item := model.OwldinalNftToken{
		TokenId:  eventItem.BoxId,
		Owner:    eventItem.User,
		TokenUri: eventItem.TokenUri,
	}
	itemResult := database.DB.Clauses().Create(&item)
	if itemResult.Error != nil {
		return itemResult.Error
	}

	return nil
}

type OwldinalNftTransferHandler struct{}

func (h *OwldinalNftTransferHandler) Handle(vlog types.Log) error {
	event, err := owldinalNftContract.ParseTransfer(vlog)
	if err != nil {
		return err
	}
	//log.Infof("[%v-%v] Transfer from = %v , to = %v , box= %v", event.Raw.TxHash, event.Raw.Index, event.From, event.To, event.TokenId)

	// save event to database
	eventItem := model.OwldinalNftTransferEvent{
		Event:    model.NewEvent(&event.Raw),
		FromUser: event.From.Hex(),
		ToUser:   event.To.Hex(),
		BoxId:    event.TokenId.Uint64(),
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
		//
		return nil
	}

	// is staking or unstaking, handle it in owl_game_handler
	if eventItem.FromUser == OwlGameAddr || eventItem.ToUser == OwlGameAddr {
		return nil
	}

	// Check previous record is correct
	var token model.OwldinalNftToken
	// this token should exist
	if err := database.DB.Where("token_id = ?", eventItem.BoxId).First(&token).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Warnf("Token not found with id: %v", eventItem.BoxId)
			return err
		} else {
			log.Warnf("Error Is: %v", err)
			return err
		}
	}

	// wrong owner
	if token.Owner != eventItem.FromUser {
		err := fmt.Errorf("token owner mismatch: token owner is %v, event from user is %v", token.Owner, eventItem.FromUser)
		log.Warnf("Error Is: %v", err)
		return err
	}

	token.Owner = eventItem.ToUser
	if err := database.DB.Save(&token).Error; err != nil {
		log.Warnf("Error updating token owner: %v", err)
		return err
	}

	return nil
}
