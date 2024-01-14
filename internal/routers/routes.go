package routers

import (
	"github.com/TechBuilder-360/business-directory-backend/internal/configs"
	"github.com/TechBuilder-360/business-directory-backend/internal/controllers"
	"github.com/TechBuilder-360/business-directory-backend/internal/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	log "github.com/sirupsen/logrus"
	"time"
)

func SetupRoutes() *fiber.App {
	var (
		organisationController = controllers.DefaultOrganisationController()
		branchController       = controllers.DefaultBranchController()
		authController         = controllers.DefaultAuthController()
		usersController        = controllers.DefaultUserController()
		controller             = controllers.DefaultController()
	)

	app := fiber.New(fiber.Config{
		ErrorHandler:          middlewares.DefaultErrorHandler,
		DisableStartupMessage: true,
		StrictRouting:         true,
		ReadTimeout:           30 * time.Second,
		ReadBufferSize:        4096,
	})

	app.Use(recover.New())

	//*******************************************
	//******* Controller **********************
	//*******************************************
	controller.RegisterRoutes(app)

	//*******************************************
	//******* Authentication **********************
	//*******************************************
	authController.RegisterRoutes(app)

	//*******************************************
	//******* ORGANISATION **********************
	//*******************************************
	organisationController.RegisterRoutes(app)

	//*************************************
	//******* BRANCH **********************
	//*************************************
	branchController.RegisterRoutes(app)

	//*************************************
	//******* USERS **********************
	//*************************************
	usersController.RegisterRoutes(app)

	if !configs.Instance.IsProduction() {
		app.Get("/swagger/*", swagger.HandlerDefault)
	}

	log.Info("Routes have been initialized")
	return app
}
