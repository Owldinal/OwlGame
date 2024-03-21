package eth

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"owl-backend/internal/config"
	"time"
)

var (
	Client *ethclient.Client
)

func init() {
	client, err := ethclient.Dial(config.C.NodeUrl)
	if err != nil {
		panic(err)
	}

	Client = client
	fmt.Printf("%v eth init successfully\n", time.Now())
}
