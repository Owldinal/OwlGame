package model

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
	Cursor int `json:"cursor" form:"cursor"`
	Limit  int `json:"limit" form:"limit"`
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
	Wallet      string      `json:"wallet"`
	OwlBalance  int64       `json:"owl_balance"` // load from contract OWL_TOKEN_ADDR
	TotalEarned int64       `json:"total_earned"`
	BuffLevel   int         `json:"buff_level"`
	ElfInfo     UserBoxInfo `json:"elf_info"`
	FruitInfo   UserBoxInfo `json:"fruit_info"`
	OwlInfo     UserBoxInfo `json:"owl_info"`

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
	StakedIdList   []uint64 `json:"staked_id_list"`
	UnstakedIdList []uint64 `json:"unstaked_id_list"`
}

type GetUserOwldinalsRequest struct {
	Wallet string `json:"wallet" form:"wallet" binding:"required"`
	PaginationRequest
}

type GetUserOwldinalsResponse struct {
	TokenId  uint64 `json:"token_id"`
	Type     string `json:"type"`
	TokenUrl string `json:"token_url"`
	Earning  int64  `json:"earning"`
	Apr      int64  `json:"apr"`
	Status   string `json:"status"`
}

type GetUserInviterRequest struct {
	Wallet string `json:"wallet" form:"wallet" binding:"required"`
	PaginationRequest
}

type GetUserInviterResponse struct {
	Inviter string `json:"inviter"`
	// maybe incoming here
}

type GetGameInfoResponse struct {
	TotalRewards         int64 `json:"total_rewards"`
	TotalRewardUSD       int64 `json:"total_reward_usd"`
	OwlPrice             int64 `json:"owl_price"`
	OwlPriceChange       int64 `json:"owl_price_change"`
	TotalMarketCap       int64 `json:"total_market_cap"`
	TotalMarketCapChange int64 `json:"total_market_cap_change"`
	TotalBurned          int64 `json:"total_burned"`
	TotalBurnedChange    int64 `json:"total_burned_change"`
}

type DataPoint struct {
	Date             string `json:"date"`
	TotalPoolAmount  int64  `json:"total_pool_amount"`
	AllocatedRewards int64  `json:"allocated_rewards"`
}

type GetGameRewardsTrendingResponse struct {
	Daily struct {
		From       string      `json:"from"`
		To         string      `json:"to"`
		DataPoints []DataPoint `json:"data"`
	} `json:"daily"`
	Weekly struct {
		From       string      `json:"from"`
		To         string      `json:"to"`
		DataPoints []DataPoint `json:"data"`
	}
	Monthly struct {
		From       string      `json:"from"`
		To         string      `json:"to"`
		DataPoints []DataPoint `json:"data"`
	}
}

type TreasuryRevenueHistory struct {
	Address         string `json:"address"`
	Operation       string `json:"operation"`
	Description     string `json:"description"`
	Count           int64  `json:"count"`
	Amount          int64  `json:"amount"`
	TransactionHash string `json:"transaction_hash"`
}
