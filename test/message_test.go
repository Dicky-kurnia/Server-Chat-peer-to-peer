package test

import (
	"jubelio/config"
	"jubelio/model"
	"jubelio/repository"
	"jubelio/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMessageService(t *testing.T) {

	db, err := config.ConnectToSQLDB()
	if err != nil {
		t.Error("Failed to connect to the database:", err)
		return
	}
	defer db.Close()

	config.AutoMigrate()

	messageRepo := repository.NewMessageRepository(db)
	messageService := service.NewMessageService(messageRepo)

	t.Run("Test send message", func(t *testing.T) {
		// create new message struct
		newMessage := &model.Message{
			SenderID:   1,
			ReceiverID: 2,
			Text:       "Hello",
		}

		// call SendMessage function from messageService and pass newMessage as parameter
		err := messageService.SendMessage(1, 2, "Hallo-World")
		if err != nil {
			t.Error("Failed to send message:", err)
		}

		// get message by sender and receiver id
		message, err := messageService.GetMessagesByUserID(1)
		if err != nil {
			t.Error("Failed to get message:", err)
		}

		// assert message text
		assert.Equal(t, newMessage.Text, message[0].Text)
	})

	t.Run("Test get message by user id", func(t *testing.T) {
		// get message by sender and receiver id
		message, err := messageService.GetMessagesByUserID(1)
		if err != nil {
			t.Error("Failed to get message:", err)
		}

		// assert message length
		assert.Len(t, message, 1)
	})
}
