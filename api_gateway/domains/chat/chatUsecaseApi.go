package Chat

import (
	"context"
	"errors"
	"log"
	"strconv"
	"tublessin/common/model"
)

type ChatUsecaseApi struct {
	ChatService model.ChatClient
}

type ChatUsecaseApiInterface interface {
	GetConversation(chat *model.ChatPayload) (*model.ChatConversation, error)
	PostNewConversation(chat *model.ChatPayload) (*model.ChatPayload, error)
}

func NewChatUsecaseApi(ChatService model.ChatClient) ChatUsecaseApiInterface {
	return ChatUsecaseApi{ChatService: ChatService}
}

// Dioper ke Chat-Service untuk ditangani ada di folder services/Chat_services/domain/ChatController.go
func (s ChatUsecaseApi) GetConversation(chat *model.ChatPayload) (*model.ChatConversation, error) {
	_, err := strconv.Atoi(chat.SenderId)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("Sender ID cannot Empty or Must a valid Number")
	}

	_, err = strconv.Atoi(chat.ReceiverId)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("Receiver ID cannot Empty or Must a valid Number")
	}

	result, err := s.ChatService.GetConversation(context.Background(), chat)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return result, nil
}

func (s ChatUsecaseApi) PostNewConversation(chat *model.ChatPayload) (*model.ChatPayload, error) {
	_, err := strconv.Atoi(chat.SenderId)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("Sender ID cannot Empty or Must a valid Number")
	}

	_, err = strconv.Atoi(chat.ReceiverId)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("Receiver ID cannot Empty or Must a valid Number")
	}

	if chat.Message == "" {
		return nil, errors.New("Payload message is Empty or Message size more than 300 character")
	}

	result, err := s.ChatService.PostNewConversation(context.Background(), chat)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return result, nil
}
