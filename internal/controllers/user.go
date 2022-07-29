package controllers

import (
	"github.com/TechBuilder-360/business-directory-backend/internal/services"
	"github.com/gorilla/mux"
)

type UserController interface {
	RegisterRoutes(router *mux.Router)
}

type NewUserController struct {
	Service services.UserService
}

func (n NewUserController) RegisterRoutes(router *mux.Router) {
	_ = router.PathPrefix("/users").Subrouter()
}

func DefaultUserController() UserController {
	return &NewUserController{
		Service: services.NewUserService(),
	}
}
