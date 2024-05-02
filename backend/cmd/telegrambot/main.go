package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"owl-backend/internal/config"
	"owl-backend/pkg/log"
	"syscall"
	"time"
)

var (
	client *ethclient.Client
	wallet string
)

func init() {
	var err error
	client, err = ethclient.Dial(config.C.NodeUrl)
	if err != nil {
		log.Fatal("Failed to connect to the Ethereum client: %v", err)
	}
	wallet = config.C.BackendWalletAddress
}

func main() {
	ticker := time.NewTicker(10 * time.Minute)
	defer ticker.Stop()

	done := make(chan bool)
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
		<-quit
		done <- true
	}()

	for {
		select {
		case <-ticker.C:
			checkWalletBalance()
		case <-done:
			log.Infof("Job listener stopped.")
		}
	}

}

func checkWalletBalance() {
	account := common.HexToAddress(wallet)
	balanceByWei, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Warnf("Balance : failed to fetch : %v", err)
		return
	}

	balance := decimal.NewFromBigInt(balanceByWei, -18)
	log.Infof("balance is %v", balance)

	if balance.LessThan(decimal.NewFromFloat(0.002)) {
		sendMessage(fmt.Sprintf("Wallet balance remaining %v . Please transfer to: %v", balance, config.C.BackendWalletAddress))
	}
}

func sendMessage(msg string) {
	response, err := http.PostForm(fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", config.C.TgBotApiKey),
		url.Values{
			"chat_id": {config.C.TgBotChatId},
			"text":    {msg},
		})

	if err != nil {
		log.Infof("Error sending message: %v", err)
	} else {
		defer response.Body.Close()
		log.Infof("Message sent: %s", msg)
	}
}
