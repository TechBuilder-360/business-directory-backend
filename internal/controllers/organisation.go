package controllers

import (
	"encoding/json"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/consts"
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

type NewOrganisationController struct {
	Service services.OrganisationService
}

func (c *NewOrganisationController) RegisterRoutes(router *mux.Router) {
	_ = router.PathPrefix("/organisations").Subrouter()
}

func DefaultOrganisationController() OrganisationController {
	return &NewOrganisationController{
		Service: services.NewOrganisationService(),
	}
}

// CreateOrganisation @Summary      Add an organisation
// @Description  add by json organisation
// @Tags         organisation
// @Accept       json
// @Produce      json
// @Param        default  body	types.CreateOrgReq  true  "Add organisation"
// @Success      200      {object}  utils.ResponseObj
// @Router       /organisation [post]
func (c *NewOrganisationController) CreateOrganisation(w http.ResponseWriter, r *http.Request) {
	logger := log.WithFields(log.Fields{consts.RequestIdentifier: utils.GenerateUUID()})
	logger.Info("creating organisation")

	requestData := &types.CreateOrgReq{}
	response := &types.Organisation{}
	_ = json.NewDecoder(r.Body).Decode(&requestData)

	if validation.ValidateStruct(w, requestData, logger) {
		return
	}

	response, err := c.Service.CreateOrganisation(requestData, nil, logger)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	logger.Info("Response body: %+v", response)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.SuccessResponse{
		Status:  true,
		Message: "Successful",
		Data:    response,
	})
}

// GetOrganisations @Summary      Get list of organisation
// @Description  Get the list of organisation with pagination
// @Tags         organisation
// @Param        page  query	int    false  "int valid"
// @Success      200      {array}  types.DataView
// @Router       /get-organisations [get]
func (c *NewOrganisationController) GetOrganisations(w http.ResponseWriter, r *http.Request) {
	logger := log.WithFields(log.Fields{consts.RequestIdentifier: utils.GenerateUUID()})
	logger.Info("getting organisation")

	//apiResponse := utils.NewResponse()
	//query := r.URL.Query()
	//pageNumber := query.Get("page")
	//page := 1
	//
	//if pageNumber != "" {
	//	pNumber, err := strconv.Atoi(pageNumber)
	//	if err == nil {
	//		page = pNumber
	//	}
	//}

	//organnisations, orgErr := services.GetOrganisations(page, c.Repo, logger)
	//if orgErr != nil {
	//	w.WriteHeader(orgErr.StatusCode)
	//	json.NewEncoder(w).Encode(apiResponse.Error(orgErr.ResponseCode))
	//	return
	//}

	//json.NewEncoder(w).Encode(apiResponse.Success(utils.SYSTEM001, nil))
}

// GetSingleOrganisation @Summary      Get organisation
// @Description  Get a single organisation
// @Tags         organisation
// @Param        organisationId  path	string    false  "string valid"
// @Success      200      {object}  model.Organisation
// @Router       /organisation/{organisationId} [get]
func (c *NewOrganisationController) GetSingleOrganisation(w http.ResponseWriter, r *http.Request) {
	logger := log.WithFields(log.Fields{consts.RequestIdentifier: utils.GenerateUUID()})
	logger.Info("getting single organisation")
	//apiResponse := utils.NewResponse()
	//params := mux.Vars(r)
	//organisationID := params["organisation_id"]

	//organ, err := services.GetSingleOrganisation(organisationID, c.Repo, logger)
	//if err != nil {
	//	w.WriteHeader(err.StatusCode)
	//	json.NewEncoder(w).Encode(apiResponse.Error(err.ResponseCode))
	//	return
	//}

	w.WriteHeader(http.StatusOK)
	//json.NewEncoder(w).Encode(apiResponse.Success(utils.SYSTEM001, nil))
}

// GetBranches @Summary      Get list of branches
// @Description  Get the list of branches with pagination
// @Tags         branch
// @Param        page  query	int    false  "int valid"
// @Param        organisationId path	string    false  "string valid"
// @Success      200      {array}  model.Branch
// @Router       /branches/{organisationId} [get]
func (c *NewOrganisationController) GetBranches(w http.ResponseWriter, r *http.Request) {
	logger := log.WithFields(log.Fields{consts.RequestIdentifier: utils.GenerateUUID()})
	logger.Info("getting organisation")
	//apiResponse := utils.NewResponse()
	//query := r.URL.Query()
	//page := query.Get("page")
	//if page == "" {
	//	page = "1"
	//}
	//params := mux.Vars(r)
	//organisationID := params["organisationId"]
	//
	//branch, err := services.GetBranches(organisationID, page, c.Repo, logger)
	//if err != nil {
	//	w.WriteHeader(err.StatusCode)
	//	json.NewEncoder(w).Encode(apiResponse.Error(err.ResponseCode))
	//	return
	//}

	//json.NewEncoder(w).Encode(apiResponse.Success(utils.SYSTEM001, nil))
}

// GetSingleBranch @Summary      Get branch
// @Description  Get a single branch
// @Tags         branch
// @Param        branchId  path	string    false  "string valid"
// @Success      200      {object}  model.Organisation
// @Router       /branch/{branchId} [get]
func (c *NewOrganisationController) GetSingleBranch(w http.ResponseWriter, r *http.Request) {
	logger := log.WithFields(log.Fields{consts.RequestIdentifier: utils.GenerateUUID()})
	logger.Info("getting single branch")
	//apiResponse := utils.NewResponse()
	//params := mux.Vars(r)
	//branchID := params["branchId"]
	//
	//branch, err := services.GetSingleBranch(branchID, c.Repo, log)
	//if err != nil {
	//	w.WriteHeader(err.StatusCode)
	//	json.NewEncoder(w).Encode(apiResponse.Error(err.ResponseCode))
	//	return
	//}

	//json.NewEncoder(w).Encode(apiResponse.Success(utils.SYSTEM001, nil))
}

// OrganisationStatus @Summary      Setting the Status for an  Organisation
// @Description  Setting the Status for an  Organisation Operation
// @Tags         organisation
// @Accept       json
// @Produce      json
// @Param        default  body	types.OrganStatus  true  "activate or deactivate"
// @Success      200      {object}  utils.Response
// @Router       /active-status [post]
func (c *NewOrganisationController) OrganisationStatus(w http.ResponseWriter, r *http.Request) {
	logger := log.WithFields(log.Fields{consts.RequestIdentifier: utils.GenerateUUID()})
	logger.Info("deactivating and activating organisation")
	//apiResponse := utils.NewResponse()
	//var request *types.OrganStatus
	//err := json.NewDecoder(r.Body).Decode(&request)
	//if err != nil {
	//	log.Error("Error occured while parsing the request body, %s", err.Error())
	//	w.WriteHeader(http.StatusNotAcceptable)
	//	json.NewEncoder(w).Encode(apiResponse.Error(utils.BAD_REQUEST))
	//	return
	//}
	//
	////_, statusErr := services.OrganisationStatus(request, c.Repo, log)
	////if err != nil {
	////	w.WriteHeader(statusErr.StatusCode)
	////	json.NewEncoder(w).Encode(apiResponse.Error(statusErr.ResponseCode))
	////	return
	////}
	//
	//json.NewEncoder(w).Encode(apiResponse.PlainSuccess(utils.SYSTEM001))

}
