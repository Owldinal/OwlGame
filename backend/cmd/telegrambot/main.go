package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"owl-backend/abigen"
	"owl-backend/internal/config"
	"owl-backend/internal/eth"
	"owl-backend/pkg/log"
	"syscall"
	"time"
)

var (
	wallet           string
	owlTokenContract *abigen.OwlToken
)

func init() {
	var err error
	owlTokenContract, err = abigen.NewOwlToken(common.HexToAddress(config.C.OwlTokenAddr), eth.Client)
	if err != nil {
		log.Fatalf("Failed init telegram bot: %v", err)
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
			checkOwldinalBalance()
		case <-done:
			log.Infof("Job listener stopped.")
		}
	}
}

func checkWalletBalance() {
	account := common.HexToAddress(wallet)
	balanceByWei, err := eth.Client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Warnf("Balance : failed to fetch : %v", err)
		sendMessage(fmt.Sprintf("Failed to fetch wallet balance"))
		return
	}

	balance := decimal.NewFromBigInt(balanceByWei, -18)
	log.Infof("balance is %v", balance)

	if balance.LessThan(decimal.NewFromFloat(0.002)) {
		sendMessage(fmt.Sprintf("Wallet balance remaining %v . Please transfer BTC to: %v", balance, config.C.BackendWalletAddress))
	}
}

func checkOwldinalBalance() {
	account := common.HexToAddress(wallet)
	balanceByWei, err := owlTokenContract.BalanceOf(&bind.CallOpts{}, account)
	if err != nil {
		log.Warnf("Balance : failed to fetch : %v", err)
		sendMessage(fmt.Sprintf("Failed to fetch Owldinal balance"))
		return
	}

	balance := decimal.NewFromBigInt(balanceByWei, -18)
	log.Infof("owl token is %v", balance)

	if balance.LessThan(decimal.NewFromFloat(10000000)) {
		sendMessage(fmt.Sprintf("OwlToken remaining %v . Please transfer more OwlToken to: %v", balance.IntPart(), config.C.BackendWalletAddress))
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
