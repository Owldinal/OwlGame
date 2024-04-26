package eventlistener

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
	"math/big"
	"os"
	"os/signal"
	"owl-backend/abigen"
	"owl-backend/internal/config"
	"owl-backend/internal/database"
	"owl-backend/internal/model"
	"owl-backend/pkg/log"
	"syscall"
	"time"
)

var (
	OwlGameAddr     string
	owldinalNftAddr common.Address
	genOneBoxAddr   common.Address
	owlGameAddr     common.Address

	client *ethclient.Client

	owldinalNftContract *abigen.Owldinal
	genOneBoxContract   *abigen.OwldinalGenOneBox
	owlGameContract     *abigen.OwlGame
	owlTokenContract    *abigen.OwlToken

	contractAddress []common.Address
	eventProcessors *EventProcessor
)

const (
	// Contract Owldinal
	eventOwldinalMintBox  = "0xf5d3f864a50c2df29b92152f2936fc5520ee555438f668048785c1868cd34230"
	eventOwldinalTransfer = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"

	// Contract OwldinalGenOneBox
	eventOwldinalGenOneBoxTransfer = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
	eventOwldinalGenOneBoxMintBox  = "0x4cce2d7ca388465a90e71f76235d389abe1ede028b09c07d4f86519e5adb078c"

	// Contract OwlGame
	eventOwlGameStakeOwldinalNft   = "0x292b69a5590aefdf5de5c9da21ea45b29afd0635e4d0c7d149d1d84be9224106"
	eventOwlGameUnstakeMysteryBox  = "0x363717deff436618c48fd125aa63e987b64c96ed5976f67cce76223bc8ab2a29"
	eventOwlGameBindInvitation     = "0x3144fc7320adf238514c761c91814e301f210c0e4e0bb5f9fd88bec051b4f100"
	eventOwlGameJoinGame           = "0x3518b1830e6ec9f510e24a95f032e27013c745334b813b19f15405c90923f4bc"
	eventOwlGamePrizePoolIncreased = "0x3c231f64b16483d16ca517ccd881f34f86dd04ccc1305e65b854fe82189b7625"

	eventOwlGameRequestMint        = "0x79ab85c2314dba73749962c93d6920c946a5f6f0532f5e675fd7b1ff530ee058"
	eventOwlGameMintMysteryBox     = "0xe8d0b37cb79abcf29fd9f6b0cb6005971fd5480295fc5a1f3b25f90c5faf727e"
	eventOwlGameStakeMysteryBox    = "0xca6d3ba2c5c2c5da541c637a6da9bccafef99671140cd5fd444322bdaca03336"
	eventOwlGamePrizePoolDecreased = "0xf0870791f79c9e860000c2e3005d55ab8944ec1a18a6b61d4ad60ce184654f42"
	eventOwlGameUnstakeOwldinalNft = "0x953947fb8c8d5e1bd1fb1561e4a4077c9f02d4952ad1185d28133c0d764d9a5e"
	eventOwlGameOwlTokenBurned     = "0x818c9d2e3949d2ab3a0c76d31b137496e1b14d58e2b2d746301c2926a31c435d"
	eventOwlGameFruitRewardUpdated = "0x77cbc09843659e18ed105aec15e30ade634e6ca7b6d13e5ac92bb82cea91f8ee"
	eventOwlGameElfRewardUpdated   = "0x9b964cbd458fdb8541cd6349fd809bc91ecc1a53602693a74fd37f5d3916363c"

	eventOwlGameRebateRewardsIncreased    = "0x07e9fcef60e55f4c1c03cfeb57d85db49ecae00bc0367f586d1d67750dcebbe6"
	eventOwlGameUnlockableRebateIncreased = "0x1f464096fc7638b46885d25d4cfc646c289ce897eeb853b8c4c4ace63285b77f"
	eventOwlGameRebateClaimed             = "0x8ec2a9ed236322c625516e43eab59fcb8145e38ee8d489a28d0aacaeecc298d6"
)

func init() {
	var err error
	client, err = ethclient.Dial(config.C.NodeUrl)
	if err != nil {
		log.Fatal("Failed to connect to the Ethereum client: %v", err)
	}

	owldinalNftAddr = common.HexToAddress(config.C.NftOwlAddr)
	owldinalNftContract, err = abigen.NewOwldinal(owldinalNftAddr, client)
	if err != nil {
		log.Fatal("Failed to instantiate contract OwldinalNft: ", err)
	}

	genOneBoxAddr = common.HexToAddress(config.C.NftMysteryBoxAddr)
	genOneBoxContract, err = abigen.NewOwldinalGenOneBox(genOneBoxAddr, client)
	if err != nil {
		log.Fatal("Failed to instantiate contract GenOneBox: ", err)
	}

	owlTokenAddr := common.HexToAddress(config.C.OwlTokenAddr)
	owlTokenContract, err = abigen.NewOwlToken(owlTokenAddr, client)
	if err != nil {
		log.Fatal("Failed to instantiate contract OwlToken: ", err)
	}

	OwlGameAddr = config.C.OwlGameAddr
	owlGameAddr = common.HexToAddress(config.C.OwlGameAddr)
	owlGameContract, err = abigen.NewOwlGame(owlGameAddr, client)
	if err != nil {
		log.Fatal("Failed to instantiate contract OwlGame: ", err)
	}

	contractAddress = []common.Address{genOneBoxAddr, owldinalNftAddr, owlGameAddr}
	eventProcessors = NewEventProcessor()
	registerHandlers(eventProcessors)
}

func StartEventListening() error {
	startBlock := big.NewInt(config.C.EventStartBlock)
	header, err := getCurrentBlock()
	if err != nil {
		log.Fatal("Failed to get the latest block header: %v", err)
	}
	currentBlock := handleHistoryEvents(startBlock, header)
	pollEvents(currentBlock)

	// Prevent main exit.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	return nil
}

func StartEventChecker() error {
	startBlock := big.NewInt(config.C.EventStartBlock)
	header, err := getCurrentBlock()
	if err != nil {
		log.Fatal("Failed to get the latest block header: %v", err)
	}
	_ = checkHistoryEvents(startBlock, header)

	// Prevent main exit.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	return nil
}

func registerHandlers(eventProcessor *EventProcessor) {
	owldinalNftAddr := common.HexToAddress(config.C.NftOwlAddr)
	genOneBoxAddr := common.HexToAddress(config.C.NftMysteryBoxAddr)
	owlGameAddr := common.HexToAddress(config.C.OwlGameAddr)

	eventProcessor.RegisterHandler(owldinalNftAddr, eventOwldinalMintBox, &OwldinalNftMintBoxHandler{})
	eventProcessor.RegisterHandler(owldinalNftAddr, eventOwldinalTransfer, &OwldinalNftTransferHandler{})

	eventProcessor.RegisterHandler(genOneBoxAddr, eventOwldinalGenOneBoxMintBox, &GenOneBoxMintBoxHandler{})
	eventProcessor.RegisterHandler(genOneBoxAddr, eventOwldinalGenOneBoxTransfer, &GenOneBoxTransferHandler{})

	eventProcessor.RegisterHandler(owlGameAddr, eventOwlGameJoinGame, &OwlGameJoinGameHandler{})
	eventProcessor.RegisterHandler(owlGameAddr, eventOwlGameBindInvitation, &OwlGameBindInvitationHandler{})
	eventProcessor.RegisterHandler(owlGameAddr, eventOwlGamePrizePoolIncreased, &OwlGamePrizePoolIncreasedHandler{})
	eventProcessor.RegisterHandler(owlGameAddr, eventOwlGamePrizePoolDecreased, &OwlGamePrizePoolDecreasedHandler{})

	eventProcessor.RegisterHandler(owlGameAddr, eventOwlGameRequestMint, &OwlGameRequestMintHandler{})
	eventProcessor.RegisterHandler(owlGameAddr, eventOwlGameMintMysteryBox, &OwlGameMintMysteryBoxHandler{})

	eventProcessor.RegisterHandler(owlGameAddr, eventOwlGameStakeOwldinalNft, &OwlGameStakeOwldinalNftHandler{})
	eventProcessor.RegisterHandler(owlGameAddr, eventOwlGameUnstakeOwldinalNft, &OwlGameUnstakeOwldinalNftHandler{})
	eventProcessor.RegisterHandler(owlGameAddr, eventOwlGameStakeMysteryBox, &OwlGameStakeMysteryBoxHandler{})
	eventProcessor.RegisterHandler(owlGameAddr, eventOwlGameUnstakeMysteryBox, &OwlGameUnstakeMysteryBoxHandler{})
	eventProcessor.RegisterHandler(owlGameAddr, eventOwlGameOwlTokenBurned, &OwlGameOwlTokenBurnedHandler{})

	eventProcessor.RegisterHandler(owlGameAddr, eventOwlGameFruitRewardUpdated, &OwlGameFruitRewardUpdatedHandler{})
	eventProcessor.RegisterHandler(owlGameAddr, eventOwlGameElfRewardUpdated, &OwlGameElfRewardUpdatedHandler{})

	eventProcessor.RegisterHandler(owlGameAddr, eventOwlGameRebateRewardsIncreased, &OwlGameRebateRewardsIncreasedHandler{})
	eventProcessor.RegisterHandler(owlGameAddr, eventOwlGameUnlockableRebateIncreased, &OwlGameUnlockableRebateIncreasedHandler{})
	eventProcessor.RegisterHandler(owlGameAddr, eventOwlGameRebateClaimed, &OwlGameClaimRebateClaimedHandler{})
}

func getCurrentBlock() (*big.Int, error) {
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return header.Number, nil
}

func handleHistoryEvents(
	startBlock *big.Int,
	endBlock *big.Int,
) *big.Int {

	// Merlin chain only supports pulling events for up to 1024 blocks at a time, so need to loop to fetch them.
	maxBlocks := big.NewInt(1024)
	for startBlock.Cmp(endBlock) < 0 {
		nextBlock := big.NewInt(0).Add(startBlock, maxBlocks)
		if nextBlock.Cmp(endBlock) > 0 {
			nextBlock.Set(endBlock)
		}

		log.Debugf("Indexer: handle history from %v to %v", startBlock, nextBlock)
		eventQuery := ethereum.FilterQuery{
			FromBlock: startBlock,
			ToBlock:   nextBlock,
			Addresses: contractAddress,
		}

		logs, err := client.FilterLogs(context.Background(), eventQuery)
		if err != nil {
			log.Fatal("Failed to filter logs: %v", err)
		}

		for _, vLog := range logs {
			if err := eventProcessors.ProcessLog(vLog); err != nil {
				log.Warnf("Failed to process log: %v, %v", err, vLog)
			}
		}

		startBlock = big.NewInt(0).Add(nextBlock, big.NewInt(1))

		header, err := getCurrentBlock()
		if err != nil {
			log.Fatal("Failed to get the latest block header: %v", err)
		}
		// If the header is much larger than the endblock, it will result in subsequent subscriptions exceeding 1024
		// blocks and thus throw an exception, so the endblock needs to be updated.
		if header.Cmp(big.NewInt(0).Add(endBlock, big.NewInt(100))) > 0 {
			endBlock = header
		}
	}

	return startBlock
}

func pollEvents(
	startBlock *big.Int) {

	pollInterval := 3 * time.Second
	toBlock := startBlock
	go func() {
		for {
			select {
			case <-time.After(pollInterval):
				startBlock = big.NewInt(0).Add(toBlock, big.NewInt(1))
				endBlock, err := client.BlockNumber(context.Background())
				if err != nil {
					log.Infof("Failed to get latest block number: %v", err)
					continue
				}
				toBlock = big.NewInt(int64(endBlock))
				// if startBlock larger than toBlock, wait next poll
				if toBlock.Cmp(startBlock) < 0 {
					continue
				}

				log.Debugf("Indexer: poll event from %v to %v", startBlock, toBlock)

				logs, err := client.FilterLogs(context.Background(), ethereum.FilterQuery{
					FromBlock: startBlock,
					ToBlock:   toBlock,
					Addresses: contractAddress,
				})
				if err != nil {
					log.Infof("Failed to filter logs: %v", err)
					continue
				}

				for _, vLog := range logs {
					if err := eventProcessors.ProcessLog(vLog); err != nil {
						log.Infof("Failed to process log: %v", err)
					}
				}
			}
		}
	}()

	return
}

func ProcessLogs(startBlock *big.Int, endBlock *big.Int) {
	log.Debugf("Indexer: handle history from %v to %v", startBlock, endBlock)
	eventQuery := ethereum.FilterQuery{
		FromBlock: startBlock,
		ToBlock:   endBlock,
		Addresses: contractAddress,
	}

	logs, err := client.FilterLogs(context.Background(), eventQuery)
	if err != nil {
		log.Fatal("Failed to filter logs: %v", err)
	}

	for _, vLog := range logs {
		if err := eventProcessors.ProcessLog(vLog); err != nil {
			log.Warnf("Failed to process log: %v, %v", err, vLog)
		}
	}
}

func checkHistoryEvents(
	startBlock *big.Int,
	endBlock *big.Int,
) *big.Int {
	// Merlin chain only supports pulling events for up to 1024 blocks at a time, so need to loop to fetch them.
	maxBlocks := big.NewInt(1024)
	for startBlock.Cmp(endBlock) < 0 {
		nextBlock := big.NewInt(0).Add(startBlock, maxBlocks)
		if nextBlock.Cmp(endBlock) > 0 {
			nextBlock.Set(endBlock)
		}

		log.Debugf("Indexer: handle history from %v to %v", startBlock, nextBlock)
		eventQuery := ethereum.FilterQuery{
			FromBlock: startBlock,
			ToBlock:   nextBlock,
			Addresses: contractAddress,
		}

		logs, err := client.FilterLogs(context.Background(), eventQuery)
		if err != nil {
			log.Fatal("Failed to filter logs: %v", err)
		}

		for _, vLog := range logs {
			checkLogExist(vLog)
		}

		startBlock = big.NewInt(0).Add(nextBlock, big.NewInt(1))

		header, err := getCurrentBlock()
		if err != nil {
			log.Fatal("Failed to get the latest block header: %v", err)
		}
		// If the header is much larger than the endblock, it will result in subsequent subscriptions exceeding 1024
		// blocks and thus throw an exception, so the endblock needs to be updated.
		if header.Cmp(big.NewInt(0).Add(endBlock, big.NewInt(100))) > 0 {
			endBlock = header
		}
	}

	return startBlock
}

func checkLogExist(vlog types.Log) {
	eventHash := vlog.Topics[0].Hex()
	var err error
	if vlog.Address == owlGameAddr {
		switch eventHash {
		case eventOwlGameJoinGame:
			event, _ := owlGameContract.ParseJoinGame(vlog)
			item := model.OwlGameJoinGameEvent{
				Event: model.NewEvent(&event.Raw),
			}
			err = database.DB.Where(&item).First(&item).Error
		case eventOwlGameBindInvitation:
			event, _ := owlGameContract.ParseBindInvitation(vlog)
			item := model.OwlGameBindInvitationEvent{
				Event: model.NewEvent(&event.Raw),
			}
			err = database.DB.Where(&item).First(&item).Error
		case eventOwlGamePrizePoolIncreased:
			event, _ := owlGameContract.ParsePrizePoolIncreased(vlog)
			item := model.OwlGamePrizePoolIncreasedEvent{
				Event: model.NewEvent(&event.Raw),
			}
			err = database.DB.Where(&item).First(&item).Error
		case eventOwlGamePrizePoolDecreased:
			event, _ := owlGameContract.ParsePrizePoolDecreased(vlog)
			item := model.OwlGamePrizePoolDecreasedEvent{
				Event: model.NewEvent(&event.Raw),
			}
			err = database.DB.Where(&item).First(&item).Error
		case eventOwlGameRequestMint:
			event, _ := owlGameContract.ParseRequestMint(vlog)
			item := model.OwlGameRequestMintEvent{
				Event: model.NewEvent(&event.Raw),
			}
			err = database.DB.Where(&item).First(&item).Error
		case eventOwlGameMintMysteryBox:
			event, _ := owlGameContract.ParseMintMysteryBox(vlog)
			item := model.OwlGameMintMysteryBoxEvent{
				Event: model.NewEvent(&event.Raw),
			}
			err = database.DB.Where(&item).First(&item).Error
		case eventOwlGameStakeOwldinalNft:
			event, _ := owlGameContract.ParseStakeOwldinalNft(vlog)
			item := model.OwlGameStakeOwldinalNftEvent{
				Event: model.NewEvent(&event.Raw),
			}
			err = database.DB.Where(&item).First(&item).Error
		case eventOwlGameUnstakeOwldinalNft:
			event, _ := owlGameContract.ParseUnstakeOwldinalNft(vlog)
			item := model.OwlGameUnstakeOwldinalNftEvent{
				Event: model.NewEvent(&event.Raw),
			}
			err = database.DB.Where(&item).First(&item).Error
		case eventOwlGameStakeMysteryBox:
			event, _ := owlGameContract.ParseStakeMysteryBox(vlog)
			item := model.OwlGameStakeMysteryBoxEvent{
				Event: model.NewEvent(&event.Raw),
			}
			err = database.DB.Where(&item).First(&item).Error
		case eventOwlGameUnstakeMysteryBox:
			event, _ := owlGameContract.ParseUnstakeMysteryBox(vlog)
			item := model.OwlGameUnstakeMysteryBoxEvent{
				Event: model.NewEvent(&event.Raw),
			}
			err = database.DB.Where(&item).First(&item).Error
		case eventOwlGameOwlTokenBurned:
			event, _ := owlGameContract.ParseOwlTokenBurned(vlog)
			item := model.OwlGameOwlTokenBurnedEvent{
				Event: model.NewEvent(&event.Raw),
			}
			err = database.DB.Where(&item).First(&item).Error
		case eventOwlGameFruitRewardUpdated:
			event, _ := owlGameContract.ParseFruitRewardUpdated(vlog)
			item := model.OwlGameFruitRewardUpdateEvent{
				Event: model.NewEvent(&event.Raw),
			}
			err = database.DB.Where(&item).First(&item).Error
		case eventOwlGameElfRewardUpdated:
			event, _ := owlGameContract.ParseElfRewardUpdated(vlog)
			item := model.OwlGameElfRewardUpdateEvent{
				Event: model.NewEvent(&event.Raw),
			}
			err = database.DB.Where(&item).First(&item).Error
		case eventOwlGameRebateRewardsIncreased:
			event, _ := owlGameContract.ParseRebateRewardsIncreased(vlog)
			item := model.OwlGameRebateRewardsIncreasedEvent{
				Event: model.NewEvent(&event.Raw),
			}
			err = database.DB.Where(&item).First(&item).Error
		case eventOwlGameUnlockableRebateIncreased:
			event, _ := owlGameContract.ParseUnlockableRebateIncreased(vlog)
			// 这里故意写错的，详见 owl_game_rebate_handler -> OwlGameUnlockableRebateIncreasedHandler
			item := model.OwlGameRebateRewardsIncreasedEvent{
				Event: model.NewEvent(&event.Raw),
			}
			err = database.DB.Where(&item).First(&item).Error
		case eventOwlGameRebateClaimed:
			event, _ := owlGameContract.ParseRebateClaimed(vlog)
			item := model.OwlGameRebateClaimedEvent{
				Event: model.NewEvent(&event.Raw),
			}
			err = database.DB.Where(&item).First(&item).Error
		}

	} else if vlog.Address == owldinalNftAddr {
		switch eventHash {
		case eventOwldinalMintBox:
			event, _ := owldinalNftContract.ParseMintBox(vlog)
			item := model.OwldinalNftMintBoxEvent{
				Event: model.NewEvent(&event.Raw),
			}
			err = database.DB.Where(&item).First(&item).Error
		case eventOwldinalTransfer:
			event, _ := owldinalNftContract.ParseTransfer(vlog)
			item := model.OwldinalNftTransferEvent{
				Event: model.NewEvent(&event.Raw),
			}
			err = database.DB.Where(&item).First(&item).Error
		}

	} else if vlog.Address == genOneBoxAddr {
		if eventHash == eventOwldinalGenOneBoxMintBox {
			event, _ := genOneBoxContract.ParseMintBox(vlog)
			item := model.GenOneBoxMintBoxEvent{
				Event: model.NewEvent(&event.Raw),
			}
			err = database.DB.Where(&item).First(&item).Error

		} else if eventHash == eventOwldinalTransfer {
			event, _ := genOneBoxContract.ParseTransfer(vlog)
			item := model.GenOneBoxTransferEvent{
				Event: model.NewEvent(&event.Raw),
			}
			err = database.DB.Where(&item).First(&item).Error

		}
	}

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		log.Warnf("Missing Event: hash= %v , blockNum= %v", vlog.TxHash, vlog.BlockNumber)
	}
}
