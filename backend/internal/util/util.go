package util

func DecodeInviteCode(encoded uint32) string {
	var inviteCode string
	for i := 0; i < 5; i++ {
		charValue := (encoded >> (i * 5)) & 0x1F
		inviteCode += string(rune(charValue + 65))
	}
	return inviteCode
}
