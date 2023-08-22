package wallet

import (
	"github.com/joho/godotenv"
	"log"
	"wallet/config"
)

type Config struct {
	DB  config.Database
	App config.AppConfig
}

func newConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	return Config{
		DB:  config.LoadDatabaseConfiguration(),
		App: config.LoadAppConfig(),
	}
}
