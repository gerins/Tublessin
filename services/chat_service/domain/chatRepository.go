package domain

import (
	"database/sql"
	"strconv"
	"tublessin/common/model"
)

type ChatRepository struct {
	db *sql.DB
}

type ChatRepositoryInterface interface {
	GetConversation(chat *model.ChatPayload) (*model.ChatConversation, error)
	PostNewConversation(chat *model.ChatPayload) (*model.ChatPayload, error)
}

func NewChatRepository(db *sql.DB) ChatRepositoryInterface {
	return &ChatRepository{db}
}

func (c *ChatRepository) GetConversation(chat *model.ChatPayload) (*model.ChatConversation, error) {
	var chatConversation model.ChatConversation

	result, err := c.db.Query(`SELECT * FROM conversations WHERE (sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?) ORDER BY date_created DESC`, chat.SenderId, chat.ReceiverId, chat.ReceiverId, chat.SenderId)
	if err != nil {
		return nil, err
	}

	for result.Next() {
		var chat model.ChatPayload

		err := result.Scan(&chat.Id, &chat.SenderId, &chat.ReceiverId, &chat.Message, &chat.Status, &chat.DateCreated)
		if err != nil {
			return nil, err
		}

		chatConversation.Conversation = append(chatConversation.Conversation, &chat)
	}

	return &chatConversation, nil
}

func (c *ChatRepository) PostNewConversation(chat *model.ChatPayload) (*model.ChatPayload, error) {
	tx, err := c.db.Begin()
	if err != nil {
		return nil, err
	}

	stmnt1, _ := tx.Prepare("INSERT INTO conversations(sender_id, receiver_id, message) VALUE (?,?,?)")
	result, err := stmnt1.Exec(chat.SenderId, chat.ReceiverId, chat.Message)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	lastInsertID, _ := result.LastInsertId()
	chat.Id = strconv.Itoa(int(lastInsertID))

	tx.Commit()
	return chat, nil
}
