package routers

import (
	"github.com/TechBuilder-360/business-directory-backend/internal/controllers"
	"github.com/TechBuilder-360/business-directory-backend/middlewares"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/swaggo/http-swagger"
)

func SetupRoutes(router *mux.Router) {
	var (
		organisationController = controllers.DefaultOrganisationController()
		branchController       = controllers.DefaultBranchController()
		controller             = controllers.DefaultAuthController()
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

	router.PathPrefix("/documentation/").Handler(httpSwagger.WrapHandler)

	log.Info("Routes have been initialized")
}
