package controllers

import (
	"encoding/json"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/consts"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/types"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/utils"
	"github.com/TechBuilder-360/business-directory-backend/internal/services"
	"github.com/TechBuilder-360/business-directory-backend/internal/validation"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type AuthController interface {
	RegisterUser(w http.ResponseWriter, r *http.Request)
	ActivateEmail(w http.ResponseWriter, r *http.Request)
	AuthenticateEmail(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	RegisterRoutes(router *mux.Router)
}

type NewAuthController struct {
	Service services.AuthService
}

func (c *NewAuthController) RegisterRoutes(router *mux.Router) {
	apis := router.PathPrefix("/auth").Subrouter()
	apis.HandleFunc("/register", c.RegisterUser).Methods("POST")
	apis.HandleFunc("/activate/{token}/{email}", c.ActivateEmail).Methods("GET")
	apis.HandleFunc("/authentication", c.AuthenticateEmail).Methods("POST")
	apis.HandleFunc("/login", c.Login).Methods("POST")
	apis.HandleFunc("/resend", c.ResendToken).Methods("POST")
}

func DefaultAuthController() AuthController {
	return &NewAuthController{
		Service: services.NewAuthService(),
	}
}

// ResendToken @Summary     Resend authentication token
// @Description  Request to authentication token
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        default  body	types.EmailRequest  true  "Authenticate existing user"
// @Success      200      {object}  utils.Response
// @Router       /auth/resend [post]

func (c *NewAuthController) ResendToken(w http.ResponseWriter, r *http.Request) {
	logger := log.WithFields(log.Fields{consts.RequestIdentifier: utils.GenerateUUID()})
	logger.Info("Resending login token.")

	requestData := &types.EmailRequest{}
	json.NewDecoder(r.Body).Decode(requestData)
	logger.Info("Request data: %+v", requestData)

	if validation.ValidateStruct(w, requestData, logger) {
		return
	}
	message, data, err := c.Service.ResendToken(requestData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logger.Error(err.Error())
		json.NewEncoder(w).Encode(utils.ErrorResponse{
			Status:  false,
			Message: "email authentication failed",
			Error:   err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.SuccessResponse{
		Status:  true,
		Message: message,
		Data:    data,
	})

}

// AuthenticateEmail @Summary     Request to authentication token
// @Description  Request to authentication token
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        default  body	types.EmailRequest  true  "Authenticate existing user"
// @Success      200      {object}  utils.Response
// @Router       /auth/authentication [post]
func (c *NewAuthController) AuthenticateEmail(w http.ResponseWriter, r *http.Request) {
	logger := log.WithFields(log.Fields{consts.RequestIdentifier: utils.GenerateUUID()})
	logger.Info("Verify user email and send login token.")

	requestData := &types.EmailRequest{}

	json.NewDecoder(r.Body).Decode(requestData)
	logger.Info("Request data: %+v", requestData)

	if validation.ValidateStruct(w, requestData, logger) {
		return
	}
	message, data, err := c.Service.AuthEmail(requestData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logger.Error(err.Error())
		json.NewEncoder(w).Encode(utils.ErrorResponse{
			Status:  false,
			Message: "email authentication failed",
			Error:   err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.SuccessResponse{
		Status:  true,
		Message: message,
		Data:    data,
	})

}

// Login @Summary     Login
// @Description  Authenticate user and get jwt token
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        default  body	types.AuthRequest  true  "Login to account"
// @Success      200      {object}  types.JWTResponse
// @Router       /auth/login [post]
func (c *NewAuthController) Login(w http.ResponseWriter, r *http.Request) {
	logger := log.WithFields(log.Fields{consts.RequestIdentifier: utils.GenerateUUID()})
	logger.Info("Verify User email and send login token.")

	requestData := &types.AuthRequest{}
	response := &types.JWTResponse{}

	json.NewDecoder(r.Body).Decode(requestData)

	if validation.ValidateStruct(w, requestData, logger) {
		return
	}

	response, err := c.Service.Login(requestData)
	if err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.ErrorResponse{
			Status:  false,
			Message: "login failed",
			Error:   err.Error(),
		})
		return
	}

	logger.Info("Response successful")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.SuccessResponse{
		Status:  true,
		Message: "Successful",
		Data:    response,
	})

}

// RegisterUser @Summary     Register a new user
// @Description  Register a new user
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        default  body	types.Registration  true  "Add a new user"
// @Success      200      {object}  utils.Response
// @Router       /auth/register [post]
func (c *NewAuthController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	logger := log.WithFields(log.Fields{consts.RequestIdentifier: utils.GenerateUUID()})
	logger.Info("Adding User")

	requestData := &types.Registration{}

	if err := json.NewDecoder(r.Body).Decode(requestData); err != nil {
		logger.Error(err.Error())
		json.NewEncoder(w).Encode(utils.ErrorResponse{
			Status:  false,
			Message: "Invalid request",
			Error:   err.Error(),
		})
		return
	}

	if validation.ValidateStruct(w, requestData, logger) {

		return
	}

	profile, err := c.Service.RegisterUser(requestData, logger)
	if err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.ErrorResponse{
			Status:  false,
			Message: err.Error(),
			Error:   err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(utils.SuccessResponse{
		Status:  true,
		Message: "Successful",
		Data:    profile,
	})
	return

}

// ActivateEmail @Summary     activation email
// @Description  Request to verification token
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        default  body	types.EmailRequest  true  "Authenticate existing user"
// @Success      200      {object}  utils.Response
// @Router       /auth/activate/{token}/{email} [post]
func (c *NewAuthController) ActivateEmail(w http.ResponseWriter, r *http.Request) {
	logger := log.WithFields(log.Fields{consts.RequestIdentifier: utils.GenerateUUID()})
	logger.Info("Activating User")
	vars := mux.Vars(r)
	key := vars["email"]
	token := vars["token"]

	message, err := c.Service.ActivateEmail(token, key, logger)
	if err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.ErrorResponse{
			Status:  false,
			Message: err.Error(),
			Error:   err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(utils.SuccessResponse{
		Status:  true,
		Message: message,
	})
	return
}
