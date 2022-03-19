package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/TechBuilder-360/business-directory-backend/dto"
	"github.com/TechBuilder-360/business-directory-backend/utility"
	"github.com/google/uuid"
)





func (c *NewController) CreateOrganisation(w http.ResponseWriter, r *http.Request){
	log:= c.Logger.NewContext()
	log.SetLogID(r.Header.Get("LogID"))
	log.Info("creating organisation")
	r.Header.Set("TraceID", uuid.New().String())
	response := utility.NewResponse()
      var Organs *dto.CreateOrganisation
      err := json.NewDecoder(r.Body).Decode(&Organs)
      if err != nil {
	log.Error("Error occured while parsing the request body, %s",err.Error())
	json.NewEncoder(w).Encode(response.Error(utility.BAD_REQUEST, utility.GetCodeMsg(utility.BAD_REQUEST)))
	w.WriteHeader(http.StatusBadRequest)
	      return
      }
    OrganisationId,err := c.Repo.CreateOrganisation(Organs)
      if err != nil {
	log.Error("Error occured while creating organisation, %s",err.Error())
	json.NewEncoder(w).Encode(response.Error(utility.SERVER_ERROR, utility.GetCodeMsg(utility.SERVER_ERROR)))
	w.WriteHeader(http.StatusInternalServerError)
	      return
      }
      
      json.NewEncoder(w).Encode(response.Success(utility.SYSTEM001, utility.GetCodeMsg(utility.SYSTEM001),OrganisationId))
	w.WriteHeader(http.StatusOK)

   
} 