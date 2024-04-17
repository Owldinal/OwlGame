package main

import (
	"owl-backend/internal/eventlistener"
)

func main() {
	go eventlistener.StartJobListening()
	eventlistener.StartEventListening()
}
