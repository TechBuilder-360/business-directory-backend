package controllers

import (
	"github.com/TechBuilder-360/business-directory-backend/internal/common/apiError"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/constant"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/types"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/utils"
	"github.com/TechBuilder-360/business-directory-backend/internal/domain/user"
	"github.com/TechBuilder-360/business-directory-backend/internal/services"
	"github.com/TechBuilder-360/business-directory-backend/internal/validation"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type IAuthController interface {
	Registration(ctx *fiber.Ctx) error
	ActivateAccount(ctx *fiber.Ctx) error
	Logout(ctx *fiber.Ctx) error
	Authenticate(ctx *fiber.Ctx) error
	RefreshToken(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	RegisterRoutes(router *fiber.App)
}

type AuthController struct {
	as services.IAuthService
}

func DefaultAuthController() IAuthController {
	return &AuthController{
		as: services.NewAuthService(),
	}
}

func (c *AuthController) RegisterRoutes(router *fiber.App) {
	auth := router.Group("/auth")

	auth.Post("/registration", c.Registration)

}

func (c *AuthController) Registration(ctx *fiber.Ctx) error {
	logger := log.WithFields(log.Fields{constant.RequestIdentifier: utils.GenerateUUID()})
	logger.Info("Registration")

	body := new(types.Registration)
	err := ctx.BodyParser(body)
	if err != nil {
		logger.Error(err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(utils.Error(apiError.AppError{
			Message: "request failed",
			Error:   err.Error(),
		}))
	}

	if valErr, ok := validation.ValidateStruct(body, logger); !ok {
		return ctx.Status(http.StatusBadRequest).JSON(utils.Error(apiError.AppError{
			Message: "request failed",
			Error:   valErr,
		}))
	}

	err = c.as.Registration(ctx.UserContext(), *body, nil)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.Error(apiError.AppError{
			Message: "request failed",
			Error:   err.Error(),
		}))
	}

	return ctx.Status(http.StatusOK).JSON(
		utils.Success("Registration is successful, An email has been sent to activate your account", nil, nil),
	)
}

func (c *AuthController) ActivateAccount(ctx *fiber.Ctx) error {
	logger := log.WithFields(log.Fields{constant.RequestIdentifier: utils.GenerateUUID()})
	logger.Info("ActivateAccount")

	token := user.ExtractBearerToken(ctx)

	err := c.as.ActivateAccount(ctx.UserContext(), token, logger)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.Error(apiError.AppError{
			Message: "request failed",
			Error:   err.Error(),
		}))
	}

	return ctx.Status(http.StatusOK).JSON(utils.Success("email verification successful", nil, nil))
}

func (c *AuthController) Logout(ctx *fiber.Ctx) error {
	logger := log.WithFields(log.Fields{constant.RequestIdentifier: utils.GenerateUUID()})
	logger.Info("Logout")

	token := user.ExtractBearerToken(ctx)

	err := c.as.Logout(ctx.UserContext(), token)
	if err != nil {
		return ctx.SendStatus(http.StatusBadRequest)
	}

	return ctx.SendStatus(http.StatusOK)
}

func (c *AuthController) Authenticate(ctx *fiber.Ctx) error {
	logger := log.WithFields(log.Fields{constant.RequestIdentifier: utils.GenerateUUID()})
	logger.Info("Authenticate")

	body := types.Authenticate{}

	err := ctx.BodyParser(body)
	if err != nil {
		logger.Error(err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(utils.Error(apiError.AppError{
			Message: "request failed",
			Error:   err.Error(),
		}))
	}

	if valErr, ok := validation.ValidateStruct(body, logger); !ok {
		return ctx.Status(http.StatusBadRequest).JSON(utils.Error(apiError.AppError{
			Message: "request failed",
			Error:   valErr,
		}))
	}

	err = c.as.Authenticate(ctx.UserContext(), body)
	if err != nil {
		logger.Error("an apiError occurred ", err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(utils.Error(apiError.AppError{
			Message: "request failed",
			Error:   err.Error(),
		}))
	}

	return ctx.Status(http.StatusOK).JSON(utils.Success("successful", nil, nil))
}

func (c *AuthController) RefreshToken(ctx *fiber.Ctx) error {
	logger := log.WithFields(log.Fields{constant.RequestIdentifier: utils.GenerateUUID()})
	logger.Info("RefreshToken")

	body := types.RefreshToken{}

	err := ctx.BodyParser(body)
	if err != nil {
		logger.Error(err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(utils.Error(apiError.AppError{
			Message: "request failed",
			Error:   err.Error(),
		}))
	}

	if valErr, ok := validation.ValidateStruct(body, logger); !ok {
		return ctx.Status(http.StatusBadRequest).JSON(utils.Error(apiError.AppError{
			Message: "request failed",
			Error:   valErr,
		}))
	}

	data, err := c.as.RefreshToken(ctx.UserContext(), &body)
	if err != nil {
		logger.Error("an apiError occurred ", err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(utils.Error(apiError.AppError{
			Message: "request failed",
			Error:   err.Error(),
		}))
	}

	return ctx.Status(http.StatusOK).JSON(utils.Success("successful", data, nil))
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	logger := log.WithFields(log.Fields{constant.RequestIdentifier: utils.GenerateUUID()})
	logger.Info("Login")

	body := types.Authenticate{}

	err := ctx.BodyParser(body)
	if err != nil {
		logger.Error(err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(utils.Error(apiError.AppError{
			Message: "request failed",
			Error:   err.Error(),
		}))
	}

	if valErr, ok := validation.ValidateStruct(body, logger); !ok {
		return ctx.Status(http.StatusBadRequest).JSON(utils.Error(apiError.AppError{
			Message: "request failed",
			Error:   valErr,
		}))
	}

	data, err := c.as.Login(ctx.UserContext(), body)
	if err != nil {
		logger.Error("an apiError occurred ", err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(utils.Error(apiError.AppError{
			Message: "request failed",
			Error:   err.Error(),
		}))
	}

	return ctx.Status(http.StatusOK).JSON(utils.Success("successful", data, nil))
}
