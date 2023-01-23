package controller

import (
	"jubelio/model"
	"jubelio/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/streadway/amqp"
)

type MessageController struct {
	messageService *service.MessageService
	rabbitMQ       *amqp.Connection
}

func NewMessageController(messageService *service.MessageService, rabbitMQ *amqp.Connection) *MessageController {
	return &MessageController{
		messageService: messageService,
		rabbitMQ:       rabbitMQ,
	}
}
func (c *MessageController) SendMessage(ctx *fiber.Ctx) error {
	var message model.Message
	if err := ctx.BodyParser(&message); err != nil {
		return err
	}

	senderID, err := strconv.Atoi(ctx.Params("sender_id"))
	if err != nil {
		return err
	}

	receiverID, err := strconv.Atoi(ctx.Params("receiver_id"))
	if err != nil {
		return err
	}

	err = c.messageService.SendMessage(senderID, receiverID, message.Text)
	if err != nil {
		return err
	}

	return ctx.Send([]byte("Message sent"))
}

func (c *MessageController) GetMessagesByUserID(ctx *fiber.Ctx) error {
	userID, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return err
	}

	messages, err := c.messageService.GetMessagesByUserID(userID)
	if err != nil {
		return err
	}

	return ctx.JSON(messages)
}
