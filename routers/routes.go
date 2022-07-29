package routers

import (
	controllers2 "github.com/TechBuilder-360/business-directory-backend/internal/controllers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func SetupRoutes(router *mux.Router) {
	var (
		organisationController = controllers2.DefaultOrganisationController()
		branchController       = controllers2.DefaultBranchController()
	)

	//*******************************************
	//******* ORGANISATION **********************
	//*******************************************
	organisationController.RegisterRoutes(router)

	//*************************************
	//******* BRANCH **********************
	//*************************************
	branchController.RegisterRoutes(router)

	log.Info("Routes have been initialized")
}
