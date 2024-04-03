package eventlistener

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"math/big"
	"os"
	"os/signal"
	"owl-backend/abigen"
	"owl-backend/internal/config"
	"owl-backend/internal/service"
	"owl-backend/pkg/log"
	"syscall"
)

type EventSubscription[T any] struct {
	subscription event.Subscription
	eventChannel chan T
}

var (
	owldinalNftMintBoxEventHash = common.HexToHash("0xf5d3f864a50c2df29b92152f2936fc5520ee555438f668048785c1868cd34230")

	genOneBoxOpenBoxEventHash = common.HexToHash("0x51c074b76ba6958c8e2e5ea62ed9d335b98d5944e5e4e60349cfe2f65588eafc")
	genOneBoxMintBoxEventHash = common.HexToHash("0x92b26e707b024019810eb9f81fb26ab92c318a8efdc77204f029ed774a4ca84b")
)

func StartEventListening() error {
	client, err := ethclient.Dial(config.C.NodeUrl)
	if err != nil {
		log.Fatal("Failed to connect to the Ethereum client: %v", err)
	}

	owldinalNftAddr := common.HexToAddress(config.C.NftOwlAddr)
	owldinalNftContract, err := abigen.NewOwldinal(owldinalNftAddr, client)
	if err != nil {
		log.Fatal("Failed to instantiate contract OwldinalNft: %v", err)
	}

	genOneBoxAddr := common.HexToAddress(config.C.NftMysteryBoxAddr)
	genOneBoxContract, err := abigen.NewOwldinalGenOneBox(genOneBoxAddr, client)
	if err != nil {
		log.Fatal("Failed to instantiate contract GenOneBox: %v", err)
	}

	// start subscribe new events
	//subscribeOpenBox(genOneBoxContract)
	subscribeMintOwldinalNft(owldinalNftContract)
	subscribeMintBox(genOneBoxContract)

	// process past contract events
	eventQuery := ethereum.FilterQuery{
		FromBlock: big.NewInt(config.C.EventStartBlock),
		ToBlock:   nil,
		Addresses: []common.Address{genOneBoxAddr, owldinalNftAddr},
	}
	logs, err := client.FilterLogs(context.Background(), eventQuery)
	for _, vLog := range logs {
		switch vLog.Topics[0] {
		case genOneBoxMintBoxEventHash:
			eventData, err := genOneBoxContract.ParseMintBox(vLog)
			if err != nil {
				log.Warnf("Failed to parse MintBox event: %v", err)
				continue
			}
			handleMintBoxEvent(eventData)
		case owldinalNftMintBoxEventHash:
			eventData, err := owldinalNftContract.ParseMintBox(vLog)
			if err != nil {
				log.Warnf("Failed to parse MintBox event: %v", err)
				continue
			}
			handleOwldinalMintBoxEvent(eventData)
		}

	}

	// Prevent main exit.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	return nil
}

func subscribeMintOwldinalNft(contract *abigen.Owldinal) {
	channel := make(chan *abigen.OwldinalMintBox)
	sub, _ := contract.WatchMintBox(&bind.WatchOpts{Context: context.Background()}, channel, nil)
	eventSubscription := &EventSubscription[*abigen.OwldinalMintBox]{
		subscription: sub,
		eventChannel: channel,
	}
	eventSubscription.StartListening(func(event interface{}) {
		if ev, ok := event.(*abigen.OwldinalMintBox); ok {
			handleOwldinalMintBoxEvent(ev)
		} else {
			log.Warnf("Received unknown event type: %T", event)
		}
	})
}

func subscribeMintBox(contract *abigen.OwldinalGenOneBox) {
	channel := make(chan *abigen.OwldinalGenOneBoxMintBox)
	sub, _ := contract.WatchMintBox(&bind.WatchOpts{Context: context.Background()}, channel, nil)
	eventSubscription := &EventSubscription[*abigen.OwldinalGenOneBoxMintBox]{
		subscription: sub,
		eventChannel: channel,
	}
	eventSubscription.StartListening(func(event interface{}) {
		if ev, ok := event.(*abigen.OwldinalGenOneBoxMintBox); ok {
			handleMintBoxEvent(ev)
		} else {
			log.Warnf("Received unknown event type: %T", event)
		}
	})
}

func subscribeOpenBox(contract *abigen.OwldinalGenOneBox) {
	channel := make(chan *abigen.OwldinalGenOneBoxOpenBox)
	sub, _ := contract.WatchOpenBox(&bind.WatchOpts{Context: context.Background()}, channel, nil)
	eventSubscription := &EventSubscription[*abigen.OwldinalGenOneBoxOpenBox]{
		subscription: sub,
		eventChannel: channel,
	}
	eventSubscription.StartListening(func(event interface{}) {
		if ev, ok := event.(*abigen.OwldinalGenOneBoxOpenBox); ok {
			handleOpenBoxEvent(ev)
		} else {
			log.Warnf("Received unknown event type: %T", event)
		}
	})
}

func (es *EventSubscription[T]) StartListening(handleFunc func(event interface{})) {
	go func() {
		if es.subscription == nil {
			return
		}
		defer es.subscription.Unsubscribe()
		ch := es.eventChannel
		for {
			select {
			case err := <-es.subscription.Err():
				log.Fatal(err)
			case e := <-ch:
				handleFunc(e)
			}
		}
	}()
}

func handleOwldinalMintBoxEvent(event *abigen.OwldinalMintBox) {
	// save event to database
	//service.SaveOpenBoxEvent(event)
	log.Infof("user = %v, box= %v", event.User, event.BoxId.Uint64())
}

func handleMintBoxEvent(event *abigen.OwldinalGenOneBoxMintBox) {
	// save event to database
	//service.SaveOpenBoxEvent(event)
	log.Infof("%v, %v", event.User, *event.TokenId)
}

func handleOpenBoxEvent(event *abigen.OwldinalGenOneBoxOpenBox) {
	// save event to database
	service.SaveOpenBoxEvent(event)

	// TODO: Organize data, merge with historical database records, and calculate totals...

}
