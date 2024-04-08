package model

import (
	"github.com/ethereum/go-ethereum/core/types"
	"owl-backend/internal/database"
)

type Event struct {
	TxHash      string `gorm:"size:66;index:idx_tx_hash"`
	LogIndex    uint   `gorm:"index:idx_log_index"`
	BlockNumber uint64 `gorm:"index:idx_block_number"`
	BlockHash   string `gorm:"size:66;block_hash"`
}

func NewEvent(l *types.Log) Event {
	return Event{
		TxHash:      l.TxHash.Hex(),
		LogIndex:    l.Index,
		BlockNumber: l.BlockNumber,
		BlockHash:   l.BlockHash.Hex(),
	}
}

type OwldinalNftMintBoxEvent struct {
	Event
	database.Model
	User     string `gorm:"size:42;index:idx_user_address"`
	BoxId    uint64 `gorm:"index:idx_box_id"`
	MintType uint64 `gorm:"mint_type"`
	TokenUri string `gorm:"token_uri"`
}

type OwldinalNftTransferEvent struct {
	Event
	database.Model
	FromUser string `gorm:"size:42"`
	ToUser   string `gorm:"size:42"`
	BoxId    uint64
}
