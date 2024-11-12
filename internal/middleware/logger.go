package middleware

import (
	"context"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/utils"
	"github.com/TechBuilder-360/business-directory-backend/pkg/log"
	"github.com/gofiber/fiber/v2"
	"strings"
)

const XRequestID = "X-Request-ID"
const RequestID = "Request-ID"

func Logger(c *fiber.Ctx) error {
	// Set a custom header on all responses:
	requestID := utils.GenerateUUID()
	c.Set(XRequestID, requestID)

	ctx := context.Background()

	logger := log.WithField(RequestID, requestID)
	ctx = context.WithValue(ctx, log.LoggerInCtx, logger)

	c.SetUserContext(ctx)

	if strings.Contains(c.Route().Path, "/gql") {
		logger.Info("Request: %s", string(c.Body()))
	}

	return c.Next()
}
