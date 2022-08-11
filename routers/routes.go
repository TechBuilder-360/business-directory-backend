package routers

import (
	controllers "github.com/TechBuilder-360/business-directory-backend/internal/controllers"
	"github.com/TechBuilder-360/business-directory-backend/middlewares"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func SetupRoutes(router *mux.Router) {
	var (
		organisationController = controllers.DefaultOrganisationController()
		branchController       = controllers.DefaultBranchController()
		controller             = controllers.DefaultController()
	)

	router.Use(middlewares.Recovery)

	//*******************************************
	//******* Controller **********************
	//*******************************************
	controller.RegisterRoutes(router)

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
