package controller

import (
	"jubelio/model"
	"jubelio/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService}
}

func (c *UserController) Register(ctx *fiber.Ctx) error {
	var user struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := ctx.BodyParser(&user); err != nil {
		return err
	}

	// create new user struct and assign email and password
	newUser := &model.User{
		Email:    user.Email,
		Password: user.Password,
	}

	// call Register function from userService and pass newUser as parameter
	err := c.userService.Register(newUser)
	if err != nil {
		return err
	}
	ctx.Send([]byte("User successfully registered"))
	return nil
}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	var user model.User
	err := ctx.BodyParser(&user)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).SendString(err.Error())
	}
	_, err = c.userService.Login(user.Email, user.Password)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).SendString(err.Error())
	}
	token, err := c.userService.GenerateJWT(user)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).SendString(err.Error())
	}
	return ctx.JSON(fiber.Map{
		"token": token,
	})
}
