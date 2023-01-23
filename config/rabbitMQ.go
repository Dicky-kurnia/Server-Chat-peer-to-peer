package config

import (
	"github.com/streadway/amqp"
)

var rabbitMQ *amqp.Connection

func ConnectToRabbitMQ() (*amqp.Connection, error) {
	var err error
	rabbitMQ, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, err
	}
	return rabbitMQ, nil
}
