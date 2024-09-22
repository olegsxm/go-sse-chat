package services

import "github.com/olegsxm/go-sse-chat.git/internal/models"

type chatService struct {
}

func (s *chatService) GetChats(userid, lastChatId int) []models.Chat {

	if lastChatId == 0 {
		return repository.Chat.GetChats(userid)
	} else {
		repository.Chat.GetChatsPage(lastChatId)
	}

	return nil
}
