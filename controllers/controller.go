package controllers

import (
	"encoding/json"
	"github.com/TechBuilder-360/business-directory-backend.git/services"
	"github.com/Toflex/oris_log/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	Service services.Service
	Logger logger.Logger
}

//func NewController(serv *services.DefaultService, logger logger.Logger) *Controller {
//	return &Controller{Service: serv, Logger: logger}
//}

func (c *Controller) Ping(ct *gin.Context) {
	var body interface{}
	err:=json.NewDecoder(ct.Request.Body).Decode(&body)
	if err!= nil{
		c.Logger.Error(err.Error())
	}
	c.Logger.Info("%+v", body)
	ct.JSON(http.StatusOK, "Pongc ...")
}