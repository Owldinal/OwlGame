package service

import (
	"gorm.io/gorm/clause"
	"owl-backend/abigen"
	"owl-backend/internal/database"
	"owl-backend/internal/model"
)

func SaveOpenBoxEvent(event *abigen.OwldinalGenOneBoxOpenBox) error {
	item := model.OpenBoxEvent{
		UserAddress: event.User.Hex(),
		BoxID:       event.TokenId.Uint64(),
		BoxType:     event.BoxType,
		BlockNumber: event.Raw.BlockNumber,
		TxHash:      event.Raw.TxHash.Hex(),
		LogIndex:    event.Raw.Index,
	}
	result := database.DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&item)
	return result.Error
}
