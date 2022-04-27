package controllers

import (
	"encoding/json"
	"github.com/TechBuilder-360/business-directory-backend/dto"
	"github.com/TechBuilder-360/business-directory-backend/services"
	"github.com/TechBuilder-360/business-directory-backend/utility"
	logger "github.com/Toflex/oris_log"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"net/http"
)


type AuthController interface {
	RegisterUser(w http.ResponseWriter, r *http.Request)
	AuthenticateEmail(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

type NewAuthController struct {
	Service  services.AuthService
	Logger     logger.Logger
}

func DefaultAuthController(serv services.AuthService, log logger.Logger) AuthController {
	return &NewAuthController{
		Service:    serv,
		Logger:     log,
	}
}

// AuthenticateEmail @Summary     Request to authentication token
// @Description  Request to authentication token
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        default  body	dto.EmailRequest  true  "Authenticate existing user"
// @Success      200      {object}  utility.Response
// @Router       /request-login-token [post]
func (c *NewAuthController) AuthenticateEmail(w http.ResponseWriter, r *http.Request) {
	log := c.Logger.NewContext()
	log.SetLogID(uuid.NewString())
	log.Info("Verify User email and send login token.")

	apiResponse := utility.NewResponse()
	requestData := &dto.EmailRequest{}

	json.NewDecoder(r.Body).Decode(requestData)
	log.Info("Request data: %+v", requestData)

	validationRes := validator.New()
	if err := validationRes.Struct(requestData); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		log.Error("Validation failed on some fields : %+v", validationErrors)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(apiResponse.ValidationError(utility.VALIDATIONERR, validationErrors.Error()))
		return
	}

	if err := c.Service.AuthEmail(requestData, log); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Error(err.Error())
		json.NewEncoder(w).Encode(apiResponse.Error(err.Error()))
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
func (c *NewAuthController) Login(w http.ResponseWriter, r *http.Request) {
	log := c.Logger.NewContext()
	log.SetLogID(uuid.NewString())
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
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(apiResponse.ValidationError(utility.VALIDATIONERR, validationErrors.Error()))
		return
	}

	response, err:=c.Service.Login(requestData, log)
	if err != nil {
		log.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(apiResponse.Error(err.Error()))
		return
	}

	log.Info("Response successful")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(apiResponse.Success(utility.SYSTEM001, response))

}

// RegisterUser @Summary     Register a new user
// @Description  Register a new user
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        default  body	dto.Registration  true  "Add a new user"
// @Success      200      {object}  utility.Response
// @Router       /user-registration [post]
func (c *NewAuthController) RegisterUser(w http.ResponseWriter, r *http.Request) {
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
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(apiResponse.ValidationError(utility.VALIDATIONERR, validationErrors.Error()))
		return
	}

	profile, err := c.Service.RegisterUser(requestData, log)
	if err != nil {
		log.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(apiResponse.CustomError(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(apiResponse.Success(utility.SYSTEM001, profile))
	return

}

