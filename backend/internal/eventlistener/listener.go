package eventlistener

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"os"
	"os/signal"
	"owl-backend/abigen"
	"owl-backend/internal/config"
	"owl-backend/pkg/log"
	"syscall"
	"time"
)

var (
	owldinalNftContract *abigen.Owldinal
	genOneBoxContract   *abigen.OwldinalGenOneBox
	owlGameContract     *abigen.OwlGame
)

const (
	// Contract Owldinal
	eventOwldinalMintBox             = "0xf5d3f864a50c2df29b92152f2936fc5520ee555438f668048785c1868cd34230"
	eventOwldinalRoleAdminChanged    = "0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff"
	eventOwldinalRoleGranted         = "0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d"
	eventOwldinalRoleRevoked         = "0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b"
	eventOwldinalApproval            = "0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"
	eventOwldinalApprovalForAll      = "0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31"
	eventOwldinalBatchMetadataUpdate = "0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c"
	eventOwldinalMetadataUpdate      = "0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7"
	eventOwldinalTransfer            = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"

	// Contract OwldinalGenOneBox
	eventOwldinalGenOneBoxRoleGranted      = "0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d"
	eventOwldinalGenOneBoxRoleRevoked      = "0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b"
	eventOwldinalGenOneBoxTransfer         = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
	eventOwldinalGenOneBoxApproval         = "0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"
	eventOwldinalGenOneBoxApprovalForAll   = "0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31"
	eventOwldinalGenOneBoxMintBox          = "0x4cce2d7ca388465a90e71f76235d389abe1ede028b09c07d4f86519e5adb078c"
	eventOwldinalGenOneBoxRoleAdminChanged = "0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff"

	// Contract OwlGame
	eventOwlGameStakeOwldinalNft   = "0x292b69a5590aefdf5de5c9da21ea45b29afd0635e4d0c7d149d1d84be9224106"
	eventOwlGameUnstakeMysteryBox  = "0x363717deff436618c48fd125aa63e987b64c96ed5976f67cce76223bc8ab2a29"
	eventOwlGameBindInvitation     = "0x3144fc7320adf238514c761c91814e301f210c0e4e0bb5f9fd88bec051b4f100"
	eventOwlGameJoinGame           = "0x3518b1830e6ec9f510e24a95f032e27013c745334b813b19f15405c90923f4bc"
	eventOwlGamePrizePoolIncreased = "0x3c231f64b16483d16ca517ccd881f34f86dd04ccc1305e65b854fe82189b7625"
	eventOwlGameRoleAdminChanged   = "0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff"
	eventOwlGameRoleGranted        = "0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d"
	eventOwlGameStakeMysteryBox    = "0x85ee4e66817b80797ec9b3286b977e5b1e7c0890c0341916ff9acaa6f8210b0d"
	eventOwlGameClaimInviterReward = "0x1b066a4c1595fdf3bf1c300bca29f51f237bbc15ee85e815b39befc1bc630cfc"
	eventOwlGamePrizePoolDecreased = "0xf0870791f79c9e860000c2e3005d55ab8944ec1a18a6b61d4ad60ce184654f42"
	eventOwlGameRoleRevoked        = "0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b"
	eventOwlGameUnstakeOwldinalNft = "0x953947fb8c8d5e1bd1fb1561e4a4077c9f02d4952ad1185d28133c0d764d9a5e"
)

func StartEventListening() error {
	// Init client and contracts
	client, err := ethclient.Dial(config.C.NodeUrl)
	if err != nil {
		log.Fatal("Failed to connect to the Ethereum client: %v", err)
	}

	owldinalNftAddr := common.HexToAddress(config.C.NftOwlAddr)
	owldinalNftContract, err = abigen.NewOwldinal(owldinalNftAddr, client)
	if err != nil {
		log.Fatal("Failed to instantiate contract OwldinalNft: ", err)
	}

	genOneBoxAddr := common.HexToAddress(config.C.NftMysteryBoxAddr)
	genOneBoxContract, err = abigen.NewOwldinalGenOneBox(genOneBoxAddr, client)
	if err != nil {
		log.Fatal("Failed to instantiate contract GenOneBox: ", err)
	}

	owlGameAddr := common.HexToAddress(config.C.OwlGameAddr)
	owlGameContract, err = abigen.NewOwlGame(owlGameAddr, client)
	if err != nil {
		log.Fatal("Failed to instantiate contract OwlGame: ", err)
	}

	contractAddress := []common.Address{genOneBoxAddr, owldinalNftAddr, owlGameAddr}

	startBlock := big.NewInt(config.C.EventStartBlock)
	eventProcessor := NewEventProcessor()
	registerHandlers(eventProcessor)
	header, err := getCurrentBlock(client)
	if err != nil {
		log.Fatal("Failed to get the latest block header: %v", err)
	}
	currentBlock := handleHistoryEvents(client, startBlock, header, contractAddress, eventProcessor)
	// Poll the latest events starting from the newest block.
	pollEvents(client, currentBlock, contractAddress, eventProcessor)

	//err = subscribeEvents(client, contractAddress, eventProcessor)
	//if err != nil {
	//	log.Fatal("Failed to subscribe events: %v", err)
	//}

	// Prevent main exit.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	return nil
}

func registerHandlers(eventProcessor *EventProcessor) {
	eventProcessor.RegisterHandler(eventOwldinalMintBox, &OwldinalNftMintBoxHandler{})
	eventProcessor.RegisterHandler(eventOwldinalTransfer, &OwldinalNftTransferHandler{})

}

func getCurrentBlock(client *ethclient.Client) (*big.Int, error) {
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return header.Number, nil
}

// subscribe events ... fuck merlin chain didn't support subscribe event.
func subscribeEvents(
	client *ethclient.Client,
	addresses []common.Address,
	processors *EventProcessor,
) error {
	subs := make([]*EventSubscription, 0, len(addresses))
	for _, addr := range addresses {
		logs := make(chan types.Log)
		query := ethereum.FilterQuery{
			Addresses: []common.Address{addr},
		}
		sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
		if err != nil {
			return err
		}
		subs = append(subs, &EventSubscription{subscription: sub, eventChannel: logs})
	}

	for _, sub := range subs {
		go func(sub *EventSubscription) {
			defer sub.subscription.Unsubscribe()

			logsChan := sub.eventChannel.(chan types.Log)
			for {
				select {
				case err := <-sub.subscription.Err():
					log.Warnf("Subscription error: %v", err)
					return
				case vLog := <-logsChan:
					if handler, exists := processors.handlers[vLog.Topics[0]]; exists {
						if err := handler.Handle(vLog); err != nil {
							log.Warnf("Failed to process log: %v", err)
						}
					} else {
						//log.Warnf("No handler for event: %s", vLog.Topics[0].Hex())
					}
				}
			}
		}(sub)
	}

	return nil
}

func handleHistoryEvents(
	client *ethclient.Client,
	startBlock *big.Int,
	endBlock *big.Int,
	addresses []common.Address,
	processors *EventProcessor,
) *big.Int {

	// Merlin chain only supports pulling events for up to 1024 blocks at a time, so need to loop to fetch them.
	maxBlocks := big.NewInt(1024)
	for startBlock.Cmp(endBlock) < 0 {
		nextBlock := big.NewInt(0).Add(startBlock, maxBlocks)
		if nextBlock.Cmp(endBlock) > 0 {
			nextBlock.Set(endBlock)
		}

		eventQuery := ethereum.FilterQuery{
			FromBlock: startBlock,
			ToBlock:   nextBlock,
			Addresses: addresses,
		}

		logs, err := client.FilterLogs(context.Background(), eventQuery)
		if err != nil {
			log.Fatal("Failed to filter logs: %v", err)
		}

		for _, vLog := range logs {
			if err := processors.ProcessLog(vLog); err != nil {
				log.Warnf("Failed to process log: %v", err)
			}
		}

		startBlock = big.NewInt(0).Add(nextBlock, big.NewInt(1))

		header, err := getCurrentBlock(client)
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

func pollEvents(client *ethclient.Client,
	startBlock *big.Int,
	addresses []common.Address,
	processors *EventProcessor) {

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

				logs, err := client.FilterLogs(context.Background(), ethereum.FilterQuery{
					FromBlock: startBlock,
					ToBlock:   toBlock,
					Addresses: addresses,
				})
				if err != nil {
					log.Infof("Failed to filter logs: %v", err)
					continue
				}

				for _, vLog := range logs {
					if err := processors.ProcessLog(vLog); err != nil {
						log.Infof("Failed to process log: %v", err)
					}
				}
			}
		}
	}()

	return
}
