package main

import (
	"fmt"
	"owl-backend/internal/config"
	"owl-backend/internal/server"
	"owl-backend/pkg/log"

	_ "owl-backend/internal/eth"
)

func main() {
	log.Info("service is starting...")
	engine := server.SetupEngine()

	serverAddr := fmt.Sprintf("%v:%v", config.C.Host, config.C.Port)
	log.Infof("Server is starting on %v", serverAddr)
	if err := engine.Run(serverAddr); err != nil {
		log.Fatal("Error starting server: %v", err)
	}
}
