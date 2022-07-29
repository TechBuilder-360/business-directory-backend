package controllers

import (
	"encoding/json"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/consts"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/utils"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Controller interface {
	Ping(w http.ResponseWriter, r *http.Request)
	RegisterRoutes(router *mux.Router)
}

func (c *NewController) RegisterRoutes(router *mux.Router) {
	api := router.PathPrefix("").Subrouter()

	api.HandleFunc("/ping", c.Ping)
}

type NewController struct {
}

func DefaultController() Controller {
	return &NewController{}
}

func (c *NewController) Ping(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{consts.RequestIdentifier: utils.GenerateUUID()})

	json.NewEncoder(w).Encode(utils.SuccessResponse{
		Status:  true,
		Message: "We are up and running 🚀",
	})
	return
}
