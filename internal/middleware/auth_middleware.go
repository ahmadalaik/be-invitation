package middleware

import (
	"strings"

	"github.com/ahmadalaik/be-invitation/pkg/jwt"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	header := c.Get("Authorization")
	if header == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "authorization header required",
		})
	}

	tokenStr := strings.TrimPrefix(header, "Bearer ")
	if tokenStr == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid token format",
		})
	}

	userID, username, err := jwt.ValidateToken(tokenStr)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid token",
		})
	}

	c.Locals("userID", userID)
	c.Locals("username", username)

	return c.Next()
}
