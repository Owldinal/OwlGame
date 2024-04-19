package service

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"owl-backend/abigen"
	"owl-backend/internal/config"
	"owl-backend/internal/eth"
	"owl-backend/internal/model"
	"owl-backend/pkg/log"
)

var (
	owlGame    *abigen.OwlGame
	privateKey *ecdsa.PrivateKey
)

func init() {
	var err error
	owlGame, err = abigen.NewOwlGame(common.HexToAddress(config.C.OwlGameAddr), eth.Client)
	if err != nil {
		panic(err)
	}

	privateKey, err = crypto.HexToECDSA(config.C.BackendWalletPrivateKey)
}

func UpdateFruitRewards() (response interface{}, code model.ResponseCode, msg string) {
	log.Infof("UpdateFruitRewards: Start")

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(config.C.ChainId))
	if err != nil {
		return false, model.ServerInternalError, fmt.Sprintf("Err: %v", err)
	}

	auth.GasPrice = big.NewInt(config.C.GasPrice)
	tx, err := owlGame.UpdateAllFruitRewards(auth)
	if err != nil {
		log.Warnf("UpdateFruitRewards: Failed: tx=%+v, err=%v", tx, err)
		return false, model.ServerInternalError, fmt.Sprintf("Err: %v", err)
	}

	log.Warnf("UpdateFruitRewards: Transaction sent: %s", tx.Hash().Hex())

	return true, model.Success, ""
}
