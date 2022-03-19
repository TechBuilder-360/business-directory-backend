package controllers

import (
	//"context"
	"encoding/json"

	"net/http"

	
	"github.com/TechBuilder-360/business-directory-backend/repository"
	"github.com/TechBuilder-360/business-directory-backend/services"
	
	logger "github.com/Toflex/oris_log"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type Controller interface {
	Ping(w http.ResponseWriter, r *http.Request)
	CreateOrganisation(w http.ResponseWriter, r *http.Request)
}
type customClaims struct {
	Username string `json:"username"`
	Role     string ` json:"role"`
	jwt.StandardClaims
}
type NewController struct {
	Service services.Service
	Logger  logger.Logger
	Repo   repository.Repository
}

func DefaultController(serv services.Service, log logger.Logger,repo repository.Repository) Controller {
	return &NewController{
		Service: serv,
		Logger:  log,
		Repo : repo,
	}
}

func (c *NewController) Ping(w http.ResponseWriter, r *http.Request) {
	log:= c.Logger.NewContext()
	log.SetLogID(r.Header.Get("LogID"))
	log.Debug("Ping")
	r.Header.Set("TraceID", uuid.New().String())

	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Hi"})
	return
}

