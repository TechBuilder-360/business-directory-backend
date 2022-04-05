package controllers

import (
	"encoding/json"
	"github.com/TechBuilder-360/business-directory-backend/dto"
	"github.com/TechBuilder-360/business-directory-backend/services"
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
		json.NewEncoder(w).Encode(apiResponse.ValidationError(utility.VALIDATIONERR, validationErrors.Error()))
		return
	}

	if err := services.AuthEmail(requestData, c.Repo, log); err != nil {
		w.WriteHeader(err.StatusCode)
		json.NewEncoder(w).Encode(apiResponse.Error(err.ResponseCode))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(apiResponse.PlainSuccess(utility.SYSTEM001))

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
	response := &dto.JWTResponse{}

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

	response, err:=services.Login(requestData, c.Repo, c.JWTService, log)
	if err != nil {
		w.WriteHeader(err.StatusCode)
		json.NewEncoder(w).Encode(apiResponse.Error(err.ResponseCode))
		return
	}

	log.Info("Response successful")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(apiResponse.Success(utility.SYSTEM001, response))

}
