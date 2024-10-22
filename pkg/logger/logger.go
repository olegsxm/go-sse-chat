package logger

import (
	"github.com/MatusOllah/slogcolor"
	"github.com/fatih/color"
	"log/slog"
	"os"
	"time"
)

var opts = &slog.HandlerOptions{
	AddSource: true,
	Level:     slog.LevelDebug,
}

func NewLogger() *slog.Logger {
	hand := slogcolor.NewHandler(os.Stderr, &slogcolor.Options{
		Level:       slog.LevelDebug,
		SrcFileMode: 1,
		TimeFormat:  time.RFC3339,
		NoColor:     false,
		MsgColor:    color.New(color.BgCyan),
	})

	// TODO implement handler for production

	var logger = slog.New(hand)

	return logger
}
