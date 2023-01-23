package middleware

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) {
	token := c.Get("Authorization")
	if token == "" {
		c.SendStatus(fiber.StatusUnauthorized)
		c.JSON(fiber.Map{
			"error": "Authorization token is missing",
		})
		return
	}
	token = token[7:]
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte("secret"), nil
	})
	if err != nil {
		c.SendStatus(fiber.StatusUnauthorized)
		c.JSON(fiber.Map{
			"error": "Invalid token",
		})
		return
	}
	if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {
		c.Locals("user_id", int(claims["user_id"].(float64)))
	} else {
		c.SendStatus(fiber.StatusUnauthorized)
		c.JSON(fiber.Map{
			"error": "Invalid token",
		})
		return
	}
	c.Next()
}
