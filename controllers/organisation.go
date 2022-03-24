package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TechBuilder-360/business-directory-backend/dto"
	"github.com/TechBuilder-360/business-directory-backend/utility"
	"github.com/araddon/dateparse"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// @Summary      Add an organisation
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
	r.Header.Set("TraceID", uuid.New().String())
	apiResponse := utility.NewResponse()
	requestData := &dto.CreateOrgReq{}
	response := &dto.CreateOrgResponse{}
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		log.Error("Error occurred while parsing the request body, %s", err.Error())
		json.NewEncoder(w).Encode(apiResponse.Error(utility.BAD_REQUEST, utility.GetCodeMsg(utility.BAD_REQUEST)))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Debug("Request body: %+v", requestData)

	validationRes := validator.New()
	if err := validationRes.Struct(requestData); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		c.Logger.Error("Validation failed on some fields : %+v", validationErrors)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(apiResponse.ValidationError(utility.VALIDATIONERR, utility.GetCodeMsg(utility.VALIDATIONERR), validationErrors.Error()))
		return
	}

	_, err = dateparse.ParseLocal(requestData.FoundingDate)
	if err != nil {
		errMsg := fmt.Sprintf("Founding date is not a valid date. %s %s", requestData.FoundingDate, err.Error())
		log.Error(errMsg)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(apiResponse.Error(utility.BAD_REQUEST, errMsg))
		return
	}
	val := c.Repo.OrganisationExist(requestData)
	if val == true {
		log.Error("Organisation name Already Exist")
		json.NewEncoder(w).Encode(apiResponse.Error(utility.ALREADY_EXIST, utility.GetCodeMsg(utility.ALREADY_EXIST)))
		return
	}

	// Pass request to repo
	response, err = c.Repo.CreateOrganisation(requestData)
	if err != nil {
		log.Error("Error occurred while creating organisation, %s", err.Error())
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(apiResponse.Error(utility.SERVER_ERROR, utility.GetCodeMsg(utility.SERVER_ERROR)))
		return
	}

	log.Info("Response body: %+v", response)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(apiResponse.Success(utility.SYSTEM001, utility.GetCodeMsg(utility.SYSTEM001), response))
}

// @Summary      Add an Branch
// @Description  add by json Branch
// @Tags         branch
// @Accept       json
// @Produce      json
// @Param        default  body	dto.CreateBranch  true  "Add add branch"
// @Success      200      {object}  utility.ResponseObj
// @Router       /branch [post]
func (c *NewController) CreateBranch(w http.ResponseWriter, r *http.Request) {
	log := c.Logger.NewContext()
	log.SetLogID(r.Header.Get("LogID"))
	log.Info("creating branch")
	response := utility.NewResponse()
	var br *dto.CreateBranch
	err := json.NewDecoder(r.Body).Decode(&br)
	if err != nil {
		log.Error("Error occured while parsing the request body, %s", err.Error())
		json.NewEncoder(w).Encode(response.Error(utility.BAD_REQUEST, utility.GetCodeMsg(utility.BAD_REQUEST)))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Debug("Request body: %+v", &br)

	validationRes := validator.New()
	if err := validationRes.Struct(&br); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		c.Logger.Error("Validation failed on some fields : %+v", validationErrors)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response.ValidationError(utility.VALIDATIONERR, utility.GetCodeMsg(utility.VALIDATIONERR), validationErrors.Error()))
		return
	}
	val := c.Repo.BranchExist(br)
	if val == true {
		log.Error("Branch name Already Exist")
		json.NewEncoder(w).Encode(response.Error(utility.ALREADY_EXIST, utility.GetCodeMsg(utility.ALREADY_EXIST)))

		return
	}

	branchId, err := c.Repo.CreateBranch(br)
	if err != nil {
		log.Error("Error occured while creating branch, %s", err.Error())
		json.NewEncoder(w).Encode(response.Error(utility.SERVER_ERROR, utility.GetCodeMsg(utility.SERVER_ERROR)))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(response.Success(utility.SYSTEM001, utility.GetCodeMsg(utility.SYSTEM001), branchId))

}

// @Summary      Get list of organisation
// @Description  Get the list of organisation with pagination
// @Tags         organisation
// @Param        page  query	int    false  "int valid"
// @Success      200      {array}  models.Organisation
// @Router       /get-organisations [get]
func (c *NewController) GetOrganisations(w http.ResponseWriter, r *http.Request) {
	log := c.Logger.NewContext()
	log.SetLogID(r.Header.Get("LogID"))
	log.Info("getting organisation")
	query := r.URL.Query()
	page := query.Get("page")

	if page == "" {
		page = "1"
	}
	response := utility.NewResponse()

	data, err := c.Repo.GetOrganisations(page)
	if err != nil {
		log.Error("Error occured while getting list of organisations, %s", err.Error())
		json.NewEncoder(w).Encode(response.Error(utility.SERVER_ERROR, utility.GetCodeMsg(utility.SERVER_ERROR)))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(response.Success(utility.SYSTEM001, utility.GetCodeMsg(utility.SYSTEM001), data))
}

// @Summary      Get organisation
// @Description  Get a single organisation
// @Tags         organisation
// @Param        organisationId  path	string    false  "string valid"
// @Success      200      {object}  models.Organisation
// @Router       /organisation/{organisationId} [get]
func (c *NewController) GetSingleOrganisation(w http.ResponseWriter, r *http.Request) {
	log := c.Logger.NewContext()
	log.SetLogID(r.Header.Get("LogID"))
	log.Info("getting single organisation")
	response := utility.NewResponse()
	params := mux.Vars(r)
	organisationID := params["organisationId"]
	data, err := c.Repo.GetSingleOrganisation(organisationID)
	if err != nil {
		log.Error("Error occured while getting organisation, %s", err.Error())
		json.NewEncoder(w).Encode(response.Error(utility.SERVER_ERROR, utility.GetCodeMsg(utility.SERVER_ERROR)))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(response.Success(utility.SYSTEM001, utility.GetCodeMsg(utility.SYSTEM001), data))
}

// @Summary      Get list of branches
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
	response := utility.NewResponse()
	params := mux.Vars(r)
	organisationID := params["organisationId"]
	data, err := c.Repo.GetBranches(organisationID, page)
	if err != nil {
		log.Error("Error occured while getting organisation branches, %s", err.Error())
		json.NewEncoder(w).Encode(response.Error(utility.SERVER_ERROR, utility.GetCodeMsg(utility.SERVER_ERROR)))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(response.Success(utility.SYSTEM001, utility.GetCodeMsg(utility.SYSTEM001), data))
}

// @Summary      Get branch
// @Description  Get a single branch
// @Tags         branch
// @Param        branchId  path	string    false  "string valid"
// @Success      200      {object}  models.Organisation
// @Router       /branch/{branchId} [get]
func (c *NewController) GetSingleBranch(w http.ResponseWriter, r *http.Request) {
	log := c.Logger.NewContext()
	log.SetLogID(r.Header.Get("LogID"))
	log.Info("getting single branch")
	response := utility.NewResponse()
	params := mux.Vars(r)
	branchID := params["branchId"]
	data, err := c.Repo.GetSingleBranch(branchID)
	if err != nil {
		log.Error("Error occured while getting branch, %s", err.Error())
		json.NewEncoder(w).Encode(response.Error(utility.SERVER_ERROR, utility.GetCodeMsg(utility.SERVER_ERROR)))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(response.Success(utility.SYSTEM001, utility.GetCodeMsg(utility.SYSTEM001), data))
}

// @Summary      Setting the Status for an  Organisation
// @Description  Setting the Status for an  Organisation Operation
// @Tags         organisation
// @Accept       json
// @Produce      json
// @Param        default  body	dto.OrganStatus  true  "activate or deactivate"
// @Success      200      {object}  utility.ResponseObj
// @Router       /active-status [post]
func (c *NewController) OrganisationStatus(w http.ResponseWriter, r *http.Request) {
	log := c.Logger.NewContext()
	log.Info("deactivating and activating organisation")
	response := utility.NewResponse()
	var Organs *dto.OrganStatus
	err := json.NewDecoder(r.Body).Decode(&Organs)
	if err != nil {
		log.Error("Error occured while parsing the request body, %s", err.Error())
		json.NewEncoder(w).Encode(response.Error(utility.BAD_REQUEST, utility.GetCodeMsg(utility.BAD_REQUEST)))
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	_, err = c.Repo.OrganisationStatus(Organs)

	if err != nil {
		log.Error("Error occured while deactivating or activating organisation, %s", err.Error())
		json.NewEncoder(w).Encode(response.Error(utility.SERVER_ERROR, utility.GetCodeMsg(utility.SERVER_ERROR)))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(response.PlainSuccess(utility.SYSTEM001, utility.GetCodeMsg(utility.SYSTEM001)))

}
