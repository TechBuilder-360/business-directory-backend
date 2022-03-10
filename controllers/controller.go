package controllers

import (
	"github.com/google/uuid"
	"net/http"

	"github.com/TechBuilder-360/business-directory-backend/services"
	logger "github.com/Toflex/oris_log"
	"github.com/gin-gonic/gin"
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
	log:= c.Logger.NewContext()
	log.SetLogID(ct.MustGet("LogID").(string))
	log.Debug("Ping")
	ct.Header("TraceID", uuid.New().String())
	ct.JSON(http.StatusOK, map[string]interface{}{"message": "Hi"})
	return
}
