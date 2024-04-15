package eventlistener

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

type EventHandler interface {
	Handle(eventLog types.Log) error
}

type EventProcessor struct {
	handlers map[common.Address]map[common.Hash]EventHandler
}

func NewEventProcessor() *EventProcessor {
	return &EventProcessor{
		handlers: make(map[common.Address]map[common.Hash]EventHandler),
	}
}

func (p *EventProcessor) RegisterHandler(contractAddress common.Address, eventHash string, handler EventHandler) {
	if p.handlers[contractAddress] == nil {
		p.handlers[contractAddress] = make(map[common.Hash]EventHandler)
	}
	p.handlers[contractAddress][common.HexToHash(eventHash)] = handler
}

func (p *EventProcessor) ProcessLog(log types.Log) error {
	handlerMap, ok := p.handlers[log.Address]
	if !ok {
		return nil
	}

	eventHash := log.Topics[0]
	handler, ok := handlerMap[eventHash]
	if !ok {
		return nil
	}

	return handler.Handle(log)
}

type EventSubscription struct {
	subscription event.Subscription
	eventChannel interface{}
}

type SubscribeFunc[T any] func(opts *bind.WatchOpts, channel chan<- T) (event.Subscription, error)
