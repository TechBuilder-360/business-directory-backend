package controllers

import (
	"encoding/json"
	"github.com/TechBuilder-360/business-directory-backend/dto"
	"github.com/TechBuilder-360/business-directory-backend/services"
	"github.com/TechBuilder-360/business-directory-backend/utility"
	logger "github.com/Toflex/oris_log"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type BranchController interface {
}

type NewBranchController struct {
	Service  services.BranchService
	Logger     logger.Logger
}

func DefaultBranchController(serv services.BranchService, log logger.Logger) BranchController {
	return &NewBranchController{
		Service:    serv,
		Logger:     log,
	}
}

// CreateBranch @Summary      Add branch
// @Description  add branch to an organisation
// @Tags         branch
// @Accept       json
// @Produce      json
// @Param        default  body	dto.CreateOrgReq  true  "Add branch"
// @Success      200      {object}  utility.ResponseObj
// @Router       /organisation [post]
func (c *NewOrganisationController) CreateBranch(w http.ResponseWriter, r *http.Request) {
	log := c.Logger.NewContext()
	log.SetLogID(r.Header.Get("LogID"))

	apiResponse := utility.NewResponse()
	requestData := &dto.CreateOrgReq{}
	response := &dto.Organisation{}
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

	response, err = c.Service.CreateOrganisation(requestData, nil, log)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(apiResponse.Error(err.Error()))
		return
	}

	log.Info("Response body: %+v", response)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(apiResponse.Success(utility.SYSTEM001, response))
}
