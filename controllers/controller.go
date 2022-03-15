package controllers

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"

	"github.com/TechBuilder-360/business-directory-backend/services"
	logger "github.com/Toflex/oris_log"
	"github.com/dgrijalva/jwt-go"
)

type Controller interface {
	Ping(w http.ResponseWriter, r *http.Request)
}
type customClaims struct {
	Username string `json:"username"`
	Role     string ` json:"role"`
	jwt.StandardClaims
}
type NewController struct {
	Service services.Service
	Logger  logger.Logger
}

func DefaultController(serv services.Service, log logger.Logger) Controller {
	return &NewController{
		Service: serv,
		Logger:  log,
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
