package service

import (
	"jubelio/model"
	"jubelio/repository"
)

type MessageService struct {
	messageRepository *repository.MessageRepository
}

func NewMessageService(messageRepository *repository.MessageRepository) *MessageService {
	return &MessageService{
		messageRepository: messageRepository,
	}
}

func (s *MessageService) SendMessage(senderID, receiverID int, message string) error {
	return s.messageRepository.AddMessage(senderID, receiverID, message)
}

func (s *MessageService) GetMessagesByUserID(userID int) ([]*model.Message, error) {
	return s.messageRepository.GetMessagesByUserID(userID)
}
