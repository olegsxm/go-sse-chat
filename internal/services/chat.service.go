package services

import "github.com/olegsxm/go-sse-chat.git/internal/repository"

type ChatService struct {
	r *repository.Repository
}

func newChatService(r *repository.Repository) *ChatService {
	return &ChatService{r: r}
}
