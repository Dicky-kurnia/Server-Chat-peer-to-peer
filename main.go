package main

import (
	"fmt"
	"jubelio/config"
	"jubelio/controller"
	"jubelio/middleware"
	"jubelio/repository"
	"jubelio/service"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// Connect to the database
	db, err := config.ConnectToSQLDB()
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}
	defer db.Close()
	config.AutoMigrate()

	// Connect to RabbitMQ
	rabbitMQ, err := config.ConnectToRabbitMQ()
	if err != nil {
		fmt.Println("Failed to connect to RabbitMQ:", err)
		return
	}
	defer rabbitMQ.Close()

	// Create new instances of the service
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	messageRepository := repository.NewMessageRepository(db)
	messageService := service.NewMessageService(messageRepository)

	// Create new instances of the controller
	userController := controller.NewUserController(userService)
	messageController := controller.NewMessageController(messageService, rabbitMQ)

	// Create new instance of Fiber app
	app := fiber.New()

	// Register middlewares
	app.Use(middleware.AuthMiddleware)

	// Register routes for user
	app.Post("/register", userController.Register)
	app.Post("/login", userController.Login)

	// Register routes for message
	app.Post("/message", messageController.SendMessage)
	app.Get("/message/:id", messageController.GetMessagesByUserID)

	// Start the server
	app.Listen("3000")
}
