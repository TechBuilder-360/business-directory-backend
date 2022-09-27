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
	"strconv"
)

type BranchController interface {
	RegisterRoutes(router *mux.Router)
	GetSingleBranch(w http.ResponseWriter, r *http.Request)
	GetAllBranch(w http.ResponseWriter, r *http.Request)
}

type NewBranchController struct {
	Service services.BranchService
}

func (c *NewBranchController) RegisterRoutes(router *mux.Router) {
	apis := router.PathPrefix("/branches").Subrouter()
	apis.HandleFunc("/{id}", middlewares.Adapt(http.HandlerFunc(c.GetSingleBranch), middlewares.AuthorizeUserJWT()).ServeHTTP).Methods(http.MethodGet)
	apis.HandleFunc("/fetch-all-branch", middlewares.Adapt(http.HandlerFunc(c.GetAllBranch), middlewares.AuthorizeUserJWT()).ServeHTTP).Methods(http.MethodGet)

}

func DefaultBranchController() BranchController {
	return &NewBranchController{
		Service: services.NewBranchService(),
	}
}

// GetSingleBranch godoc
// @Summary      fetch an branch
// @Description  fetch an branch
// @Tags         Get One
// @Accept       json
// @Produce      json
// @Param        default  body	id  true  "fetch an branch"
// @Success      200      {object}  utils.SuccessResponse{Data=data}
// @Router       /branches/{id} [get]
func (c *NewBranchController) GetSingleBranch(w http.ResponseWriter, r *http.Request) {
	logger := log.WithFields(log.Fields{constant.RequestIdentifier: utils.GenerateUUID()})
	logger.Info("fetching single branch.")
	params := mux.Vars(r)
	OrganisationID := params["id"]

	data, err := c.Service.GetSingleBranch(OrganisationID, logger)
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

// GetAllBranch godoc
// @Summary      fetch all branch
// @Description  fetch all brancn
// @Tags         Get all
// @Accept       json
// @Produce      json
// @Param        default  body	id  true  "fetch all branch"
// @Success      200      {object}  utils.SuccessResponse{Data=data}
// @Router       /branches/fetch-all-branch [get]
func (c *NewBranchController) GetAllBranch(w http.ResponseWriter, r *http.Request) {
	logger := log.WithFields(log.Fields{constant.RequestIdentifier: utils.GenerateUUID()})
	logger.Info("fetching All Branch.")
	params := mux.Vars(r)
	BranchID, err := strconv.Atoi(params["page"])
	if err != nil {
		logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	data, err := c.Service.GetAllBranch(BranchID, logger)
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
