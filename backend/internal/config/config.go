package config

import (
	"fmt"
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

var C Config

// Config holds all the configuration for our application.
type Config struct {
	Host string `env:"HOST" envDefault:"localhost"`
	Port int    `env:"PORT" envDefault:"8080"`

	BackendWalletPrivateKey string `env:"BACKEND_WALLET_PRIVATE_KEY" envDefault:""`

	DBHost     string `env:"DB_HOST" envDefault:"localhost"`
	DBPort     int    `env:"DB_PORT" envDefault:"3306"`
	DBUser     string `env:"DB_USER" envDefault:"root"`
	DBPassword string `env:"DB_PASSWORD"`
	DBName     string `env:"DB_NAME" envDefault:"mydatabase"`
	LogLevel   string `env:"LOG_LEVEL" envDefault:"info"`
	LogOutput  string `env:"LOG_OUTPUT" envDefault:"console"` // "file" or "console"
	LogFile    string `env:"LOG_FILE" envDefault:"app.log"`   // Only required if LogOutput is set to "file"

	NodeUrl    string `env:"NODE_URL" envDefault:""`
	NftOwlAddr string `env:"NFT_OWL_ADDR" envDefault:""`

	OwlTokenAddr string `env:"OWL_TOKEN_ADDR" envDefault:""`
	NftFairyAddr string `env:"NFT_FAIRY_ADDR" envDefault:""`
	NftFruitAddr string `env:"NFT_FRUIT_ADDR" envDefault:""`
	BlindboxAddr string `env:"BLINDBOX_ADDR" envDefault:""`
	OwlGameAddr  string `env:"OWL_GAME_ADDR" envDefault:""`
}

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
	}

	if err := env.Parse(&C); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", C)
}
