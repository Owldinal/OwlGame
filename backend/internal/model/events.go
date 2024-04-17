package model

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/shopspring/decimal"
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
	MintType uint64
	TokenUri string
}

type OwldinalNftTransferEvent struct {
	Event
	database.Model
	FromUser string `gorm:"size:42"`
	ToUser   string `gorm:"size:42"`
	BoxId    uint64
}

type GenOneBoxMintBoxEvent struct {
	Event
	database.Model
	User    string `gorm:"size:42;index:idx_user"`
	TokenId uint64 `gorm:"index:idx_token_id"`
	BoxType uint8
}

type GenOneBoxTransferEvent struct {
	Event
	database.Model
	FromUser string `gorm:"size:42"`
	ToUser   string `gorm:"size:42"`
	TokenId  uint64
}

// eventOwlGameJoinGame
type OwlGameJoinGameEvent struct {
	Event
	database.Model
	User       string `gorm:"size:42"`
	InviteCode uint32
}

// eventOwlGameBindInvitation
type OwlGameBindInvitationEvent struct {
	Event
	database.Model
	Inviter string `gorm:"size:42"`
	Invitee string `gorm:"size:42"`
}

// eventOwlGamePrizePoolIncreased
type OwlGamePrizePoolIncreasedEvent struct {
	Event
	database.Model
	Amount decimal.Decimal `gorm:"type:decimal(36,18);"`
}

// eventOwlGamePrizePoolDecreased
type OwlGamePrizePoolDecreasedEvent struct {
	Event
	database.Model
	Amount decimal.Decimal `gorm:"type:decimal(36,18);"`
}

// eventOwlGameMintMysteryBox
type OwlGameMintMysteryBoxEvent struct {
	Event
	database.Model
	User     string `gorm:"size:42"`
	Count    int64
	TokenIds database.IdList
}

// eventOwlGameStakeOwldinalNft
type OwlGameStakeOwldinalNftEvent struct {
	Event
	database.Model
	User     string `gorm:"size:42"`
	TokenIds database.IdList
}

// eventOwlGameUnstakeOwldinalNft
type OwlGameUnstakeOwldinalNftEvent struct {
	Event
	database.Model
	User     string `gorm:"size:42"`
	TokenIds database.IdList
}

// eventOwlGameStakeMysteryBox
type OwlGameStakeMysteryBoxEvent struct {
	Event
	database.Model
	User      string `gorm:"size:42"`
	TokenIds  database.IdList
	BuffLevel uint8
}

// eventOwlGameUnstakeMysteryBox
type OwlGameUnstakeMysteryBoxEvent struct {
	Event
	database.Model
	User    string `gorm:"size:42"`
	TokenId uint64
	BoxType uint8
	Rewards decimal.Decimal `gorm:"type:decimal(36,18);"`
}

type OwlGameOwlTokenBurnedEvent struct {
	Event
	database.Model
	User      string `gorm:"size:42"`
	MintCount uint64
	Amount    decimal.Decimal `gorm:"type:decimal(36,18);"`
}

// eventOwlGameRebateRewardsIncreased
type OwlGameRebateRewardsIncreasedEvent struct {
	Event
	database.Model
	User      string `gorm:"size:42"`
	Invitee   string `gorm:"size:42"`
	MintCount uint64
	Amount    decimal.Decimal `gorm:"type:decimal(36,18);"`
}

// eventOwlGameUnlockableRebateIncreased
type OwlGameUnlockableRebateIncreasedEvent struct {
	Event
	database.Model
	User      string `gorm:"size:42"`
	MintCount uint64
	Amount    decimal.Decimal `gorm:"type:decimal(36,18);"`
}

// eventOwlGameRebateClaimed
type OwlGameClaimRebateClaimedEvent struct {
	Event
	database.Model
	User   string          `gorm:"size:42"`
	Amount decimal.Decimal `gorm:"type:decimal(36,18);"`
}

type OwlGameFruitRewardUpdateEvent struct {
	Event
	database.Model
	Amount          decimal.Decimal `gorm:"type:decimal(36,18);"`
	Count           uint64
	TotalFruitCount uint64
	TotalElfCount   uint64
}

type OwlGameElfRewardUpdateEvent struct {
	Event
	database.Model
	Amount decimal.Decimal `gorm:"type:decimal(36,18);"`
	Count  uint64
}

type OwlGameRequestMintEvent struct {
	Event
	database.Model
	User      string `gorm:"size:42"`
	RequestId uint64
	Count     uint64
}
