package controllers

import (
	"encoding/json"
	"github.com/TechBuilder-360/business-directory-backend.git/services"
	logger "github.com/Toflex/oris_log"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller interface {
	Ping(ctx *gin.Context)
}

type NewController struct {
	Service services.Service
	Logger  logger.Logger
}

func DefaultController(serv services.Service, log logger.Logger) Controller {
	return &NewController{
		Service: serv,
		Logger: log,
	}
}

func (c *NewController) Ping(ct *gin.Context) {
	var body interface{}
	err := json.NewDecoder(ct.Request.Body).Decode(&body)
	if err != nil {
		c.Logger.Error(err.Error())
	}
	c.Logger.Info("%+v", body)
	ct.JSON(http.StatusOK, "Pong ...")
}
