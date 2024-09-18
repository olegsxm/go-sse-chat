package services

type chatService struct {
}

func (s *chatService) GetChats() {
	repository.Chat.GetChats()
}
