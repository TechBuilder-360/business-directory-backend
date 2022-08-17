package controllers

import (
	"github.com/TechBuilder-360/business-directory-backend/internal/services"
	"github.com/gorilla/mux"
)

type UserController interface {
	RegisterRoutes(router *mux.Router)
}

type userController struct {
	Service services.UserService
}

func (n userController) RegisterRoutes(router *mux.Router) {
	_ = router.PathPrefix("/users").Subrouter()

}

func DefaultUserController() UserController {
	return &userController{
		Service: services.NewUserService(),
	}
}
