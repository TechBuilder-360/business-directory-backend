package controllers

import (
	"github.com/TechBuilder-360/business-directory-backend/internal/services"
	"github.com/gofiber/fiber/v2"
)

type IUserController interface {
	//UpgradeTier(ctx *fiber.Ctx) error
	RegisterRoutes(router *fiber.App)
}

type UserController struct {
	as services.UserService
}

func (c *UserController) RegisterRoutes(router *fiber.App) {
	_ = router.Group("/users")
	//apis.HandleFunc("/upgrade/tier-one", middlewares.Adapt(http.HandlerFunc(c.UpgradeTier),
	//	middlewares.AuthorizeUserJWT()).ServeHTTP).Methods(http.MethodPost)

}

func DefaultUserController() IUserController {
	return &UserController{
		as: services.NewUserService(),
	}
}

//func (c *UserController) UpgradeTier(ctx *fiber.Ctx) error {
//	logger := log.WithFields(log.Fields{constant.RequestIdentifier: utils.GenerateUUID()})
//	logger.Info("Upgrading user tiers")
//	body := &types.UpgradeUserTierRequest{}
//	json.NewDecoder(r.Body).Decode(body)
//	logger.Info("Request data: %+v", body)
//
//	if validation.ValidateStruct(w, body, logger) {
//		return
//	}
//
//	// get user from context
//	user, err := middlewares.UserFromContext(r)
//	if err != nil {
//		logger.Error(err.Error())
//		w.WriteHeader(http.StatusBadRequest)
//		json.NewEncoder(w).Encode(utils.ErrorResponse{
//			Status:  false,
//			Message: err.Error(),
//		})
//		return
//	}
//
//	err = c.as.UpgradeStatus(body, user, logger)
//	if err != nil {
//		w.WriteHeader(http.StatusBadRequest)
//		logger.Error(err.Error())
//		json.NewEncoder(w).Encode(utils.ErrorResponse{
//			Status:  false,
//			Message: err.Error(),
//		})
//		return
//	}
//
//	w.WriteHeader(http.StatusOK)
//	json.NewEncoder(w).Encode(utils.SuccessResponse{
//		Status:  true,
//		Message: "Tier one upgrade successful",
//	})
//}
