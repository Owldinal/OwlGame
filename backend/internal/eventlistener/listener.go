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
	openBoxEventHash = common.HexToHash("0x51c074b76ba6958c8e2e5ea62ed9d335b98d5944e5e4e60349cfe2f65588eafc")
	mintBoxEventHash = common.HexToHash("0x92b26e707b024019810eb9f81fb26ab92c318a8efdc77204f029ed774a4ca84b")
)

func StartEventListening() error {
	client, err := ethclient.Dial(config.C.NodeUrl)
	if err != nil {
		log.Fatal("Failed to connect to the Ethereum client: %v", err)
	}
	blindBoxAddr := common.HexToAddress(config.C.BlindboxAddr)
	blindBoxContract, err := abigen.NewOwldinalGenOneBox(blindBoxAddr, client)
	if err != nil {
		log.Fatal("Failed to instantiate contract blindBox: %v", err)
	}

	// start subscribe new events
	subscribeOpenBox(blindBoxContract)

	// process past contract events
	eventQuery := ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		ToBlock:   nil,
		Addresses: []common.Address{blindBoxAddr},
	}
	logs, err := client.FilterLogs(context.Background(), eventQuery)
	for _, vLog := range logs {
		switch vLog.Topics[0] {
		case openBoxEventHash:
			eventData, err := blindBoxContract.ParseOpenBox(vLog)
			if err != nil {
				log.Warnf("Failed to parse OpenBox event: %v", err)
				continue
			}
			handleOpenBoxEvent(eventData)
		}
	}

	// Prevent main exit.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	return nil
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

func handleOpenBoxEvent(event *abigen.OwldinalGenOneBoxOpenBox) {
	// save event to database
	service.SaveOpenBoxEvent(event)

	// TODO: Organize data, merge with historical database records, and calculate totals...

}
