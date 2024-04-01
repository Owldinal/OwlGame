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
	boxGen0Contract        *abigen.Owldinal
	signatureArguments     abi.Arguments
	backendPrivateKey      *ecdsa.PrivateKey
)

func init() {
	boxGen0ContractAddress = common.HexToAddress(config.C.NftOwlAddr)
	var err error
	boxGen0Contract, err = abigen.NewOwldinal(boxGen0ContractAddress, eth.Client)
	if err != nil {
		panic(err)
	}

	parameterTypeAddress, _ := abi.NewType("address", "", nil)
	signatureArguments = abi.Arguments{
		{
			Type: parameterTypeAddress,
		},
		{
			Type: parameterTypeAddress,
		},
	}

	backendPrivateKey, err = crypto.HexToECDSA(config.C.BackendWalletPrivateKey)

	//ecdsaPublicKey, ok := backendPrivateKey.Public().(*ecdsa.PublicKey)
	//if !ok {
	//	log.Fatal("error casting public key to ECDSA")
	//}
	//// 打印公钥的十六进制表示
	//publicKeyBytes := elliptic.Marshal(ecdsaPublicKey.Curve, ecdsaPublicKey.X, ecdsaPublicKey.Y)
	//fmt.Printf("Public key (hex): %x\n", publicKeyBytes)
	//address := crypto.PubkeyToAddress(*ecdsaPublicKey).Hex()
	//fmt.Printf("Public key to address: %s\n", address)
	//if err != nil {
	//	panic(err)
	//}
}

func GenerateHashAndSignForMint(wallet string) (response interface{}, code model.ResponseCode, msg string) {
	//tokenId, err := getTokenId()
	//if err != nil {
	//	return nil, model.ServerInternalError, err.Error()
	//}

	hash, signature, err := generateHashAndSignature(wallet, nil)
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
	bytes, err := signatureArguments.Pack(common.HexToAddress(wallet), boxGen0ContractAddress)
	if err != nil {
		return "", "", err
	}

	hash := crypto.Keccak256Hash(bytes)
	signatureBytes, _ := crypto.Sign(accounts.TextHash(hash.Bytes()), backendPrivateKey)

	// The ethereum need v=27 || 28, but golang gives 0 || 1
	if signatureBytes[64] == 0 || signatureBytes[64] == 1 {
		signatureBytes[64] += 27
	}
	signature := hexutil.Encode(signatureBytes)

	return hash.Hex(), signature, nil

}
