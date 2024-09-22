package main

import (
	_ "github.com/joho/godotenv/autoload"
	sse_chat "github.com/olegsxm/go-sse-chat.git/internal/apps/sse-chat"
)

func main() {
	sse_chat.Run()
}
