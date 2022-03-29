package controllers

import (
	"encoding/json"
	"github.com/TechBuilder-360/business-directory-backend/dto"
	"github.com/TechBuilder-360/business-directory-backend/models"
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
		json.NewEncoder(w).Encode(apiResponse.ValidationError(utility.VALIDATIONERR, utility.GetCodeMsg(utility.VALIDATIONERR), validationErrors.Error()))
		return
	}

	// Check if email address exist
	ok,err:=c.Repo.DoesUserEmailExist(requestData.EmailAddress)
	if err != nil {
		log.Error("An Error occurred while checking if user email exist. %s", err.Error())
		w.WriteHeader(http.StatusFailedDependency)
		json.NewEncoder(w).Encode(apiResponse.Error(utility.SMMERROR, utility.GetCodeMsg(utility.SMMERROR)))
		return
	}
	if ok {
		log.Info("Email address already exist. '%s'", requestData.EmailAddress)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(apiResponse.Error(utility.EMAILALREADYEXIST, utility.GetCodeMsg(utility.EMAILALREADYEXIST)))
		return
	}

	// Save user details
	userId, err := c.Repo.RegisterUser(requestData)
	if err != nil {
		log.Error("Error occurred when saving new user. %s", err.Error())
		w.WriteHeader(http.StatusFailedDependency)
		json.NewEncoder(w).Encode(apiResponse.Error(utility.SMMERROR, utility.GetCodeMsg(utility.SMMERROR)))
		return
	}

	// Activity log
	activity := &models.Activity{By: userId, Message: "Registered"}
	go func() {
		if err = c.Repo.AddActivity(activity); err!=nil {
			log.Error("User activity failed to log")
		}
	}()


	// TODO: Send Activate email

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(apiResponse.PlainSuccess(utility.SYSTEM001, utility.GetCodeMsg(utility.SYSTEM001)))
	return

}

