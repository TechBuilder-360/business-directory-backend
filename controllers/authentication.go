package controllers

import (
	"encoding/json"
	"github.com/TechBuilder-360/business-directory-backend/dto"
	"github.com/TechBuilder-360/business-directory-backend/models"
	"github.com/TechBuilder-360/business-directory-backend/utility"
	"github.com/go-playground/validator/v10"
	"net/http"
)


// AuthenticateEmail @Summary     Request to authentication token
// @Description  Request to authentication token
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        default  body	dto.EmailRequest  true  "Authenticate existing user"
// @Success      200      {object}  utility.Response
// @Router       /request-login-token [post]
func (c *NewController) AuthenticateEmail(w http.ResponseWriter, r *http.Request) {
	log := c.Logger.NewContext()
	log.SetLogID(r.Header.Get("LogID"))
	log.Info("Verify User email and send login token.")

	apiResponse := utility.NewResponse()
	requestData := &dto.EmailRequest{}

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
	if !ok {
		log.Info("Email address does not exist. '%s'", requestData.EmailAddress)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(apiResponse.Error(utility.EMAILDOESNOTEXIST, utility.GetCodeMsg(utility.EMAILDOESNOTEXIST)))
		return
	}

	token, err := c.Repo.CreateUserToken(requestData.EmailAddress)
	if err != nil {
		log.Error("Error occurred when saving sign-in token. %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(apiResponse.Error(utility.SMMERROR, utility.GetCodeMsg(utility.SMMERROR)))
		return
	}

	// TODO: Send Token to user email
	token = token

	// Activity log
	activity := &models.Activity{Message: "Requested for sign in token"}
	go func() {
		if err := c.Repo.AddActivity(activity); err!=nil {
			log.Error("User activity failed to log")
		}
	}()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(apiResponse.PlainSuccess(utility.SYSTEM001, utility.GetCodeMsg(utility.SYSTEM001)))

}

// Login @Summary     Login
// @Description  Authenticate user and get jwt token
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        default  body	dto.AuthRequest  true  "Login to account"
// @Success      200      {object}  dto.JWTResponse
// @Router       /login [post]
func (c *NewController) Login(w http.ResponseWriter, r *http.Request) {
	log := c.Logger.NewContext()
	log.SetLogID(r.Header.Get("LogID"))
	log.Info("Verify User email and send login token.")

	apiResponse := utility.NewResponse()
	requestData := &dto.AuthRequest{}
	response := dto.JWTResponse{}

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

	// Validate user token
	isTokenValid, err := c.Repo.IsTokenValid(requestData)
	if err != nil {
		log.Error("An Error occurred when validating login token. %s", err.Message)
		w.WriteHeader(http.StatusFailedDependency)
		json.NewEncoder(w).Encode(apiResponse.Error(utility.SMMERROR, err.Message))
		return
	}
	if isTokenValid == false {
		log.Error("Invalid login token. %s", err.AsMessage())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(apiResponse.Error(utility.SMMERROR, utility.GetCodeMsg(utility.SMMERROR)))
		return
	}

	profile, profileErr := c.Repo.GetUserInformation(requestData.Email)
	if profileErr != nil {
		log.Error("An error occurred when fetching user profile. %s", profileErr.Error())
		w.WriteHeader(http.StatusFailedDependency)
		json.NewEncoder(w).Encode(apiResponse.Error(utility.SMMERROR, utility.GetCodeMsg(utility.SMMERROR)))
		return
	}

	// Generate JWT for user
	token, Tokenerr := c.JWTService.GenerateToken(profile.ID)
	if Tokenerr != nil {
		log.Error("An error occurred when generating jwt token. %s", Tokenerr.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(apiResponse.Error(utility.SMMERROR, utility.GetCodeMsg(utility.SMMERROR)))
		return
	}

	response.Profile = profile
	response.AccessToken = token

	// Activity log
	activity := &models.Activity{By: response.Profile.ID, Message: "Successful login"}
	go func() {
		if err := c.Repo.AddActivity(activity); err!=nil {
			log.Error("User activity failed to log")
		}
	}()

	log.Info("Response successful")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(apiResponse.Success(utility.SYSTEM001, utility.GetCodeMsg(utility.SYSTEM001), response))

}
