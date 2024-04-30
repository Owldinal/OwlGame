package eventlistener

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"math/big"
	"os"
	"os/signal"
	"owl-backend/abigen"
	"owl-backend/internal/config"
	"owl-backend/internal/constant"
	"owl-backend/internal/database"
	"owl-backend/internal/model"
	"owl-backend/pkg/log"
	"strings"
	"syscall"
	"time"
)

var (
	privateKey *ecdsa.PrivateKey
	ethClient  *ethclient.Client
	owlGame    *abigen.OwlGame
)

func init() {
	var err error
	ethClient, err = ethclient.Dial(config.C.NodeUrl)
	if err != nil {
		log.Fatal("Failed to connect to the Ethereum client: %v", err)
	}

	privateKey, err = crypto.HexToECDSA(config.C.BackendWalletPrivateKey)
}

func StartJobListening() {
	owlGameAddr = common.HexToAddress(config.C.OwlGameAddr)
	owlGame, _ = abigen.NewOwlGame(owlGameAddr, ethClient)

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	// Timer used to check for unconfirmed mint jobs.
	checkUnconfirmJobTicker := time.NewTicker(10 * time.Minute)
	defer checkUnconfirmJobTicker.Stop()

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
			processJobs(owlGame)
		case <-checkUnconfirmJobTicker.C:
			CheckUnconfirmedJob()
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
		database.DB.Model(&job).Omit("HasConfirmed").Updates(job)

		auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(config.C.ChainId))
		if err != nil {
			job.Status = constant.MintJobStatusFailed
			job.RetryCount++
			job.Result = job.Result + fmt.Sprintf("Err: %v;", err)
			database.DB.Model(&job).Omit("HasConfirmed").Updates(job)
			continue
		}

		// Print Gas Limit
		if config.C.GasPrice > 0 {
			auth.GasPrice = big.NewInt(config.C.GasPrice)
		}
		//auth.GasLimit = 100000

		tx, err := owlGame.MintMysteryBox(auth, big.NewInt(int64(job.RequestId)))
		if err != nil {
			job.Status = constant.MintJobStatusFailed
			job.RetryCount++
			job.Result = job.Result + fmt.Sprintf("Err: %v;", err)
		} else {
			job.Status = constant.MintJobStatusSuccess
			job.Result = job.Result + fmt.Sprintf("Success;")
			job.JobTxHash = tx.Hash().Hex()

			// receipt:
			if true {
				receipt, err := bind.WaitMined(context.Background(), ethClient, tx)
				if err != nil {
					// do nothing
					//job.Status = constant.MintJobStatusFailed
					job.Result = job.Result + fmt.Sprintf("WaitMined Err: %v;", err)
				} else {
					job.Status = constant.MintJobStatusSuccess
					job.Result = job.Result + fmt.Sprintf("Success Mint (Gas = %v)", receipt.GasUsed)

					job.JobBlockHash = receipt.BlockHash.Hex()
					job.JobBlockNumber = receipt.BlockNumber.Uint64()
					for _, eventLog := range receipt.Logs {
						if eventLog.Topics[0].Hex() == "0xdf3efac43e1fbda7cd95675458501cb76f194c5347392348f2275a375a5f9863" {
							job.JobLogIndex = eventLog.Index
							break
						}
					}
				}
			}
		}

		database.DB.Model(&job).Omit("HasConfirmed").Updates(job)
		time.Sleep(time.Millisecond * 500)
	}
}

func CheckUnconfirmedJob() {
	var jobs []model.RequestMintJob
	now := time.Now()
	err := database.DB.
		Where("created_at <= ?", now.Add(-5*time.Minute)).
		Where("has_confirmed = ?", false).
		Where("status = ? OR (status = ? AND retry_count >= ?)", constant.MintJobStatusSuccess, constant.MintJobStatusFailed, 3).
		Find(&jobs).Error
	if err != nil {
		log.Warnf("Error retrieving jobs: %v", err)
		return
	}

	requestIdList := make([]uint64, 0)
	for _, job := range jobs {
		// This situation indicates that the contract has already processed this request and can be skipped directly.
		if strings.Contains(job.Result, "no request") {
			job.Status = constant.MintJobStatusSuccess
			job.HasConfirmed = true
			job.Result = job.Result + fmt.Sprintf("Confirmed by checker;")
			database.DB.Model(&job).Updates(job)
			continue
		}

		if job.Status == constant.MintJobStatusSuccess {
			job.Result = job.Result + fmt.Sprintf("Prev Tx: %v;", job.JobTxHash)
		}

		job.Status = constant.MintJobStatusNotStart
		database.DB.Model(&job).Updates(map[string]interface{}{
			"Status": constant.MintJobStatusNotStart,
			"Result": job.Result,
		})

		requestIdList = append(requestIdList, job.RequestId)
	}

	log.Infof("JobRetry: size=%v, id=%v", len(requestIdList), requestIdList)
}

func UpdateFruitReward(compareTime *time.Time) {
	var fruits []model.MysteryBoxToken
	if compareTime == nil {
		// 应该是 -4*time.Hour，保险点留个余量
		t := time.Now().Add(-230 * time.Minute).UTC()
		compareTime = &t
	}
	query := database.DB.
		Where("is_staking = ?", 1).
		Where("box_type = ?", constant.BoxTypeFruit)
	query = query.Where("staking_time <= ?", compareTime)
	err := query.Find(&fruits).Error
	if err != nil {
		log.Warnf("Error retrieiving fruit: %v", err)
		return
	}

	if len(fruits) == 0 {
		log.Warnf("No fruit need to be updated")
		return
	}

	var snapshot model.DailyPoolSnapshot
	if err := database.DB.Order("id desc").First(&snapshot).Error; err != nil {
		log.Warnf("Error retrieving the latest DailyPoolSnapshot: %v", err)
		return
	}

	fruitRewardsProportion := calculateFruitRewardsProportion(len(fruits))
	totalRewards := snapshot.TotalPoolAmount.Mul(fruitRewardsProportion).Div(decimal.NewFromInt(100000000))
	eachFruitRewards := totalRewards.Div(decimal.NewFromInt(int64(len(fruits))))

	log.Infof("Update fruit rewards: time=%v, pool=%v, fruits=%v, proportion=%v, rewards=%v, eachRewards=%v",
		compareTime,
		snapshot.TotalPoolAmount,
		len(fruits),
		fruitRewardsProportion,
		totalRewards,
		eachFruitRewards,
	)

	//update current_rewards
	for _, fruit := range fruits {
		fruit.CurrentRewards = fruit.CurrentRewards.Add(eachFruitRewards)
		fruit.TotalRewards = fruit.TotalRewards.Add(eachFruitRewards)

		dbSaveErr := database.DB.Save(&fruit).Error
		if dbSaveErr != nil {
			log.Warnf("Error saving updated fruit reward to the database: %v", dbSaveErr)
		}
	}

	err = UpdateDailyPoolSnapshot(DailyPoolUpdater{Decrease: totalRewards})
	if err != nil {
		log.Warnf("Failed to update daily pool snapshot, %v", err)
		return
	}

	// update apr snapshot
	var elfCount, fruitCount int64
	err = database.DB.Model(&model.MysteryBoxToken{}).
		Where("box_type = ?", constant.BoxTypeElf).
		Where("is_staking = ?", true).
		Count(&elfCount).Error
	if err != nil {
		log.Warnf("updateFruitRewards: Failed to load elfCount, %v", err)
	}
	err = database.DB.Model(&model.MysteryBoxToken{}).
		Where("box_type = ?", constant.BoxTypeFruit).
		Where("is_staking = ?", true).
		Count(&fruitCount).Error
	if err != nil {
		log.Warnf("updateFruitRewards: Failed to load fruitCount, %v", err)
	}
	UpdateAprSnapshot(&model.OwlGameFruitRewardUpdateEvent{
		Count:           uint64(len(fruits)),
		TotalFruitCount: uint64(fruitCount),
		TotalElfCount:   uint64(elfCount),
		Amount:          eachFruitRewards,
	})
}

func calculateFruitRewardsProportion(rewardFruitCount int) decimal.Decimal {
	var fruitRewardsProportion decimal.Decimal
	switch {
	case rewardFruitCount <= 1:
		fruitRewardsProportion = decimal.NewFromInt(100) // 0.0001000% * 100_000_000
	case rewardFruitCount <= 50:
		fruitRewardsProportion = decimal.NewFromInt(5000) // 0.0050000% * 100_000_000
	case rewardFruitCount <= 100:
		fruitRewardsProportion = decimal.NewFromInt(10000) // 0.0100000% * 100_000_000
	case rewardFruitCount <= 200:
		fruitRewardsProportion = decimal.NewFromInt(21000) // 0.0210000% * 100_000_000
	case rewardFruitCount <= 400:
		fruitRewardsProportion = decimal.NewFromInt(43040) // 0.0430400% * 100_000_000
	case rewardFruitCount <= 800:
		fruitRewardsProportion = decimal.NewFromInt(88230) // 0.0882300% * 100_000_000
	case rewardFruitCount <= 1200:
		fruitRewardsProportion = decimal.NewFromInt(136760) // 0.1367600% * 100_000_000
	case rewardFruitCount <= 1800:
		fruitRewardsProportion = decimal.NewFromInt(211980) // 0.2119800% * 100_000_000
	case rewardFruitCount <= 2300:
		fruitRewardsProportion = decimal.NewFromInt(276750) // 0.2767500% * 100_000_000
	case rewardFruitCount <= 2800:
		fruitRewardsProportion = decimal.NewFromInt(342930) // 0.3429300% * 100_000_000
	case rewardFruitCount <= 3300:
		fruitRewardsProportion = decimal.NewFromInt(416410) // 0.4164100% * 100_000_000
	case rewardFruitCount <= 4000:
		fruitRewardsProportion = decimal.NewFromInt(522400) // 0.5224000% * 100_000_000
	case rewardFruitCount <= 5000:
		fruitRewardsProportion = decimal.NewFromInt(679130) // 0.6791300% * 100_000_000
	case rewardFruitCount <= 6000:
		fruitRewardsProportion = decimal.NewFromInt(842120) // 0.8421200% * 100_000_000
	case rewardFruitCount <= 7500:
		fruitRewardsProportion = decimal.NewFromInt(1190480) // 1.1904800% * 100_000_000
	default:
		fruitRewardsProportion = decimal.NewFromInt(1349210) // 1.3492100% * 100_000_000
	}

	return fruitRewardsProportion
}
