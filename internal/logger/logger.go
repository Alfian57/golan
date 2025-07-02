package logger

import (
	"go.uber.org/zap"
)

var Log *zap.SugaredLogger

func Init() {
	log, _ := zap.NewProduction()
	defer log.Sync()
	Log = log.Sugar()
}
