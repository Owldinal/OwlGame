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
	handlers map[common.Hash]EventHandler
}

func NewEventProcessor() *EventProcessor {
	return &EventProcessor{
		handlers: make(map[common.Hash]EventHandler),
	}
}

func (p *EventProcessor) RegisterHandler(eventHash string, handler EventHandler) {
	p.handlers[common.HexToHash(eventHash)] = handler
}

func (p *EventProcessor) ProcessLog(log types.Log) error {
	eventHash := log.Topics[0]
	handler, ok := p.handlers[eventHash]
	if !ok {
		//return errors.New("no handler for event")
		return nil
	}
	return handler.Handle(log)
}

type EventSubscription struct {
	subscription event.Subscription
	eventChannel interface{}
}

type SubscribeFunc[T any] func(opts *bind.WatchOpts, channel chan<- T) (event.Subscription, error)
