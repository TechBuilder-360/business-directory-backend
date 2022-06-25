package routers

import (
	"github.com/TechBuilder-360/business-directory-backend/controllers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func SetupRoutes(router *mux.Router) {
	var (
		organisationController = controllers.DefaultOrganisationController()
		branchController       = controllers.DefaultBranchController()
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
