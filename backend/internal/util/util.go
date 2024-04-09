package util

import "strconv"

func DecodeInviteCode(encoded uint32) string {
	var inviteCode string
	for i := 0; i < 5; i++ {
		charValue := (encoded >> (uint32(i) * 5)) & 0x1F
		inviteCode += strconv.FormatUint(uint64(charValue)+0x41, 10)
	}
	return inviteCode
}
