package eventlistener

import (
	"github.com/ethereum/go-ethereum/core/types"
	"owl-backend/abigen"
	"owl-backend/pkg/log"
)

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
	log.Infof("[%v-%v] Mint box: user = %v, boxId = %v", event.Raw.TxHash, event.Raw.Index, event.User, event.BoxId.Uint64())
}

type OwldinalNftTransferHandler struct{}

func (h *OwldinalNftTransferHandler) Handle(vlog types.Log) error {
	event, err := owldinalNftContract.ParseTransfer(vlog)
	if err != nil {
		return err
	}

	log.Infof("[%v-%v] Transfer from = %v , to = %v , box= %v", event.Raw.TxHash, event.Raw.Index, event.From, event.To, event.TokenId)
	return nil
}
