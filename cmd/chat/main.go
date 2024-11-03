package main

import (
	"chat/internal/apps/chat"
	"chat/internal/pkg/config"
	_ "chat/internal/pkg/logger"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	chat.Run(cfg)
}
