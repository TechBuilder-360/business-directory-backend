package controllers

import (
	"encoding/json"
	"github.com/TechBuilder-360/business-directory-backend/dto"
	"github.com/TechBuilder-360/business-directory-backend/services"
	"github.com/TechBuilder-360/business-directory-backend/utility"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// RegisterUser @Summary     Register a new user
// @Description  Register a new user
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        default  body	dto.Registration  true  "Add a new user"
// @Success      200      {object}  utility.Response
// @Router       /user-registration [post]
func (c *NewController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	log := c.Logger.NewContext()
	log.SetLogID(r.Header.Get("LogID"))
	log.Info("Adding User")

	apiResponse := utility.NewResponse()
	requestData := &dto.Registration{}

	json.NewDecoder(r.Body).Decode(requestData)
	log.Info("Request data: %+v", requestData)

	validationRes := validator.New()
	if err := validationRes.Struct(requestData); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		log.Error("Validation failed on some fields : %+v", validationErrors)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(apiResponse.ValidationError(utility.VALIDATIONERR, validationErrors.Error()))
		return
	}

	err:=services.RegisterUser(requestData, c.Repo, log)
	if err != nil {
		w.WriteHeader(err.StatusCode)
		json.NewEncoder(w).Encode(apiResponse.Error(err.ResponseCode))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(apiResponse.PlainSuccess(utility.SYSTEM001))
	return

}

