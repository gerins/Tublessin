package domain

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"strconv"
	"time"
	"tublessin/common/model"

	"github.com/go-redis/redis/v8"
)

type ChatUsecase struct {
	ChatRepository ChatRepositoryInterface
	RedisDatabase  *redis.Client
}

type ChatUsecaseInterface interface {
	GetConversation(chat *model.ChatPayload) (*model.ChatConversation, error)
	PostNewConversation(chat *model.ChatPayload) (*model.ChatPayload, error)
}

func NewChatUsecase(db *sql.DB, rdb *redis.Client) ChatUsecaseInterface {
	return &ChatUsecase{ChatRepository: NewChatRepository(db), RedisDatabase: rdb}
}

func (c *ChatUsecase) GetConversation(chat *model.ChatPayload) (*model.ChatConversation, error) {
	senderId, _ := strconv.Atoi(chat.SenderId)
	receiverId, _ := strconv.Atoi(chat.ReceiverId)
	key := (senderId * receiverId) + (senderId + receiverId)

	value, err := c.RedisDatabase.Get(context.Background(), strconv.Itoa(int(key))).Result()
	if err == nil {
		var conversation model.ChatConversation
		json.Unmarshal([]byte(value), &conversation)
		if err != nil {
			log.Println("Something wrong when Unmarshal data to Conversation", err)
		}
		return &conversation, nil
	} else if err != nil && err != redis.Nil {
		log.Println("Something wrong when read data from Redis", err)
	}

	result, err := c.ChatRepository.GetConversation(chat)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	convertResult, err := json.Marshal(result)
	err = c.RedisDatabase.Set(context.Background(), strconv.Itoa(int(key)), convertResult, 3600*time.Second).Err()
	if err != nil {
		log.Println("Cannot save Conversation to Redis", err)
	}

	return result, nil
}

func (c *ChatUsecase) PostNewConversation(chat *model.ChatPayload) (*model.ChatPayload, error) {
	senderId, _ := strconv.Atoi(chat.SenderId)
	receiverId, _ := strconv.Atoi(chat.ReceiverId)
	key := (senderId * receiverId) + (senderId + receiverId)

	result, err := c.ChatRepository.PostNewConversation(chat)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = c.RedisDatabase.Set(context.Background(), strconv.Itoa(int(key)), "", 1*time.Second).Err()
	if err != nil {
		log.Println("Cannot remove Conversation from Redis", err)
	}

	return result, nil
}
