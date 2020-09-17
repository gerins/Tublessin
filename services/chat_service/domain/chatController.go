package domain

import (
	"context"
	"database/sql"
	"tublessin/common/model"

	"github.com/go-redis/redis/v8"
)

type ChatServer struct {
	ChatUsecase ChatUsecaseInterface
}

func NewChatController(db *sql.DB, rdb *redis.Client) *ChatServer {
	return &ChatServer{ChatUsecase: NewChatUsecase(db, rdb)}
}

func (c *ChatServer) GetConversation(ctx context.Context, param *model.ChatPayload) (*model.ChatConversation, error) {
	result, err := c.ChatUsecase.GetConversation(param)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ChatServer) PostNewConversation(ctx context.Context, param *model.ChatPayload) (*model.ChatPayload, error) {
	result, err := c.ChatUsecase.PostNewConversation(param)
	if err != nil {
		return nil, err
	}

	return result, nil
}
