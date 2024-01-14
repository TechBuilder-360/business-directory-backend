package middlewares

import (
	auth_server_sdk "github.com/TechBuilder-360/business-directory-backend/internal/infrastructure/auth_sdk"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func ValidateToken() fiber.Handler {
	return func(c *fiber.Ctx) error {

		token := ExtractBearerToken(c)

		if !auth_server_sdk.New().ValidateToken(c.UserContext(), token) {
			return c.SendStatus(http.StatusOK)
		}

		return c.Next()
	}
}
