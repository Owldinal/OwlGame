package model

import (
	"github.com/shopspring/decimal"
	"owl-backend/internal/constant"
	"time"
)

type ResponseCode int

const (
	Success             ResponseCode = 0
	WrongParam          ResponseCode = 401
	NotFound            ResponseCode = 404
	HCaptchaFailed      ResponseCode = 402
	ServerInternalError ResponseCode = 500
)

type ResponseData struct {
	Success bool         `json:"success"`
	Code    ResponseCode `json:"code"`
	Message string       `json:"message"`
	Data    interface{}  `json:"data"`
}

type CursorRequest struct {
	Cursor int `json:"cursor,omitempty" form:"cursor"`
	Limit  int `json:"limit,omitempty" form:"limit"`
}

type CursorResponse[T any] struct {
	Cursor     int  `json:"cursor"`
	Limit      int  `json:"limit"`
	NextCursor int  `json:"next_cursor"`
	HasMore    bool `json:"has_more"`
	List       []T  `json:"list"`
}

type PaginationRequest struct {
	Page    int `json:"page,omitempty" form:"page"`
	PerPage int `json:"per_page,omitempty" form:"per_page"`
}

type PaginationResponse[T any] struct {
	Total     int `json:"total"`
	Page      int `json:"page"`
	PerPage   int `json:"per_page"`
	PageCount int `json:"page_count"`
	List      []T `json:"list"`
}

type GetMintSignatureRequest struct {
	Wallet   string `json:"wallet" binding:"required"`
	HCaptcha string `json:"hcaptcha" binding:"required"`
}

type SignedRequest struct {
	Address   string `json:"address" form:"address" binding:"required"`
	Message   string `json:"message" form:"message" binding:"required"`
	Signature string `json:"signature" form:"signature" binding:"required"`
}

type GetMintSignatureResponse struct {
	Wallet    string `json:"wallet"`
	Hash      string `json:"hash"`
	Signature string `json:"signature"`
}

type GetUserInfoRequest struct {
	Wallet string `json:"wallet" form:"wallet" binding:"required"`
	PaginationRequest
}

type GetUserInfoResponse struct {
	Wallet    string `json:"wallet"`
	HasJoined bool   `json:"has_joined"`

	OwlBalance  int64           `json:"owl_balance"` // load from contract OWL_TOKEN_ADDR
	TotalEarned decimal.Decimal `json:"total_earned"`
	BuffLevel   int             `json:"buff_level"`
	IsMoonBoost bool            `json:"is_moon_boost"`
	ElfInfo     UserBoxInfo     `json:"elf_info"`
	FruitInfo   UserBoxInfo     `json:"fruit_info"`
	OwlInfo     UserBoxInfo     `json:"owl_info"`

	ReferralRewards UserReferralRewards `json:"referral_rewards"`

	InvitationCode string `json:"invitation_code"`
	InviteCount    int    `json:"invite_count"`
}

type GetUserBoxInfoRequest struct {
	Wallet string `json:"wallet" form:"wallet" binding:"required"`
	PaginationRequest
}

type UserReferralRewards struct {
	Total     int64 `json:"total"`
	Claimed   int64 `json:"claimed"`
	Available int64 `json:"available"`
	Locked    int64 `json:"locked"`
}

type UserBoxInfo struct {
	Total          int      `json:"total"`
	Staked         int      `json:"staked"`
	Apr            float64  `json:"apr"`
	Apy            float64  `json:"apy"`
	StakedIdList   []uint64 `json:"staked_id_list"`
	UnstakedIdList []uint64 `json:"unstaked_id_list"`
}

type GetUserOwldinalsRequest struct {
	Wallet string `json:"wallet" form:"wallet" binding:"required"`
	PaginationRequest
}

type UserOwldinal struct {
	TokenId uint64 `json:"token_id"`
	//Type      string `json:"type"`
	TokenUri  string `json:"token_uri"`
	IsStaking bool   `json:"is_staking"`
}

type UserMysteryBox struct {
	TokenId   uint64           `json:"token_id"`
	BoxType   constant.BoxType `json:"box_type"`
	Earning   decimal.Decimal  `json:"earning"`
	Claimed   decimal.Decimal  `json:"claimed"`
	Apr       float64          `json:"apr"`
	Apy       float64          `json:"apy"`
	IsStaking bool             `json:"is_staking"`
}

type GetUserInviterRequest struct {
	Wallet string `json:"wallet" form:"wallet" binding:"required"`
	PaginationRequest
}

type GetUserInviterResponse struct {
	Inviter string `json:"inviter"`
	// maybe incoming here
}

type GetMintTxRequest struct {
	Tx string `json:"tx" form:"tx" binding:"required"`
}

type GetMintTxResponse struct {
	RequestTx string                 `json:"request_tx"`
	MintTx    string                 `json:"mint_tx"`
	JobStatus constant.MintJobStatus `json:"job_status"`
	JobMsg    string                 `json:"job_msg"`
}

type GetGameInfoResponse struct {
	TotalRewards         decimal.Decimal `json:"total_rewards"`
	TotalRewardUSD       decimal.Decimal `json:"total_reward_usd"`
	OwlPrice             decimal.Decimal `json:"owl_price"`
	OwlPriceChange       decimal.Decimal `json:"owl_price_change"`
	TotalMarketCap       decimal.Decimal `json:"total_market_cap"`
	TotalMarketCapChange decimal.Decimal `json:"total_market_cap_change"`
	TotalBurned          decimal.Decimal `json:"total_burned"`
	TotalBurnedChange    decimal.Decimal `json:"total_burned_change"`
	StakedFruitCount     uint64          `json:"staked_fruit_count"`
	StakedElfCount       uint64          `json:"staked_elf_count"`
}

type DataPoint struct {
	Date             time.Time       `json:"date"`
	TotalPoolAmount  decimal.Decimal `json:"total_pool_amount"`
	AllocatedRewards decimal.Decimal `json:"allocated_rewards"`
}

type GetGameRewardsTrendingResponse struct {
	Daily   *RewardTrendingDetail `json:"daily"`
	Weekly  *RewardTrendingDetail `json:"weekly"`
	Monthly *RewardTrendingDetail `json:"monthly"`
}

type RewardsHistoryRequest struct {
	CursorRequest
	Address string `json:"address,omitempty" form:"address"`
}

type RewardTrendingDetail struct {
	From       *time.Time  `json:"from"`
	To         *time.Time  `json:"to"`
	DataPoints []DataPoint `json:"data"`
}

type TreasuryRevenueHistory struct {
	Address         string          `json:"address"`
	Operation       string          `json:"operation"`
	Description     string          `json:"description"`
	Count           int64           `json:"count"`
	Amount          decimal.Decimal `json:"amount"`
	TransactionHash string          `json:"transaction_hash"`
}

type AdminSecret struct {
	Secret string `json:"secret" form:"secret"`
}

type UpdateFruitRequest struct {
	Secret      string     `json:"secret" form:"secret"`
	CompareTime *time.Time `json:"compare_time" form:"compare_time" time_format:"2006-01-02T15:04:05Z" time_utc:"1"`
}

type RequestJobRequest struct {
	Tx string `json:"tx" form:"tx"`
}

type RequestJobResponse struct {
	RequestTx  string     `json:"request_tx"`
	JobTx      string     `json:"job_tx"`
	JobUrl     string     `json:"job_url"`
	IsSuccess  bool       `json:"is_success"`
	LastUpdate *time.Time `json:"last_update"`
}

type ReloadLogRequest struct {
	//FromBlock int64 `json:"from_block"`
	//ToBlock   int64 `json:"to_block"`
	Block int64 `json:"block" form:"block"`
}

type ClaimBoxRequest struct {
	SignedRequest
	TokenIds []uint64 `json:"token_ids" form:"token_ids"`
}

type ClaimBoxResponse struct {
	TotalClaimed    decimal.Decimal   `json:"total_claimed"`
	TotalBurned     decimal.Decimal   `json:"total_burned"`
	TransactionHash string            `json:"transaction_hash"`
	ClaimedBoxes    []MysteryBoxToken `json:"claimed_boxes"`
}
