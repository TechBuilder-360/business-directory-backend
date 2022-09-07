package controllers

import (
	"encoding/json"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/constant"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/types"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/utils"
	"github.com/TechBuilder-360/business-directory-backend/internal/middlewares"
	"github.com/TechBuilder-360/business-directory-backend/internal/services"
	"github.com/TechBuilder-360/business-directory-backend/internal/validation"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type UserController interface {
	UpgradeUserTier(w http.ResponseWriter, r *http.Request)
	RegisterRoutes(router *mux.Router)
}

type NewUserController struct {
	as services.UserService
}

func (c NewUserController) RegisterRoutes(router *mux.Router) {
	apis := router.PathPrefix("/users").Subrouter()
	apis.HandleFunc("/upgrade", middlewares.Adapt(http.HandlerFunc(c.UpgradeUserTier),
		middlewares.AuthorizeUserJWT()).ServeHTTP).Methods(http.MethodPost)

}

func DefaultUserController() UserController {
	return &NewUserController{
		as: services.NewUserService(),
	}
}

func (c *NewUserController) UpgradeUserTier(w http.ResponseWriter, r *http.Request) {
	logger := log.WithFields(log.Fields{constant.RequestIdentifier: utils.GenerateUUID()})
	logger.Info("Upgrading user tiers")

	body := &types.UpgradeUserTierRequest{}
	json.NewDecoder(r.Body).Decode(body)
	logger.Info("Request data: %+v", body)

	if validation.ValidateStruct(w, body, logger) {
		return
	}

	err := c.as.UpgradeUserTier(body, logger)
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
		Message: "Upgraded",
	})
}
