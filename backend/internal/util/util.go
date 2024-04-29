package util

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"strconv"
)

func DecodeInviteCode(encoded uint32) string {
	var inviteCode string
	for i := 0; i < 5; i++ {
		charValue := (encoded >> (i * 5)) & 0x1F
		inviteCode += string(rune(charValue + 65))
	}
	return inviteCode
}

func ValidSignature(address, message, signature string) error {
	hashedMessage := []byte("\x19Ethereum Signed Message:\n" + strconv.Itoa(len(message)) + message)
	hash := crypto.Keccak256Hash(hashedMessage)
	//fmt.Printf("hash: %x\n", hash)

	decodedMessage := hexutil.MustDecode(signature)
	if decodedMessage[64] == 27 || decodedMessage[64] == 28 {
		decodedMessage[64] -= 27
	}

	sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), decodedMessage)
	if sigPublicKeyECDSA == nil {
		err = errors.New("could not get a public get from the message signature")
	}
	if err != nil {
		return err
	}

	publicKey := crypto.PubkeyToAddress(*sigPublicKeyECDSA).String()
	if publicKey != address {
		return errors.New(fmt.Sprintf("signature not same with expect. expect: %v, actual: %v", address, publicKey))
	}

	return nil
}
