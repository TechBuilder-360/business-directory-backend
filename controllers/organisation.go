package controllers

import (
	"encoding/json"
	"github.com/TechBuilder-360/business-directory-backend/dto"
	"github.com/TechBuilder-360/business-directory-backend/services"
	"github.com/TechBuilder-360/business-directory-backend/utility"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// CreateOrganisation @Summary      Add an organisation
// @Description  add by json organisation
// @Tags         organisation
// @Accept       json
// @Produce      json
// @Param        default  body	dto.CreateOrgReq  true  "Add add organisation"
// @Success      200      {object}  utility.ResponseObj
// @Router       /organisation [post]
func (c *NewController) CreateOrganisation(w http.ResponseWriter, r *http.Request) {
	log := c.Logger.NewContext()
	log.SetLogID(r.Header.Get("LogID"))
	log.Info("creating organisation")

	apiResponse := utility.NewResponse()
	requestData := &dto.CreateOrgReq{}
	response := &dto.CreateOrgResponse{}
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		log.Error("Error occurred while parsing the request body, %s", err.Error())
		json.NewEncoder(w).Encode(apiResponse.Error(utility.BAD_REQUEST))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Debug("Request body: %+v", requestData)

	validationRes := validator.New()
	if validateErr := validationRes.Struct(requestData); validateErr != nil {
		validationErrors := validateErr.(validator.ValidationErrors)
		log.Error("Validation failed on some fields : %+v", validationErrors)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(apiResponse.ValidationError(utility.VALIDATIONERR, validationErrors.Error()))
		return
	}

	response, serviceErr := services.CreateOrganisation(requestData, c.Repo, log)
	if serviceErr != nil {
		w.WriteHeader(serviceErr.StatusCode)
		json.NewEncoder(w).Encode(apiResponse.Error(serviceErr.ResponseCode))
		return
	}

	log.Info("Response body: %+v", response)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(apiResponse.Success(utility.SYSTEM001, response))
}

// CreateBranch @Summary   Add a Branch
// @Description  add by json Branch
// @Tags         branch
// @Accept       json
// @Produce      json
// @Param        default  body	dto.CreateBranch  true  "Add branch"
// @Success      201      {object}  utility.Response
// @Router       /branch [post]
func (c *NewController) CreateBranch(w http.ResponseWriter, r *http.Request) {
	log := c.Logger.NewContext()
	log.SetLogID(r.Header.Get("LogID"))
	log.Info("creating branch")
	apiResponse := utility.NewResponse()
	var request *dto.CreateBranch
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Error("Error occurred while parsing the request body, %s", err.Error())
		json.NewEncoder(w).Encode(apiResponse.Error(utility.BAD_REQUEST))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Debug("Request body: %+v", &request)

	validationRes := validator.New()
	if err := validationRes.Struct(&request); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		log.Error("Validation failed on some fields : %+v", validationErrors)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(apiResponse.ValidationError(utility.VALIDATIONERR, validationErrors.Error()))
		return
	}

	if serviceErr:=services.CreateBranch(request, c.Repo, log); serviceErr != nil {
		w.WriteHeader(serviceErr.StatusCode)
		json.NewEncoder(w).Encode(apiResponse.Error(serviceErr.ResponseCode))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(apiResponse.PlainSuccess(utility.SYSTEM001))
}

// GetOrganisations @Summary      Get list of organisation
// @Description  Get the list of organisation with pagination
// @Tags         organisation
// @Param        page  query	int    false  "int valid"
// @Success      200      {array}  dto.DataView
// @Router       /get-organisations [get]
func (c *NewController) GetOrganisations(w http.ResponseWriter, r *http.Request) {
	log := c.Logger.NewContext()
	log.SetLogID(r.Header.Get("LogID"))
	log.Info("getting organisation")

	apiResponse := utility.NewResponse()
	query := r.URL.Query()
	pageNumber := query.Get("page")
	page := 1

	if pageNumber != "" {
		pNumber, err := strconv.Atoi(pageNumber)
		if err == nil {
			page = pNumber
		}
	}

	organnisations, orgErr:=services.GetOrganisations(page, c.Repo, log)
	if orgErr != nil {
		w.WriteHeader(orgErr.StatusCode)
		json.NewEncoder(w).Encode(apiResponse.Error(orgErr.ResponseCode))
		return
	}

	json.NewEncoder(w).Encode(apiResponse.Success(utility.SYSTEM001, organnisations))
}

// GetSingleOrganisation @Summary      Get organisation
// @Description  Get a single organisation
// @Tags         organisation
// @Param        organisationId  path	string    false  "string valid"
// @Success      200      {object}  models.Organisation
// @Router       /organisation/{organisationId} [get]
func (c *NewController) GetSingleOrganisation(w http.ResponseWriter, r *http.Request) {
	log := c.Logger.NewContext()
	log.SetLogID(r.Header.Get("LogID"))
	log.Info("getting single organisation")
	apiResponse := utility.NewResponse()
	params := mux.Vars(r)
	organisationID := params["organisation_id"]

	organ, err := services.GetSingleOrganisation(organisationID, c.Repo, log)
	if err != nil {
		w.WriteHeader(err.StatusCode)
		json.NewEncoder(w).Encode(apiResponse.Error(err.ResponseCode))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(apiResponse.Success(utility.SYSTEM001, organ))
}

// GetBranches @Summary      Get list of branches
// @Description  Get the list of branches with pagination
// @Tags         branch
// @Param        page  query	int    false  "int valid"
// @Param        organisationId path	string    false  "string valid"
// @Success      200      {array}  models.Branch
// @Router       /branches/{organisationId} [get]
func (c *NewController) GetBranches(w http.ResponseWriter, r *http.Request) {
	log := c.Logger.NewContext()
	log.SetLogID(r.Header.Get("LogID"))
	log.Info("getting organisation")
	query := r.URL.Query()
	page := query.Get("page")
	if page == "" {
		page = "1"
	}
	apiResponse := utility.NewResponse()
	params := mux.Vars(r)
	organisationID := params["organisationId"]

	branch, err:= services.GetBranches(organisationID, page, c.Repo, log)
	if err != nil {
		w.WriteHeader(err.StatusCode)
		json.NewEncoder(w).Encode(apiResponse.Error(err.ResponseCode))
		return
	}

	json.NewEncoder(w).Encode(apiResponse.Success(utility.SYSTEM001, branch))
}

// GetSingleBranch @Summary      Get branch
// @Description  Get a single branch
// @Tags         branch
// @Param        branchId  path	string    false  "string valid"
// @Success      200      {object}  models.Organisation
// @Router       /branch/{branchId} [get]
func (c *NewController) GetSingleBranch(w http.ResponseWriter, r *http.Request) {
	log := c.Logger.NewContext()
	log.SetLogID(r.Header.Get("LogID"))
	log.Info("getting single branch")
	apiResponse := utility.NewResponse()
	params := mux.Vars(r)
	branchID := params["branchId"]

	branch, err := services.GetSingleBranch(branchID, c.Repo, log)
	if err != nil {
		w.WriteHeader(err.StatusCode)
		json.NewEncoder(w).Encode(apiResponse.Error(err.ResponseCode))
		return
	}

	json.NewEncoder(w).Encode(apiResponse.Success(utility.SYSTEM001, branch))
}

// OrganisationStatus @Summary      Setting the Status for an  Organisation
// @Description  Setting the Status for an  Organisation Operation
// @Tags         organisation
// @Accept       json
// @Produce      json
// @Param        default  body	dto.OrganStatus  true  "activate or deactivate"
// @Success      200      {object}  utility.Response
// @Router       /active-status [post]
func (c *NewController) OrganisationStatus(w http.ResponseWriter, r *http.Request) {
	log := c.Logger.NewContext()
	log.Info("deactivating and activating organisation")
	apiResponse := utility.NewResponse()
	var request *dto.OrganStatus
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Error("Error occured while parsing the request body, %s", err.Error())
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(apiResponse.Error(utility.BAD_REQUEST))
		return
	}

	_, statusErr := services.OrganisationStatus(request, c.Repo, log)
	if err != nil {
		w.WriteHeader(statusErr.StatusCode)
		json.NewEncoder(w).Encode(apiResponse.Error(statusErr.ResponseCode))
		return
	}

	json.NewEncoder(w).Encode(apiResponse.PlainSuccess(utility.SYSTEM001))

}
