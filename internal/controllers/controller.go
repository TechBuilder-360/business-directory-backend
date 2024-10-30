package controllers

import (
	"github.com/TechBuilder-360/business-directory-backend/internal/common/utils"
	"github.com/TechBuilder-360/business-directory-backend/pkg/log"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Controller interface {
	Ping(ctx *fiber.Ctx) error
	RegisterRoutes(router *fiber.App)
}

func (c *NewController) RegisterRoutes(router *fiber.App) {
	api := router.Group("")

	api.Get("", c.Ping)
}

type NewController struct {
}

func DefaultController() Controller {
	return &NewController{}
}

func (c *NewController) Ping(ctx *fiber.Ctx) error {
	logger := log.LoggerInContext(ctx.UserContext())
	logger.Error("Ping Pong Error")
	logrus.Error("Test Sentry on ping")

	return ctx.Status(http.StatusOK).JSON(utils.Success("We are up and running 🚀", nil, nil))
}
