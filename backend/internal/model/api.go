package model

type ResponseCode int

const (
	Success             ResponseCode = 0
	WrongParam          ResponseCode = 401
	ServerInternalError ResponseCode = 500
)

type ResponseData struct {
	Success bool         `json:"success"`
	Code    ResponseCode `json:"code"`
	Message string       `json:"message"`
	Data    interface{}  `json:"data"`
}

type GetMintSignatureRequest struct {
	Wallet string `json:"wallet" binding:"required"`
}

type GetMintSignatureResponse struct {
	Wallet    string `json:"wallet"`
	Hash      string `json:"hash"`
	Signature string `json:"signature"`
}
