package model

import (
	"owl-backend/internal/database"
)

type OwldinalNftToken struct {
	database.Model
	TokenId    uint64          `json:"token_id" gorm:"index:idx_token_id"`
	Owner      string          `json:"owner" gorm:"size:42;index:idx_owner"`
	TokenUri   string          `json:"token_uri"`
	IsStaking  bool            `json:"is_staking"`
	BuffingIds database.IdList `json:"buffing_ids"`
}
