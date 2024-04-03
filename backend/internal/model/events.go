package model

import "gorm.io/gorm"

type OpenBoxEvent struct {
	gorm.Model
	UserAddress string `gorm:"size:42;index:idx_user_address"`
	BoxID       uint64 `gorm:"index:idx_box_id"`
	BoxType     uint8
	BlockNumber uint64 `gorm:"index:idx_block_number"`
	TxHash      string `gorm:"size:66;index:idx_tx_hash"`
	LogIndex    uint   `gorm:"index:idx_log_index"`
}

type MintBoxEvent struct {
	gorm.Model
	UserAddress string `gorm:"size:42;index:idx_user_address"`
	BoxID       uint64 `gorm:"index:idx_box_id"`
	BlockNumber uint64 `gorm:"index:idx_block_number"`
	TxHash      string `gorm:"size:66;index:idx_tx_hash"`
	LogIndex    uint   `gorm:"index:idx_log_index"`
}
