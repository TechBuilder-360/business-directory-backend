package controllers

import (
	"github.com/TechBuilder-360/business-directory-backend.git/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	Service services.Service
	//Logger logger.Logger
}

//func NewController(serv *services.DefaultService, logger logger.Logger) *Controller {
//	return &Controller{Service: serv, Logger: logger}
//}

func (c *Controller) Ping(ct *gin.Context) {
	//c.Logger.Info("Pinging...")
	author:= c.Service.GetAuthor()
	ct.JSON(http.StatusOK, author)
}