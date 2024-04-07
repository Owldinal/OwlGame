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
)

var (
	owldinalNftContract *abigen.Owldinal
	genOneBoxContract   *abigen.OwldinalGenOneBox
	owlGameContract     *abigen.OwlGame
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

	handleHistoryEvents(client, contractAddress, startBlock, eventProcessor)
	err = subscribeEvents(client, contractAddress, eventProcessor)
	if err != nil {
		log.Fatal("Failed to subscribe events: %v", err)
	}

	// Prevent main exit.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	return nil
}

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
	addresses []common.Address,
	startBlock *big.Int,
	processors *EventProcessor,
) {
	// process past contract events
	eventQuery := ethereum.FilterQuery{
		FromBlock: startBlock,
		ToBlock:   nil,
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
}

func registerHandlers(eventProcessor *EventProcessor) {
	//eventProcessor.RegisterHandler(genOneBoxMintBoxEventHash, &GenOneBoxMintBoxHandler{})
	eventProcessor.RegisterHandler("0xf5d3f864a50c2df29b92152f2936fc5520ee555438f668048785c1868cd34230", &OwldinalNftMintBoxHandler{})
}

type OwldinalNftMintBoxHandler struct{}

func (h *OwldinalNftMintBoxHandler) Handle(vlog types.Log) error {
	event, err := owldinalNftContract.ParseMintBox(vlog)
	if err != nil {
		return err
	}
	handleOwldinalMintBoxEvent(event)
	return nil
}

func handleOwldinalMintBoxEvent(event *abigen.OwldinalMintBox) {
	// save event to database
	//service.SaveOpenBoxEvent(event)
	log.Infof("user = %v, box= %v", event.User, event.BoxId.Uint64())
}

//type GenOneBoxMintBoxHandler struct{}
//
//func (h *GenOneBoxMintBoxHandler) Handle(vlog types.Log) error {
//	event, err := owldinalNftContract.ParseMintBox(vlog)
//	if err != nil {
//		return err
//	}
//	handleOwldinalMintBoxEvent(event)
//	return nil
//}
//func handleMintBoxEvent(event *abigen.OwldinalGenOneBoxMintBox) {
//	save event to database
//	service.SaveOpenBoxEvent(event)
//
//}
