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
	"owl-backend/internal/database"
	"owl-backend/internal/eth"
	"owl-backend/internal/eventlistener"
	"owl-backend/internal/model"
	"owl-backend/pkg/log"
	"time"
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

	if config.C.GasPrice > 0 {
		auth.GasPrice = big.NewInt(config.C.GasPrice)
	}
	tx, err := owlGame.UpdateAllFruitRewards(auth)
	if err != nil {
		log.Warnf("UpdateFruitRewards: Failed: tx=%+v, err=%v", tx, err)
		return false, model.ServerInternalError, fmt.Sprintf("Err: %v", err)
	}

	log.Warnf("UpdateFruitRewards: Transaction sent: %s", tx.Hash().Hex())

	return true, model.Success, ""
}

func RetryAllJobs() (response interface{}, code model.ResponseCode, msg string) {
	log.Infof("Retry all job")
	jobIds := processJobs()
	return jobIds, model.Success, ""
}

func processJobs() (jobIds []uint) {
	jobIds = make([]uint, 0)
	var jobs []model.RequestMintJob
	now := time.Now()
	err := database.DB.
		Where("created_at <= ?", now.Add(-2*time.Minute)).
		Where("has_confirmed = ? ", 0).
		Find(&jobs).Error
	if err != nil {
		log.Warnf("Error retrieving jobs: %v", err)
		return
	}

	if len(jobs) == 0 {
		return
	}

	for _, job := range jobs {
		jobIds = append(jobIds, job.ID)
	}
	//
	//err = database.DB.Model(&model.RequestMintJob{}).
	//	Where("id IN ?", jobIds).
	//	Omit("HasConfirmed").
	//	Updates(map[string]interface{}{
	//		"RetryCount": 0,
	//		"Status":     constant.MintJobStatusFailed,
	//	}).Error
	//if err != nil {
	//	log.Warnf("Error updating jobs: %v", err)
	//}

	return
}

func RetryTransferMultiple(taskId int64) (response interface{}, code model.ResponseCode, msg string) {
	log.Infof("Retry Transfer Multiple for %v", taskId)
	jobIds := eventlistener.RetryClaimMultipleTask(taskId)
	return jobIds, model.Success, ""
}
