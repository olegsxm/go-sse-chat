package services

import (
	"errors"

	"github.com/gofiber/fiber/v2/log"
	"github.com/olegsxm/go-sse-chat.git/internal/models"
)

const messages_query_limit = 20

type messageService struct {
}

func (s *messageService) CreateMessage(message models.Message) (models.Message, error) {
	log.Debug("[SERVICE] Create message")

	if message.ChatId == 0 {
		chat, err := s.CreateChat(message)

		if err != nil {
			return models.Message{}, err
		}

		message.ChatId = chat.Id
	}

	message, err := repository.Message.CreateMessage(message)

	if err != nil {
		log.Error("[SERVICE] Create message error", err)
		return message, errors.New("creating message error")
	}

	return message, nil
}

func (s *messageService) CreateChat(message models.Message) (models.Chat, error) {
	log.Debug("[SERVICE] Create chat")
	return repository.Message.CreateChat(message)
}

func (s *messageService) GetMessages(chatId, limit, lastMessageId int) ([]models.Message, error) {
	if limit == 0 {
		limit = messages_query_limit
	}

	if lastMessageId == 0 {
		return repository.Message.GetLastMessages(chatId, limit)
	}

	return nil, nil
}

//func (s *messageService) FindChat(id int64) models.Chat {
//	repository.Message.FindChat(id)
//}
