package controllers

import (
	"encoding/json"
	"github.com/TechBuilder-360/business-directory-backend/configs"
	"github.com/TechBuilder-360/business-directory-backend/repository"
	"github.com/TechBuilder-360/business-directory-backend/services"
	"github.com/TechBuilder-360/business-directory-backend/utility"
	"net/http"

	logger "github.com/Toflex/oris_log"
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
	AuthenticateEmail(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

type NewController struct {
	Service    services.Service
	JWTService services.JWTService
	Repo       repository.Repository
	Logger     logger.Logger
	Config     *configs.Config
}

func DefaultController(serv services.Service, auth services.JWTService, log logger.Logger, repo repository.Repository, config *configs.Config) Controller {
	return &NewController{
		Service:    serv,
		JWTService: auth,
		Logger:     log,
		Repo:       repo,
		Config:     config,
	}
}

func (c *NewController) Ping(w http.ResponseWriter, r *http.Request) {
	log := c.Logger.NewContext()
	log.SetLogID(r.Header.Get("LogID"))
	apiResponse := utility.NewResponse()

	json.NewEncoder(w).Encode(apiResponse.PlainSuccess(utility.SYSTEM001))
	return
}
