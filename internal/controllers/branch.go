package controllers

import (
	"github.com/TechBuilder-360/business-directory-backend/internal/services"
	"github.com/gorilla/mux"
)

type BranchController interface {
	RegisterRoutes(router *mux.Router)
}

type NewBranchController struct {
	Service services.BranchService
}

func (c *NewBranchController) RegisterRoutes(router *mux.Router) {
	_ = router.PathPrefix("/branches").Subrouter()
}

func DefaultBranchController() BranchController {
	return &NewBranchController{
		Service: services.NewBranchService(),
	}
}
