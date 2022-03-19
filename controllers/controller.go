package controllers

import (
	"encoding/json"
	"github.com/TechBuilder-360/business-directory-backend/repository"
	"github.com/TechBuilder-360/business-directory-backend/utility"
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
	Repo    repository.Repository
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
	apiResponse := utility.NewResponse()
	log.Debug("Ping")
	r.Header.Set("TraceID", uuid.New().String())

	d:=c.Service.GetAuthor(log)
	//_ := c.Repo.GetClientByID("sdd")

	//json.NewEncoder(w).Encode(map[string]interface{}{"message": "Hi"})
	json.NewEncoder(w).Encode(apiResponse.Success(utility.SYSTEM001, utility.GetCodeMsg(utility.SYSTEM001), d))
	return
}

