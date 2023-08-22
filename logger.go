package wallet

import (
	"go.uber.org/zap"
	"log"
	"os"
)

func NewZapLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer logger.Sync() // flushes buffer, if any

	return logger
}
