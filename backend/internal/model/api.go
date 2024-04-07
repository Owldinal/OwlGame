package model

type ResponseCode int

const (
	Success             ResponseCode = 0
	WrongParam          ResponseCode = 401
	HCaptchaFailed      ResponseCode = 402
	ServerInternalError ResponseCode = 500
)

type ResponseData struct {
	Success bool         `json:"success"`
	Code    ResponseCode `json:"code"`
	Message string       `json:"message"`
	Data    interface{}  `json:"data"`
}

type PaginationRequest struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
}

type PaginationResponse[T any] struct {
	Total     int `json:"total"`
	Page      int `json:"page"`
	PerPage   int `json:"per_page"`
	PageCount int `json:"page_count"`
	List      []T `json:"list"`
}

type CursorPaginationRequest struct {
	Cursor int `json:"cursor"`
	Limit  int `json:"limit"`
}

type CursorPaginationResponse[T any] struct {
	Cursor     int  `json:"cursor"`
	Limit      int  `json:"limit"`
	NextCursor int  `json:"next_cursor"`
	HasMore    bool `json:"has_more"`
	List       []T  `json:"list"`
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
	Wallet     string            `json:"wallet" binding:"required"`
	Pagination PaginationRequest `json:"pagination,omitempty"`
}

type GetUserInfoResponse struct {
	Wallet      string      `json:"wallet"`
	OwlBalance  int64       `json:"owl_balance"` // load from contract OWL_TOKEN_ADDR
	TotalEarned int64       `json:"total_earned"`
	BuffLevel   int         `json:"buff_level"`
	ElfInfo     UserBoxInfo `json:"elf_info"`
	FruitInfo   UserBoxInfo `json:"fruit_info"`

	ReferralRewards struct {
		Total     int64 `json:"total"`
		Claimed   int64 `json:"claimed"`
		Available int64 `json:"available"`
		Locked    int64 `json:"locked"`
	} `json:"referral_rewards"`

	InvitationCode string `json:"invitation_code"`
	InviteCount    int    `json:"invite_count"`
}

type UserBoxInfo struct {
	Total          int64   `json:"total"`
	Staked         int64   `json:"staked"`
	Apr            float64 `json:"apr"`
	StakedIdList   []int64 `json:"staked_id_list"`
	UnstakedIdList []int64 `json:"unstaked_id_list"`
}
