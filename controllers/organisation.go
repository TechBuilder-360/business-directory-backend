package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/araddon/dateparse"
	"github.com/go-playground/validator/v10"
	"net/http"

	"github.com/TechBuilder-360/business-directory-backend/dto"
	"github.com/TechBuilder-360/business-directory-backend/utility"
	"github.com/google/uuid"
)


//CreateOrganisation ...
func (c *NewController) CreateOrganisation(w http.ResponseWriter, r *http.Request){
	log:= c.Logger.NewContext()
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

	// Pass request to repo
    response ,err = c.Repo.CreateOrganisation(requestData)
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
