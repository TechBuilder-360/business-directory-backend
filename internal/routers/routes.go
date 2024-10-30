package routers

import (
	"github.com/TechBuilder-360/business-directory-backend/domain/business/controller"
	controller2 "github.com/TechBuilder-360/business-directory-backend/domain/user/controller"
	"github.com/TechBuilder-360/business-directory-backend/internal/configs"
	"github.com/TechBuilder-360/business-directory-backend/internal/controllers"
	"github.com/TechBuilder-360/business-directory-backend/internal/middleware"
	"github.com/TechBuilder-360/business-directory-backend/pkg/apm/sentry"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	log "github.com/sirupsen/logrus"
	"time"
)

func SetupRoutes() *fiber.App {
	var (
		organisationController = controller.DefaultBusinessController()
		branchController       = controller.DefaultBranchController()
		authController         = controllers.DefaultAuthController()
		usersController        = controller2.DefaultUserController()
		controller             = controllers.DefaultController()
	)

	app := fiber.New(fiber.Config{
		ErrorHandler:          middleware.DefaultErrorHandler,
		DisableStartupMessage: true,
		StrictRouting:         true,
		ReadTimeout:           30 * time.Second,
		ReadBufferSize:        4096,
	})

	app.Use(recover.New(), sentry.Middleware(), middleware.Logger)

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
