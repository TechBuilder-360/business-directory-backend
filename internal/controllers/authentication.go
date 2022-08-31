package controllers

import (
	"encoding/json"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/constant"
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
	Authenticate(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	RegisterRoutes(router *mux.Router)
}

type NewAuthController struct {
	as services.AuthService
}

func (c *NewAuthController) RegisterRoutes(router *mux.Router) {
	apis := router.PathPrefix("/auth").Subrouter()

	apis.HandleFunc("/register", c.RegisterUser).Methods(http.MethodPost)
	apis.HandleFunc("/activate", c.ActivateEmail).Methods(http.MethodGet)
	apis.HandleFunc("/authentication", c.Authenticate).Methods(http.MethodPost)
	apis.HandleFunc("/login", c.Login).Methods(http.MethodPost)
}

func DefaultAuthController() AuthController {
	return &NewAuthController{
		as: services.NewAuthService(),
	}
}

// Authenticate
// @Summary      request to authentication token
// @Description  Request to authentication token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        default  body	types.EmailRequest  true  "Authenticate existing user"
// @Success      200      {object}  utils.SuccessResponse
// @Router       /auth/authentication [post]
func (c *NewAuthController) Authenticate(w http.ResponseWriter, r *http.Request) {
	logger := log.WithFields(log.Fields{constant.RequestIdentifier: utils.GenerateUUID()})
	logger.Info("Authenticate")

	body := &types.EmailRequest{}
	json.NewDecoder(r.Body).Decode(body)
	logger.Info("Request data: %+v", body)

	if validation.ValidateStruct(w, body, logger) {
		return
	}

	err := c.as.RequestToken(body, logger)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logger.Error(err.Error())
		json.NewEncoder(w).Encode(utils.ErrorResponse{
			Status:  false,
			Message: err.Error(),
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
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        default  body	types.AuthRequest  true  "Login to account"
// @Success      200      {object}  utils.SuccessResponse{Data=types.JWTResponse}
// @Router       /auth/login [post]
func (c *NewAuthController) Login(w http.ResponseWriter, r *http.Request) {
	logger := log.WithFields(log.Fields{constant.RequestIdentifier: utils.GenerateUUID()})
	logger.Info("Verify User email and send login token.")

	body := &types.AuthRequest{}

	err := json.NewDecoder(r.Body).Decode(body)
	if err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.ErrorResponse{
			Status:  false,
			Message: "bad request",
		})
		return
	}

	if validation.ValidateStruct(w, body, logger) {
		return
	}

	response, err := c.as.Login(body)
	if err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.SuccessResponse{
		Status:  true,
		Message: "Successful",
		Data:    response,
	})

}

// RegisterUser @Summary     Register a new user
// @Description  Register a new user
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        default  body	types.Registration  true  "Add a new user"
// @Success      200      {object}  utils.SuccessResponse
// @Router       /auth/register [post]
func (c *NewAuthController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	logger := log.WithFields(log.Fields{constant.RequestIdentifier: utils.GenerateUUID()})
	logger.Info("Adding User")

	requestData := &types.Registration{}

	if err := json.NewDecoder(r.Body).Decode(requestData); err != nil {
		logger.Error(err.Error())
		json.NewEncoder(w).Encode(utils.ErrorResponse{
			Status:  false,
			Message: "Invalid request",
		})
		return
	}

	if validation.ValidateStruct(w, requestData, logger) {

		return
	}

	err := c.as.RegisterUser(requestData, logger)
	if err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(utils.SuccessResponse{
		Status:  true,
		Message: "Successful",
	})
}

// ActivateEmail @Summary     activation email
// @Description  Request to verification token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        token    query     string  false  "token"
// @Param        uid    query     string  false  "uid"
// @Success      200      {object}  utils.SuccessResponse
// @Router       /auth/activate [get]
func (c *NewAuthController) ActivateEmail(w http.ResponseWriter, r *http.Request) {
	logger := log.WithFields(log.Fields{constant.RequestIdentifier: utils.GenerateUUID()})
	logger.Info("Activating User")

	vars := r.URL.Query()

	uid := vars.Get("uid")
	token := vars.Get("token")

	err := c.as.ActivateEmail(token, uid, logger)
	if err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(utils.SuccessResponse{
		Status:  true,
		Message: "account activation successful",
	})
	return
}
