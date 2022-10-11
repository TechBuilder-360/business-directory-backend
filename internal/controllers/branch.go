package controllers

import (
	"encoding/json"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/constant"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/utils"
	"github.com/TechBuilder-360/business-directory-backend/internal/middlewares"
	"github.com/TechBuilder-360/business-directory-backend/internal/services"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type BranchController interface {
	RegisterRoutes(router *mux.Router)
	GetBranches(w http.ResponseWriter, r *http.Request)
}

type NewBranchController struct {
	Service services.BranchService
}

func (c *NewBranchController) RegisterRoutes(router *mux.Router) {
	apis := router.PathPrefix("/branches").Subrouter()

	apis.HandleFunc("", middlewares.Adapt(http.HandlerFunc(c.GetBranches), middlewares.AuthorizeUserJWT()).ServeHTTP).Methods(http.MethodGet)
}

func DefaultBranchController() BranchController {
	return &NewBranchController{
		Service: services.NewBranchService(),
	}
}

// GetBranches godoc
// @Summary      get branches
// @Description  Get branches
// @Tags         Branch
// @Accept       json
// @Produce      json
// @Success      200      {object}  utils.SuccessResponse{Data=[]types.Branch}
// @Router       /branches [get]
func (c *NewBranchController) GetBranches(w http.ResponseWriter, r *http.Request) {
	logger := log.WithFields(log.Fields{constant.RequestIdentifier: utils.GenerateUUID()})
	logger.Info("fetching single branch.")

	// get organisation from context
	organisation, err := middlewares.OrganisationFromContext(r)
	if err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.ErrorResponse{
			Status:  false,
			Message: "organisation not found",
		})
		return
	}

	data, err := c.Service.GetOrganisationBranches(organisation, logger)
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
