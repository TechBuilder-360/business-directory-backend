package controllers

import (
	"github.com/TechBuilder-360/business-directory-backend/internal/services"
	"github.com/gorilla/mux"
)

type OrganisationController interface {
	RegisterRoutes(router *mux.Router)
}

type organisationController struct {
	Service services.OrganisationService
}

func (c *organisationController) RegisterRoutes(router *mux.Router) {
	_ = router.PathPrefix("/organisations").Subrouter()

}

func DefaultOrganisationController() OrganisationController {
	return &organisationController{
		Service: services.NewOrganisationService(),
	}
}
