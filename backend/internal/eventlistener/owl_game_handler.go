package eventlistener

import (
	"errors"
	"github.com/ethereum/go-ethereum/core/types"
	"gorm.io/gorm"
	"owl-backend/internal/database"
	"owl-backend/internal/model"
	"owl-backend/pkg/log"
)

type OwlGameJoinGameHandler struct{}

func (h *OwlGameJoinGameHandler) Handle(vlog types.Log) error {
	event, err := owlGameContract.ParseJoinGame(vlog)
	if err != nil {
		return err
	}
	// save event to database
	//log.Infof("[%v-%v] Mint box: user = %v, boxId = %v", event.Raw.TxHash, event.Raw.Index, event.User, event.TokenId.Uint64())
	eventItem := model.OwlGameJoinGameEvent{
		Event:      model.NewEvent(&event.Raw),
		User:       event.User.Hex(),
		InviteCode: event.InviteCode,
	}
	eventResult := database.DB.Clauses().Create(&eventItem)
	if eventResult.Error != nil {
		if errors.Is(eventResult.Error, gorm.ErrDuplicatedKey) {
			return nil
		} else {
			log.Warnf("Error Is: %v", eventResult.Error)
			return eventResult.Error
		}
	}

	// TODO:

	return nil
}

type OwlGameBindInvitationHandler struct{}

func (h *OwlGameBindInvitationHandler) Handle(vlog types.Log) error {
	event, err := owlGameContract.ParseBindInvitation(vlog)
	if err != nil {
		return err
	}
	// save event to database
	//log.Infof("[%v-%v] Mint box: user = %v, boxId = %v", event.Raw.TxHash, event.Raw.Index, event.User, event.TokenId.Uint64())
	eventItem := model.OwlGameBindInvitationEvent{
		Event:   model.NewEvent(&event.Raw),
		Inviter: event.Inviter.Hex(),
		Invitee: event.Invitee.Hex(),
	}
	eventResult := database.DB.Clauses().Create(&eventItem)
	if eventResult.Error != nil {
		if errors.Is(eventResult.Error, gorm.ErrDuplicatedKey) {
			return nil
		} else {
			log.Warnf("Error Is: %v", eventResult.Error)
			return eventResult.Error
		}
	}

	// TODO:

	return nil
}

type OwlGamePrizePoolIncreasedHandler struct{}

func (h *OwlGamePrizePoolIncreasedHandler) Handle(vlog types.Log) error {
	event, err := owlGameContract.ParsePrizePoolIncreased(vlog)
	if err != nil {
		return err
	}
	// save event to database
	//log.Infof("[%v-%v] Mint box: user = %v, boxId = %v", event.Raw.TxHash, event.Raw.Index, event.User, event.TokenId.Uint64())
	eventItem := model.OwlGamePrizePoolIncreasedEvent{
		Event:  model.NewEvent(&event.Raw),
		Amount: database.Amount{Int: event.Amount},
	}
	eventResult := database.DB.Clauses().Create(&eventItem)
	if eventResult.Error != nil {
		if errors.Is(eventResult.Error, gorm.ErrDuplicatedKey) {
			return nil
		} else {
			log.Warnf("Error Is: %v", eventResult.Error)
			return eventResult.Error
		}
	}

	// TODO:

	return nil
}

type OwlGamePrizePoolDecreasedHandler struct{}

func (h *OwlGamePrizePoolDecreasedHandler) Handle(vlog types.Log) error {
	event, err := owlGameContract.ParsePrizePoolDecreased(vlog)
	if err != nil {
		return err
	}
	// save event to database
	//log.Infof("[%v-%v] Mint box: user = %v, boxId = %v", event.Raw.TxHash, event.Raw.Index, event.User, event.TokenId.Uint64())
	eventItem := model.OwlGamePrizePoolDecreasedEvent{
		Event:  model.NewEvent(&event.Raw),
		Amount: database.Amount{Int: event.Amount},
	}
	eventResult := database.DB.Clauses().Create(&eventItem)
	if eventResult.Error != nil {
		if errors.Is(eventResult.Error, gorm.ErrDuplicatedKey) {
			return nil
		} else {
			log.Warnf("Error Is: %v", eventResult.Error)
			return eventResult.Error
		}
	}

	// TODO:

	return nil
}

type OwlGameStakeOwldinalNftHandler struct{}

func (h *OwlGameStakeOwldinalNftHandler) Handle(vlog types.Log) error {
	event, err := owlGameContract.ParseStakeOwldinalNft(vlog)
	if err != nil {
		return err
	}
	// save event to database
	//log.Infof("[%v-%v] Mint box: user = %v, boxId = %v", event.Raw.TxHash, event.Raw.Index, event.User, event.TokenId.Uint64())
	tokenIds := make([]uint64, len(event.TokenId))
	for i, bigInt := range event.TokenId {
		tokenIds[i] = bigInt.Uint64()
	}
	eventItem := model.OwlGameStakeOwldinalNftEvent{
		Event:    model.NewEvent(&event.Raw),
		User:     event.User.Hex(),
		TokenIds: tokenIds,
	}
	eventResult := database.DB.Clauses().Create(&eventItem)
	if eventResult.Error != nil {
		if errors.Is(eventResult.Error, gorm.ErrDuplicatedKey) {
			return nil
		} else {
			log.Warnf("Error Is: %v", eventResult.Error)
			return eventResult.Error
		}
	}

	// TODO:

	return nil
}

type OwlGameUnstakeOwldinalNftHandler struct{}

func (h *OwlGameUnstakeOwldinalNftHandler) Handle(vlog types.Log) error {
	event, err := owlGameContract.ParseUnstakeOwldinalNft(vlog)
	if err != nil {
		return err
	}
	// save event to database
	//log.Infof("[%v-%v] Mint box: user = %v, boxId = %v", event.Raw.TxHash, event.Raw.Index, event.User, event.TokenId.Uint64())
	tokenIds := make([]uint64, len(event.TokenId))
	for i, bigInt := range event.TokenId {
		tokenIds[i] = bigInt.Uint64()
	}
	eventItem := model.OwlGameUnstakeOwldinalNftEvent{
		Event:    model.NewEvent(&event.Raw),
		User:     event.User.Hex(),
		TokenIds: tokenIds,
	}
	eventResult := database.DB.Clauses().Create(&eventItem)
	if eventResult.Error != nil {
		if errors.Is(eventResult.Error, gorm.ErrDuplicatedKey) {
			return nil
		} else {
			log.Warnf("Error Is: %v", eventResult.Error)
			return eventResult.Error
		}
	}

	// TODO:

	return nil
}

type OwlGameStakeMysteryBoxHandler struct{}

func (h *OwlGameStakeMysteryBoxHandler) Handle(vlog types.Log) error {
	event, err := owlGameContract.ParseStakeMysteryBox(vlog)
	if err != nil {
		return err
	}
	// save event to database
	//log.Infof("[%v-%v] Mint box: user = %v, boxId = %v", event.Raw.TxHash, event.Raw.Index, event.User, event.TokenId.Uint64())
	tokenIds := make([]uint64, len(event.TokenId))
	for i, bigInt := range event.TokenId {
		tokenIds[i] = bigInt.Uint64()
	}
	eventItem := model.OwlGameStakeMysteryBoxEvent{
		Event:    model.NewEvent(&event.Raw),
		User:     event.User.Hex(),
		TokenIds: tokenIds,
	}
	eventResult := database.DB.Clauses().Create(&eventItem)
	if eventResult.Error != nil {
		if errors.Is(eventResult.Error, gorm.ErrDuplicatedKey) {
			return nil
		} else {
			log.Warnf("Error Is: %v", eventResult.Error)
			return eventResult.Error
		}
	}

	// TODO:

	return nil
}

type OwlGameUnstakeMysteryBoxHandler struct{}

func (h *OwlGameUnstakeMysteryBoxHandler) Handle(vlog types.Log) error {
	event, err := owlGameContract.ParseUnstakeMysteryBox(vlog)
	if err != nil {
		return err
	}
	// save event to database
	//log.Infof("[%v-%v] Mint box: user = %v, boxId = %v", event.Raw.TxHash, event.Raw.Index, event.User, event.TokenId.Uint64())
	eventItem := model.OwlGameUnstakeMysteryBoxEvent{
		Event:   model.NewEvent(&event.Raw),
		User:    event.User.Hex(),
		TokenId: event.TokenId.Uint64(),
		Rewards: database.Amount{Int: event.Rewards},
	}
	eventResult := database.DB.Clauses().Create(&eventItem)
	if eventResult.Error != nil {
		if errors.Is(eventResult.Error, gorm.ErrDuplicatedKey) {
			return nil
		} else {
			log.Warnf("Error Is: %v", eventResult.Error)
			return eventResult.Error
		}
	}

	// TODO:

	return nil
}

type OwlGameClaimInviterRewardsHandler struct{}

func (h *OwlGameClaimInviterRewardsHandler) Handle(vlog types.Log) error {
	event, err := owlGameContract.ParseClaimInviterReward(vlog)
	if err != nil {
		return err
	}
	// save event to database
	//log.Infof("[%v-%v] Mint box: user = %v, boxId = %v", event.Raw.TxHash, event.Raw.Index, event.User, event.TokenId.Uint64())
	eventItem := model.OwlGameClaimInviterRewardEvent{
		Event:          model.NewEvent(&event.Raw),
		User:           event.User.Hex(),
		WithdrawAmount: database.Amount{Int: event.WithdrawAmount},
	}
	eventResult := database.DB.Clauses().Create(&eventItem)
	if eventResult.Error != nil {
		if errors.Is(eventResult.Error, gorm.ErrDuplicatedKey) {
			return nil
		} else {
			log.Warnf("Error Is: %v", eventResult.Error)
			return eventResult.Error
		}
	}

	// TODO:

	return nil
}
