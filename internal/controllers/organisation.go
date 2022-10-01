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

type IOrganisationController interface {
	CreateOrganisation(w http.ResponseWriter, r *http.Request)
	GetOrganisation(w http.ResponseWriter, r *http.Request)
	RegisterRoutes(router *mux.Router)
}

type organisationController struct {
	Service services.IOrganisationService
}

func (c *organisationController) RegisterRoutes(router *mux.Router) {
	apis := router.PathPrefix("/organisations").Subrouter()

	apis.HandleFunc("/", middlewares.Adapt(http.HandlerFunc(c.CreateOrganisation), middlewares.AuthorizeUserJWT()).ServeHTTP).Methods(http.MethodPost)
	apis.HandleFunc("/{id}", c.GetOrganisation).Methods(http.MethodGet)
}

func DefaultOrganisationController() IOrganisationController {
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
// @Success      201      {object}  utils.SuccessResponse{Data=types.CreateOrganisationResponse}
// @Router       /organisations [post]
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

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(utils.SuccessResponse{
		Status:  true,
		Message: "Successful",
		Data:    data,
	})

}

// GetOrganisation godoc
// @Summary      get organisation
// @Description  get organisation
// @Tags         Organisation
// @Accept       json
// @Produce      json
// @Param        default  path	string  true  "organisation ID"
// @Success      200      {object}  utils.SuccessResponse{types.Organisation}
// @Router       /organisation/{id} [get]
func (c *organisationController) GetOrganisation(w http.ResponseWriter, r *http.Request) {
	logger := log.WithFields(log.Fields{constant.RequestIdentifier: utils.GenerateUUID()})
	logger.Info("GetOrganisation")

	vars := mux.Vars(r)
	id := vars["id"]

	data, err := c.Service.GetOrganisation(id)
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

// GetOrganisations godoc
// @Summary      get organisation
// @Description  get organisation
// @Tags         Organisation
// @Accept       json
// @Produce      json
// @Success      200      {object}  utils.SuccessResponse{Data=types.Organisations}
// @Router       /organisations [get]
//func (c *organisationController) GetOrganisations(w http.ResponseWriter, r *http.Request) {
//	logger := log.WithFields(log.Fields{constant.RequestIdentifier: utils.GenerateUUID()})
//	logger.Info("GetOrganisation")
//
//	vars := mux.Vars(r)
//	id := vars["id"]
//
//	data, err := c.Service.GetOrganisations(id)
//	if err != nil {
//		logger.Error(err.Error())
//		w.WriteHeader(http.StatusBadRequest)
//		json.NewEncoder(w).Encode(utils.ErrorResponse{
//			Status:  false,
//			Message: err.Error(),
//		})
//		return
//	}
//
//	w.WriteHeader(http.StatusOK)
//	json.NewEncoder(w).Encode(utils.SuccessResponse{
//		Status:  true,
//		Message: "Successful",
//		Data:    data,
//	})
//
//}
