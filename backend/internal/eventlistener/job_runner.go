package eventlistener

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"os"
	"os/signal"
	"owl-backend/abigen"
	"owl-backend/internal/config"
	"owl-backend/internal/constant"
	"owl-backend/internal/database"
	"owl-backend/internal/model"
	"owl-backend/pkg/log"
	"syscall"
	"time"
)

var (
	privateKey *ecdsa.PrivateKey
)

func StartJobListening() {
	client, err := ethclient.Dial(config.C.NodeUrl)
	if err != nil {
		log.Fatal("Failed to connect to the Ethereum client: %v", err)
	}

	privateKey, err = crypto.HexToECDSA(config.C.BackendWalletPrivateKey)

	owlGameAddr := common.HexToAddress(config.C.OwlGameAddr)
	owlGame, err := abigen.NewOwlGame(owlGameAddr, client)

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	done := make(chan bool)
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
		<-quit

		done <- true
	}()
	for {
		select {
		case <-ticker.C:
			// 这里是处理作业的逻辑
			processJobs(owlGame)
		case <-done:
			log.Infof("Job listener stopped.")
		}
	}
}

func processJobs(owlGame *abigen.OwlGame) {
	var jobs []model.RequestMintJob
	now := time.Now()
	err := database.DB.
		Where("created_at <= ?", now.Add(-10*time.Second)).
		Where("status = ? OR (status = ? AND retry_count < ?)", constant.MintJobStatusNotStart, constant.MintJobStatusFailed, 3).
		Find(&jobs).Error
	if err != nil {
		log.Warnf("Error retrieving jobs: %v", err)
		return
	}

	for _, job := range jobs {
		job.Status = constant.MintJobStatusProcessing
		database.DB.Save(&job)

		auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(config.C.ChainId))
		if err != nil {
			job.Status = constant.MintJobStatusFailed
			job.RetryCount++
			job.Result = job.Result + fmt.Sprintf("Err: %v;", err)
			database.DB.Save(&job)
			continue
		}

		auth.GasPrice = big.NewInt(50000000)

		tx, err := owlGame.MintMysteryBox(auth, big.NewInt(int64(job.RequestId)))
		if err != nil {
			job.Status = constant.MintJobStatusFailed
			job.RetryCount++
			job.Result = job.Result + fmt.Sprintf("Err: %v;", err)
		} else {
			job.Status = constant.MintJobStatusSuccess
			job.Result = fmt.Sprintf("Success: %v", tx.Hash().String())
		}

		database.DB.Save(&job)
	}
}
