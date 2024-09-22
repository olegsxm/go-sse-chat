package logger

import (
	"github.com/gofiber/fiber/v2/log"
	"go.uber.org/zap"
)

var logger *zap.Logger

func Logger() *zap.Logger {
	if logger == nil {
		log.Info("Initializing logger")
		logger = zap.Must(zap.NewProduction())
		defer logger.Sync()
	}

	return logger
}
