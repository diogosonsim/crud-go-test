package middlewares

import (
	"github.com/diogosonsim/crud-go-test/utils"
	"github.com/gofiber/fiber/v2"
)

// Middeware to verify if the the token is authenticated to access route
func IsAuthenticated(c *fiber.Ctx) error {

	cookie := c.Cookies("jwt")

	if _, err := utils.ParseJwt(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	return c.Next()
}
