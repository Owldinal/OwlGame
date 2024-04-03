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

	HCaptchaKey     string `env:"HCAPTCHA_KEY" envDefault:""`
	DisableHCaptcha bool   `env:"DISABLE_HCAPTCHA" envDefault:"false"`

	DBHost     string `env:"DB_HOST" envDefault:"localhost"`
	DBPort     int    `env:"DB_PORT" envDefault:"3306"`
	DBUser     string `env:"DB_USER" envDefault:"root"`
	DBPassword string `env:"DB_PASSWORD"`
	DBName     string `env:"DB_NAME" envDefault:"mydatabase"`
	LogLevel   string `env:"LOG_LEVEL" envDefault:"info"`
	LogOutput  string `env:"LOG_OUTPUT" envDefault:"console"` // "file" or "console"
	LogFile    string `env:"LOG_FILE" envDefault:"app.log"`   // Only required if LogOutput is set to "file"

	NodeUrl         string `env:"NODE_URL" envDefault:""`
	EventStartBlock int64  `env:"EVENT_START_BLOCK" envDefault:"0"`

	NftOwlAddr        string `env:"NFT_OWL_ADDR" envDefault:""`
	OwlTokenAddr      string `env:"OWL_TOKEN_ADDR" envDefault:""`
	NftMysteryBoxAddr string `env:"NFT_MYSTERY_BOX_ADDR" envDefault:""`
	OwlGameAddr       string `env:"OWL_GAME_ADDR" envDefault:""`
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
