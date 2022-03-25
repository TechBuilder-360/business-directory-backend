package controllers

import (
	"encoding/json"
	"github.com/TechBuilder-360/business-directory-backend/repository"
	"github.com/TechBuilder-360/business-directory-backend/services"
	"github.com/TechBuilder-360/business-directory-backend/utility"
	"net/http"

	logger "github.com/Toflex/oris_log"
	"github.com/dgrijalva/jwt-go"
)

type Controller interface {
	Ping(w http.ResponseWriter, r *http.Request)
	CreateOrganisation(w http.ResponseWriter, r *http.Request)
	CreateBranch(w http.ResponseWriter, r *http.Request)
	GetOrganisations(w http.ResponseWriter, r *http.Request)
	GetSingleOrganisation(w http.ResponseWriter, r *http.Request)
	GetBranches(w http.ResponseWriter, r *http.Request)
	GetSingleBranch(w http.ResponseWriter, r *http.Request)
	OrganisationStatus(w http.ResponseWriter, r *http.Request)
	RegisterUser(w http.ResponseWriter, r *http.Request)
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

func DefaultController(serv services.Service, log logger.Logger, repo repository.Repository) Controller {
	return &NewController{
		Service: serv,
		Logger:  log,
		Repo:    repo,
	}
}

func (c *NewController) Ping(w http.ResponseWriter, r *http.Request) {
	log := c.Logger.NewContext()
	log.SetLogID(r.Header.Get("LogID"))
	apiResponse := utility.NewResponse()

	json.NewEncoder(w).Encode(apiResponse.PlainSuccess(utility.SYSTEM001, utility.GetCodeMsg(utility.SYSTEM001)))
	return
}
