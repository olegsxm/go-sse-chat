package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"

	"github.com/olegsxm/go-sse-chat.git/internal/apps/sse"

	_ "github.com/olegsxm/go-sse-chat.git/pkg/logger"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := sse.Run(ctx); err != nil {
		slog.Error("sse run error: ", err.Error())
	}
}
