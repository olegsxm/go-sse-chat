package main

import (
	"context"
	_ "github.com/olegsxm/go-sse-chat.git/docs"
	"github.com/olegsxm/go-sse-chat.git/internal/apps/sse"
	_ "github.com/olegsxm/go-sse-chat.git/pkg/logger"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"
)

//	@title		Chat API
//	@version	1.0
//
// @host localhost:443
// @BasePath  /api/v1
func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	server := sse.New(ctx)
	// Start server
	go func() {
		if err := server.ListenAndServeTLS("cert.pem", "key.pem"); err != nil && err != http.ErrServerClosed {
			slog.Error("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		slog.Error(err.Error())
	}
}
