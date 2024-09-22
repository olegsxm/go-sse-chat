package main

import (
	_ "github.com/joho/godotenv/autoload"
	ws_chat "github.com/olegsxm/go-sse-chat.git/internal/apps/ws-chat"
)

func main() {
	ws_chat.Run()
}
