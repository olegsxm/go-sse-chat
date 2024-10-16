package logger

import (
	"log/slog"
	"os"
	"time"

	"github.com/fatih/color"

	"github.com/MatusOllah/slogcolor"

	_ "github.com/MatusOllah/slogcolor"
)

func Init(productions bool) {
	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}

	var hand slog.Handler

	if productions {
		hand = slog.NewJSONHandler(os.Stdout, opts)
	} else {
		hand = slogcolor.NewHandler(os.Stderr, &slogcolor.Options{
			Level:       slog.LevelDebug,
			SrcFileMode: 1,
			TimeFormat:  time.RFC3339,
			NoColor:     false,
			MsgColor:    color.New(color.FgGreen),
		})
	}

	logger := slog.New(hand)
	slog.SetDefault(logger)
}
