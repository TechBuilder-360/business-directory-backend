package controllers

import (
	"github.com/TechBuilder-360/business-directory-backend/internal/common/constant"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/utils"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Controller interface {
	Ping(ctx *fiber.Ctx) error
	RegisterRoutes(router *fiber.App)
}

func (c *NewController) RegisterRoutes(router *fiber.App) {
	api := router.Group("")

	api.Get("/", c.Ping)
}

type NewController struct {
}

func DefaultController() Controller {
	return &NewController{}
}

func (c *NewController) Ping(ctx *fiber.Ctx) error {
	log.WithFields(log.Fields{constant.RequestIdentifier: utils.GenerateUUID()})

	return ctx.Status(http.StatusOK).JSON(utils.SuccessResponse{
		Status:  true,
		Message: "We are up and running 🚀",
	})
}
