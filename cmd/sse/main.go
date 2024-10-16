package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/olegsxm/go-sse-chat.git/pkg/logger"

	_ "github.com/olegsxm/go-sse-chat.git/docs"
	"github.com/olegsxm/go-sse-chat.git/internal/apps/sse"
	"github.com/olegsxm/go-sse-chat.git/internal/config"
	_ "github.com/olegsxm/go-sse-chat.git/pkg/logger"
)

//	@title		Chat API
//	@version	1.0
//
// @host localhost:3000
// @BasePath  /api/v1
func main() {
	cfg, err := config.New()
	if err != nil || cfg == nil {
		slog.Error("Error loading config")
		panic(err)
	}

	logger.Init(cfg.Production)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	server := sse.New(ctx, cfg)
	// Start server
	go func() {
		var err error

		if cfg.Production {
			err = runProdServer(server)
		} else {
			err = runDevServer(server)
		}

		if err = server.ListenAndServeTLS("cert.pem", "key.pem"); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with a timeout of 10 seconds.
	<-ctx.Done()

	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(c); err != nil {
		slog.Error(err.Error())
	}
}

func runProdServer(s *http.Server) error {
	return s.ListenAndServeTLS("cert.pem", "key.pem")
}

func runDevServer(s *http.Server) error {
	return s.ListenAndServe()
}
