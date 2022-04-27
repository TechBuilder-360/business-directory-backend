package controllers

import (
	"github.com/TechBuilder-360/business-directory-backend/configs"
	"github.com/TechBuilder-360/business-directory-backend/services"
	logger "github.com/Toflex/oris_log"
)

type UserController interface {
}

type NewUserController struct {
	Service  services.UserService
	Logger     logger.Logger
	Config     *configs.Config
}

func DefaultUserController(serv services.UserService, log logger.Logger, config *configs.Config) UserController {
	return &NewUserController{
		Service:    serv,
		Logger:     log,
		Config:     config,
	}
}
