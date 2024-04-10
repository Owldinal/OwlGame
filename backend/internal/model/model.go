package model

import (
	"github.com/shopspring/decimal"
	"owl-backend/internal/constant"
	"owl-backend/internal/database"
	"time"
)

type OwldinalNftToken struct {
	database.Model
	TokenId    uint64          `json:"token_id" gorm:"index:idx_token_id"`
	Owner      string          `json:"owner" gorm:"size:42;index:idx_owner"`
	TokenUri   string          `json:"token_uri"`
	IsStaking  bool            `json:"is_staking"`
	BuffingIds database.IdList `json:"buffing_ids"`
}

type MysteryBoxToken struct {
	database.Model
	TokenId        uint64           `json:"token_id" gorm:"index:idx_token_id"`
	Owner          string           `json:"owner" gorm:"size:42;index:idx_owner"`
	BoxType        constant.BoxType `json:"box_type" gorm:"type:TINYINT UNSIGNED NOT NULL"`
	IsStaking      bool             `json:"is_staking"`
	BuffLevel      uint8            `json:"buff_level"`
	CurrentRewards decimal.Decimal  `json:"current_rewards" gorm:"type:decimal(36,18);"`
	TotalRewards   decimal.Decimal  `json:"total_rewards" gorm:"type:decimal(36,18);"`
}

type InviteRelation struct {
	database.Model
	Inviter         string          `gorm:"size:42;index:idx_Inviter"`
	Invitee         string          `gorm:"size:42;"`
	ReferralRewards decimal.Decimal `json:"referral_rewards" gorm:"type:decimal(36,18);"`
}

type UserInfo struct {
	database.Model
	Address     string
	InviteCode  string `json:"invite_code"`
	BuffLevel   int    `json:"buff_level"`
	InviteCount int    `json:"invite_count"`
	//TotalEarned decimal.Decimal `json:"total_earned" gorm:"type:decimal(36,18)"`
	UnclaimedReferral  decimal.Decimal `json:"-" gorm:"type:decimal(36,18)"`
	UnlockableReferral decimal.Decimal `json:"-" gorm:"type:decimal(36,18)"`
	ClaimedReferral    decimal.Decimal `json:"claimed_referral" gorm:"type:decimal(36,18)"`

	AvailableReferral decimal.Decimal `json:"available_referral" gorm:"-"`
	LockedReferral    decimal.Decimal `json:"locked_referral" gorm:"-"`
}

type DailyPoolSnapshot struct {
	database.Model
	Date time.Time

	TotalPoolAmount  decimal.Decimal `json:"total_pool_amount" gorm:"type:decimal(36,18)"`
	AllocatedRewards decimal.Decimal `json:"allocated_rewards" gorm:"type:decimal(36,18)"`
	TotalMarketCap   decimal.Decimal `json:"total_market_cap" gorm:"type:decimal(36,18)"`
	TotalBurn        decimal.Decimal `json:"total_burn" gorm:"type:decimal(36,18)"`
}

type RewardPoolTransactionRecord struct {
	database.Model
	User        string
	Operation   string
	Description string
	Count       int
	Amount      decimal.Decimal `gorm:"type:decimal(36,18)"`
	Event
}

// region ---- Enums ----

// endregion
