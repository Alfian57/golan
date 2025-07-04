package logger

import (
	"log"

	"go.uber.org/zap"
)

var Log *zap.SugaredLogger

func Init() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync()
	Log = logger.Sugar()
}
