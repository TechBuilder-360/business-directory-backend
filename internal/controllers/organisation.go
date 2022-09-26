package controllers

import (
	"encoding/json"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/constant"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/types"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/utils"
	"github.com/TechBuilder-360/business-directory-backend/internal/middlewares"
	"github.com/TechBuilder-360/business-directory-backend/internal/services"
	"github.com/TechBuilder-360/business-directory-backend/internal/validation"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type OrganisationController interface {
	CreateOrganisation(w http.ResponseWriter, r *http.Request)
	ActivateOrganisation(w http.ResponseWriter, r *http.Request)
	DeactivateOrganisation(w http.ResponseWriter, r *http.Request)
	RegisterRoutes(router *mux.Router)
}

type organisationController struct {
	Service services.OrganisationService
}

func (c *organisationController) RegisterRoutes(router *mux.Router) {
	apis := router.PathPrefix("/organisations").Subrouter()
	apis.HandleFunc("/create", middlewares.Adapt(http.HandlerFunc(c.CreateOrganisation), middlewares.AuthorizeUserJWT()).ServeHTTP).Methods(http.MethodPost)
	apis.HandleFunc("/activate-organisation/{id}", middlewares.Adapt(http.HandlerFunc(c.ActivateOrganisation), middlewares.AuthorizeUserJWT()).ServeHTTP).Methods(http.MethodGet)
	apis.HandleFunc("/de-activate-organisation/{id}", middlewares.Adapt(http.HandlerFunc(c.DeactivateOrganisation), middlewares.AuthorizeUserJWT()).ServeHTTP).Methods(http.MethodGet)

}

func DefaultOrganisationController() OrganisationController {
	return &organisationController{
		Service: services.NewOrganisationService(),
	}
}

// CreateOrganisation godoc
// @Summary      create an organisation
// @Description  create an organisation
// @Tags         Create
// @Accept       json
// @Produce      json
// @Param        default  body	types.CreateOrganisationReq  true  "create this organisation"
// @Success      200      {object}  utils.SuccessResponse{Data=types.CreateOrganisationResponse}
// @Router       /organisations/create [post]
func (c *organisationController) CreateOrganisation(w http.ResponseWriter, r *http.Request) {
	logger := log.WithFields(log.Fields{constant.RequestIdentifier: utils.GenerateUUID()})
	logger.Info("Creating Organisation.")

	body := &types.CreateOrganisationReq{}

	err := json.NewDecoder(r.Body).Decode(body)
	if err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.ErrorResponse{
			Status:  false,
			Message: "bad request",
		})
		return
	}

	if validation.ValidateStruct(w, body, logger) {
		return
	}

	// get user from context
	user, err := middlewares.UserFromContext(r)
	if err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	data, err := c.Service.CreateOrganisation(body, user, logger)
	if err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.SuccessResponse{
		Status:  true,
		Message: "Successful",
		Data:    data,
	})

}

// CreateOrganisation godoc
// @Summary      create an organisation
// @Description  create an organisation
// @Tags         Create
// @Accept       json
// @Produce      json
// @Param        default  body	types.CreateOrganisationReq  true  "create this organisation"
// @Success      200      {object}  utils.SuccessResponse{Data=types.CreateOrganisationResponse}
// @Router       /organisations/create [post]
func (c *organisationController) ActivateOrganisation(w http.ResponseWriter, r *http.Request) {
	logger := log.WithFields(log.Fields{constant.RequestIdentifier: utils.GenerateUUID()})
	logger.Info("Activating Organisation.")
	params := mux.Vars(r)
	OrganisationID := params["id"]

	err := c.Service.ActivateOrganisation(OrganisationID, logger)
	if err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.SuccessResponse{
		Status:  true,
		Message: "Successful",
		Data:    nil,
	})

}

// CreateOrganisation godoc
// @Summary      create an organisation
// @Description  create an organisation
// @Tags         Create
// @Accept       json
// @Produce      json
// @Param        default  body	types.CreateOrganisationReq  true  "create this organisation"
// @Success      200      {object}  utils.SuccessResponse{Data=types.CreateOrganisationResponse}
// @Router       /organisations/create [post]
func (c *organisationController) DeactivateOrganisation(w http.ResponseWriter, r *http.Request) {
	logger := log.WithFields(log.Fields{constant.RequestIdentifier: utils.GenerateUUID()})
	logger.Info("De-activating Organisation.")
	params := mux.Vars(r)
	OrganisationID := params["id"]

	err := c.Service.DeactivateOrganisation(OrganisationID, logger)
	if err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.SuccessResponse{
		Status:  true,
		Message: "Successful",
		Data:    nil,
	})

}