package constant

import "github.com/shopspring/decimal"

type BoxType uint8

const (
	BoxTypeNotExist BoxType = iota
	BoxTypeElf
	BoxTypeFruit
	BoxTypeBurned
)

type MintJobStatus uint8

const (
	MintJobStatusNotStart MintJobStatus = iota
	MintJobStatusProcessing
	MintJobStatusSuccess
	MintJobStatusFailed
)

const NoneAddr = "0x0000000000000000000000000000000000000000"

var MysteryBoxMintPrice = decimal.NewFromInt(100000)
