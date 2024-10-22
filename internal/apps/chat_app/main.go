package chat_app

import (
	"context"
	"errors"
	"github.com/olegsxm/go-sse-chat/ent"
	"github.com/olegsxm/go-sse-chat/internal/db"
	"github.com/olegsxm/go-sse-chat/internal/repository"
	"github.com/olegsxm/go-sse-chat/internal/router"
	"github.com/olegsxm/go-sse-chat/internal/services"
	"github.com/olegsxm/go-sse-chat/pkg/logger"
	"github.com/olegsxm/go-sse-chat/pkg/server"
	"go.uber.org/fx"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func CreateChatApp() fx.Option {
	return fx.Options(
		fx.Provide(
			logger.NewLogger,
			db.NewDB,
			repository.New,
			services.New,
			router.New,
			server.New,
		),
		fx.Invoke(
			func(l *slog.Logger) {
				slog.SetDefault(l)
				slog.Info("INIT LOGGER")
			},
			func(d *db.Db) {
				slog.Info("INIT DB")
				if err := d.SQL().Schema.Create(context.Background()); err != nil {
					log.Fatalf("failed creating schema resources: %v", err)
				}
			},
			runServer,
		),
	)
}

func runServer(s *http.Server, d *db.Db) {
	shutdownChan := make(chan bool, 1)

	go func() {
		if err := s.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server error: %v", err)
		}

		slog.Debug("IMPLEMENT STOPPING")
		defer func(sql *ent.Client) {
			err := sql.Close()
			if err != nil {
				slog.Error("failed to close DB connection")
			}
		}(d.SQL())
		shutdownChan <- true
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := s.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("HTTP shutdown error: %v", err)
	}

	<-shutdownChan
	slog.Info("Graceful shutdown complete.")
}
