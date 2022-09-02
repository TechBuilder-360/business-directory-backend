package controllers

import (
	"encoding/json"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/constant"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/types"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/utils"
	"github.com/TechBuilder-360/business-directory-backend/internal/services"
	"github.com/TechBuilder-360/business-directory-backend/internal/validation"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type OrganisationController interface {
	CreateOrganisation(w http.ResponseWriter, r *http.Request)
	RegisterRoutes(router *mux.Router)
}

type organisationController struct {
	Service services.OrganisationService
}

func (c *organisationController) RegisterRoutes(router *mux.Router) {
	apis := router.PathPrefix("/organisations").Subrouter()
	apis.HandleFunc("/create", c.CreateOrganisation).Methods(http.MethodPost)
}

func DefaultOrganisationController() OrganisationController {
	return &organisationController{
		Service: services.NewOrganisationService(),
	}
}

// Create Organisation
// @Summary      create an organisation
// @Description  create an organisation
// @Tags         Create
// @Accept       json
// @Produce      json
// @Param        default  body	types.CreateOrgReq  true  "create this organisation"
// @Success      200      {object}  utils.SuccessResponse
// @Router       /organisations/create [post]
func (c *organisationController) CreateOrganisation(w http.ResponseWriter, r *http.Request) {
	logger := log.WithFields(log.Fields{constant.RequestIdentifier: utils.GenerateUUID()})
	logger.Info("Creating Organisation.")

	body := &types.CreateOrgReq{}

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

	data, err := c.Service.CreateOrganisation(body, logger)
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
		Message: "Successful created",
		Data:    data,
	})

}
