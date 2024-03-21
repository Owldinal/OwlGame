package service

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"owl-backend/abigen"
	"owl-backend/internal/config"
	"owl-backend/internal/eth"
	"owl-backend/internal/model"
)

var (
	boxGen0ContractAddress common.Address
	boxGen0Contract        *abigen.MysteryBoxGen0
	signatureArguments     abi.Arguments
	backendPrivateKey      *ecdsa.PrivateKey
)

func init() {
	boxGen0ContractAddress = common.HexToAddress(config.C.NftOwlAddr)
	var err error
	boxGen0Contract, err = abigen.NewMysteryBoxGen0(boxGen0ContractAddress, eth.Client)
	if err != nil {
		panic(err)
	}

	parameterTypeUint256, _ := abi.NewType("uint256", "", nil)
	parameterTypeAddress, _ := abi.NewType("address", "", nil)
	signatureArguments = abi.Arguments{
		{
			Type: parameterTypeAddress,
		},
		{
			Type: parameterTypeUint256,
		},
		{
			Type: parameterTypeAddress,
		},
	}

	backendPrivateKey, err = crypto.HexToECDSA(config.C.BackendWalletPrivateKey)
	if err != nil {
		panic(err)
	}
}

func GenerateHashAndSignForMint(wallet string) (response interface{}, code model.ResponseCode, msg string) {
	tokenId, err := getTokenId()
	if err != nil {
		return nil, model.ServerInternalError, err.Error()
	}

	hash, signature, err := generateHashAndSignature(wallet, tokenId)
	if err != nil {
		return nil, model.ServerInternalError, err.Error()
	}

	response = &model.GetMintSignatureResponse{
		Wallet:    wallet,
		Hash:      hash,
		Signature: signature,
	}

	return response, model.Success, ""
}

func getTokenId() (*big.Int, error) {
	tokenIdCounter, err := boxGen0Contract.TokenIdCounter(nil)
	if err != nil {
		return nil, err
	}

	return tokenIdCounter, nil
}

func generateHashAndSignature(wallet string, tokenId *big.Int) (string, string, error) {
	bytes, err := signatureArguments.Pack(common.HexToAddress(wallet), tokenId, boxGen0ContractAddress)
	if err != nil {
		return "", "", err
	}

	hash := crypto.Keccak256Hash(bytes)

	signatureBytes, _ := crypto.Sign(accounts.TextHash(hash.Bytes()), backendPrivateKey)
	signature := hexutil.Encode(signatureBytes)

	return hash.Hex(), signature, nil

}
