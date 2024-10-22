package main

import (
	"github.com/olegsxm/go-sse-chat/internal/apps/chat_app"
	"go.uber.org/fx"
)

//	@title		Chat API
//	@version	1.0
//
// @host localhost:3000
// @BasePath  /api/v1
func main() {
	fx.New(chat_app.CreateChatApp(), fx.NopLogger).Run()
}
