package repository

import (
	"database/sql"
	"jubelio/model"
)

type MessageRepository struct {
	db *sql.DB
}

func NewMessageRepository(db *sql.DB) *MessageRepository {
	return &MessageRepository{db}
}

func (r *MessageRepository) AddMessage(senderID, receiverID int, message string) error {
	_, err := r.db.Exec("INSERT INTO messages (sender_id, receiver_id, message) VALUES ($1, $2, $3)", senderID, receiverID, message)
	if err != nil {
		return err
	}
	return nil
}

func (r *MessageRepository) GetMessagesByUserID(userID int) ([]*model.Message, error) {
	rows, err := r.db.Query("SELECT id, sender_id, receiver_id, message FROM messages WHERE sender_id = $1 OR receiver_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []*model.Message
	for rows.Next() {
		var message model.Message
		if err := rows.Scan(&message.ID, &message.SenderID, &message.ReceiverID, &message.Message); err != nil {
			return nil, err
		}
		messages = append(messages, &message)
	}
	return messages, nil
}
