package controllers

import (
	"encoding/json"
	"github.com/TechBuilder-360/business-directory-backend/common/consts"
	"github.com/TechBuilder-360/business-directory-backend/common/types"
	"github.com/TechBuilder-360/business-directory-backend/services"
	"github.com/TechBuilder-360/business-directory-backend/utility"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type BranchController interface {
	CreateBranch(w http.ResponseWriter, r *http.Request)
	RegisterRoutes(router *mux.Router)
}

type NewBranchController struct {
	Service services.BranchService
}

func (c *NewBranchController) RegisterRoutes(router *mux.Router) {
	_ = router.PathPrefix("/branches").Subrouter()
}

func DefaultBranchController() BranchController {
	return &NewBranchController{
		Service: services.NewBranchService(),
	}
}

// CreateBranch @Summary      Add branch
// @Description  add branch to an organisation
// @Tags         branch
// @Accept       json
// @Produce      json
// @Param        default  body	types.CreateOrgReq  true  "Add branch"
// @Success      200      {object}  utility.ResponseObj
// @Router       /organisation [post]
func (c *NewBranchController) CreateBranch(w http.ResponseWriter, r *http.Request) {
	logger := log.WithFields(log.Fields{consts.RequestIdentifier: utility.GenerateUUID()})

	requestData := &types.CreateOrgReq{}
	_ = &types.Organisation{}
	_ = json.NewDecoder(r.Body).Decode(&requestData)

	if utility.ValidateStruct(w, requestData, logger) {
		return
	}

	//response, err := nil, nil //c.Service.CreateOrganisation(requestData, nil, logger)
	//if err != nil {
	//	w.WriteHeader(http.StatusBadRequest)
	//	json.NewEncoder(w).Encode(apiResponse.Error(err.Error()))
	//	return
	//}

	//logger.Info("Response body: %+v", response)
	//w.WriteHeader(http.StatusOK)
	//json.NewEncoder(w).Encode(apiResponse.Success(utility.SYSTEM001, response))
}
