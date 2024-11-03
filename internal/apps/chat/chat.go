package chat

import (
	database "chat/internal/db"
	"chat/internal/handlers"
	"chat/internal/pkg/config"
	"chat/internal/pkg/server"
	"chat/internal/repository"
	"chat/internal/services"
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Run(cfg config.Config) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	db := database.New(ctx, cfg.DSN())

	srv := server.New()

	api := srv.Group("/api")

	r := repository.New(db)

	s := services.NewServices(r)

	handlers.New(handlers.Dependencies{
		Api:      api,
		Services: s,
	}).Mount()

	go func() {
		if err := srv.Start(cfg.ServerPort); err != nil && !errors.Is(err, http.ErrServerClosed) {
			srv.Logger.Fatal("shutting down the server")
		}
	}()

	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		srv.Logger.Fatal(err)
	}

	defer db.Pg.Close()
}
