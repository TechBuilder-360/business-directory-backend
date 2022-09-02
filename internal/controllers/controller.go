package controllers

import (
	"encoding/json"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/consts"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/utils"
	"github.com/TechBuilder-360/business-directory-backend/middlewares"
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

	api.HandleFunc("/ping", middlewares.Chain(c.Ping, middlewares.Method("GET"), middlewares.Logging())) //.Methods(http.MethodGet)
}

type NewController struct {
}

func DefaultController() Controller {
	return &NewController{}
}

func (c *NewController) Ping(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{consts.RequestIdentifier: utils.GenerateUUID()})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(utils.SuccessResponse{
		Status:  true,
		Message: "We are up and running 🚀",
	})
	return
}
