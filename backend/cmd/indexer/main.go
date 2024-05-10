package main

import (
	"owl-backend/internal/eventlistener"
)

func main() {
	//go eventlistener.StartJobListening()
	go eventlistener.StartEventCheckerWithDelay()
	_ = eventlistener.StartEventListening()
}
