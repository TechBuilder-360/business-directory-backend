package controllers

import (
	"encoding/json"
	"github.com/TechBuilder-360/business-directory-backend/utility"
	"github.com/google/uuid"
	"net/http"

	logger "github.com/Toflex/oris_log"
)

type Controller interface {
	Ping(w http.ResponseWriter, r *http.Request)
}

type NewController struct {
	Logger     logger.Logger
}

func DefaultController(log logger.Logger) Controller {
	return &NewController{
		Logger:     log,
	}
}

func (c *NewController) Ping(w http.ResponseWriter, r *http.Request) {
	log := c.Logger.NewContext()
	log.SetLogID(uuid.NewString())
	apiResponse := utility.NewResponse()

	json.NewEncoder(w).Encode(apiResponse.PlainSuccess("We are up and running 🚀"))
	return
}
