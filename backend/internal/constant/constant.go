package constant

import "github.com/shopspring/decimal"

type BoxType uint8

const (
	BoxTypeNotExist BoxType = iota
	BoxTypeElf
	BoxTypeFruit
	BoxTypeBurned
)

const NoneAddr = "0x0000000000000000000000000000000000000000"

var MysteryBoxMintPrice = decimal.NewFromInt(100000)
