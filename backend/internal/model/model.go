package model

import (
	"github.com/shopspring/decimal"
	"owl-backend/internal/constant"
	"owl-backend/internal/database"
	"time"
)

type OwldinalNftToken struct {
	database.Model
	TokenId     uint64          `json:"token_id" gorm:"index:idx_token_id"`
	Owner       string          `json:"owner" gorm:"size:42;index:idx_owner"`
	TokenUri    string          `json:"token_uri"`
	IsStaking   bool            `json:"is_staking"`
	StakingTime *time.Time      `json:"staking_time"`
	BuffingIds  database.IdList `json:"buffing_ids"` // not in use
}

type MysteryBoxToken struct {
	database.Model
	TokenId        uint64           `json:"token_id" gorm:"index:idx_token_id"`
	Owner          string           `json:"owner" gorm:"size:42;index:idx_owner"`
	BoxType        constant.BoxType `json:"box_type" gorm:"type:TINYINT UNSIGNED NOT NULL"`
	IsStaking      bool             `json:"is_staking"`
	StakingTime    *time.Time       `json:"staking_time"`
	BuffLevel      uint8            `json:"buff_level"`
	CurrentRewards decimal.Decimal  `json:"current_rewards" gorm:"type:decimal(36,18);"`
	TotalRewards   decimal.Decimal  `json:"total_rewards" gorm:"type:decimal(36,18);"`
	ClaimedRewards decimal.Decimal  `json:"claimed_rewards" gorm:"type:decimal(36,18);"`
	// ClaimedRewards + CurrentRewards = TotalRewards
}

type InviteRelation struct {
	database.Model
	Event
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

	// total earned from mystery box claim
	TotalEarned decimal.Decimal `json:"total_earned" gorm:"type:decimal(36,18)"`

	// invite referral
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

type AprSnapshot struct {
	database.Model
	Event
	RewardCount     uint64
	TotalFruitCount uint64
	TotalElfCount   uint64
	FruitRewards    decimal.Decimal `gorm:"type:decimal(36,18)"`
	FruitApr        float64
	FruitApy        float64
	ElfApr          float64
	ElfApy          float64
}

type RewardPoolTransactionRecord struct {
	database.Model
	User        string
	Operation   string
	Description string
	Count       uint64
	Amount      decimal.Decimal `gorm:"type:decimal(36,18)"`
	Event
}

type RequestMintJob struct {
	database.Model
	User      string
	RequestId uint64
	Count     uint64
	// Status: 0 means waiting to do ; 1 means processing ; 2 means success; 3 means failed
	Status       constant.MintJobStatus
	HasConfirmed bool

	// Result: if success, there will be a json string like {elf:[1,2,3], fruit:[5,6,7], burn:[8,9]}
	// Failed: if failed, there will be a error msg
	Result     string
	RetryCount uint8

	RequestTxHash      string `gorm:"size:66;"`
	RequestLogIndex    uint
	RequestBlockNumber uint64
	RequestBlockHash   string `gorm:"size:66;"`

	JobTxHash      string `gorm:"size:66;"`
	JobLogIndex    uint
	JobBlockNumber uint64
	JobBlockHash   string `gorm:"size:66"`
}

type TransferRewards struct {
	database.Model
	User          string
	TokenId       uint64
	BoxType       constant.BoxType `json:"box_type" gorm:"type:TINYINT UNSIGNED NOT NULL"`
	Rewards       decimal.Decimal  `gorm:"type:decimal(36,18)"`
	BurnedRewards decimal.Decimal  `gorm:"type:decimal(36,18)"` // for fruit，give it to elf; for elf, burn it.
	BuffLevel     uint8
	MoonBoost     bool
	IsConfirmed   bool
	Status        uint8
	Result        string

	ClaimTxHash      string `gorm:"size:66;"`
	ClaimLogIndex    uint
	ClaimBlockNumber uint64
	ClaimBlockHash   string `gorm:"size:66;"`

	TransferTxHash      string `gorm:"size:66;"`
	TransferLogIndex    uint
	TransferBlockNumber uint64
	TransferBlockHash   string `gorm:"size:66"`

	BurnTxHash string `gorm:"size:66;"`
}
