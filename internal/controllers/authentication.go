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
	AuthenticateEmail(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	RegisterRoutes(router *mux.Router)
}

type NewAuthController struct {
	Service services.AuthService
}

func (c *NewAuthController) RegisterRoutes(router *mux.Router) {
	_ = router.PathPrefix("/auth").Subrouter()
}

func DefaultAuthController() AuthController {
	return &NewAuthController{
		Service: services.NewAuthService(),
	}
}

// AuthenticateEmail @Summary     Request to authentication token
// @Description  Request to authentication token
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        default  body	types.EmailRequest  true  "Authenticate existing user"
// @Success      200      {object}  utils.Response
// @Router       /request-login-token [post]
func (c *NewAuthController) AuthenticateEmail(w http.ResponseWriter, r *http.Request) {
	logger := log.WithFields(log.Fields{consts.RequestIdentifier: utils.GenerateUUID()})
	logger.Info("Verify user email and send login token.")

	requestData := &types.EmailRequest{}

	json.NewDecoder(r.Body).Decode(requestData)
	logger.Info("Request data: %+v", requestData)

	if validation.ValidateStruct(w, requestData, logger) {
		return
	}

	if err := c.Service.AuthEmail(requestData); err != nil {
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
		Message: "Successful",
	})

}

// Login @Summary     Login
// @Description  Authenticate user and get jwt token
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        default  body	types.AuthRequest  true  "Login to account"
// @Success      200      {object}  types.JWTResponse
// @Router       /login [post]
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
// @Router       /user-registration [post]
func (c *NewAuthController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	logger := log.WithFields(log.Fields{consts.RequestIdentifier: utils.GenerateUUID()})
	logger.Info("Adding User")

	requestData := &types.Registration{}

	json.NewDecoder(r.Body).Decode(requestData)

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
