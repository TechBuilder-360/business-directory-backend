package controllers

import (
	"encoding/json"
	"github.com/TechBuilder-360/business-directory-backend.git/configs"
	"github.com/TechBuilder-360/business-directory-backend.git/repository"
	"github.com/TechBuilder-360/business-directory-backend.git/services"
	log "github.com/Toflex/oris_log"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	Service services.Service
	Repository repository.Repository
	Logger log.Logger
	Config *configs.Config
}


func (c *Controller) Ping(ct *gin.Context) {
	var body interface{}
	err:=json.NewDecoder(ct.Request.Body).Decode(&body)
	if err!= nil{
		c.Logger.Error(err.Error())
	}
	c.Logger.Info("%+v", body)
	ct.JSON(http.StatusOK, "Pong ...")
}