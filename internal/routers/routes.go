package routers

import (
	"github.com/TechBuilder-360/business-directory-backend/domain/business/controller"
	controller2 "github.com/TechBuilder-360/business-directory-backend/domain/user/controller"
	"github.com/TechBuilder-360/business-directory-backend/graph"
	"github.com/TechBuilder-360/business-directory-backend/graph/generated"
	"github.com/arsmn/fastgql/graphql/handler"
	"github.com/arsmn/fastgql/graphql/playground"

	"github.com/TechBuilder-360/business-directory-backend/internal/configs"
	"github.com/TechBuilder-360/business-directory-backend/internal/controllers"
	"github.com/TechBuilder-360/business-directory-backend/internal/middleware"
	"github.com/TechBuilder-360/business-directory-backend/pkg/apm/sentry"
	"github.com/TechBuilder-360/business-directory-backend/pkg/log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
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
	//******* Controller ************************
	//*******************************************
	controller.RegisterRoutes(app)

	//*******************************************
	//******* GraphQL **********************
	//*******************************************
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	gqlHandler := srv.Handler()
	query := playground.Handler("GraphQL query", "/query")

	app.All("/query", func(c *fiber.Ctx) error {
		gqlHandler(c.Context())
		return nil
	})

	app.All("/gql", func(c *fiber.Ctx) error {
		query(c.Context())
		return nil
	})

	//*******************************************
	//******* End GraphQL **********************
	//*******************************************

	r := app.Group("/api")

	//*******************************************
	//******* Authentication **********************
	//*******************************************
	authController.RegisterRoutes(r)

	//*******************************************
	//******* ORGANISATION **********************
	//*******************************************
	organisationController.RegisterRoutes(r)

	//*************************************
	//******* BRANCH **********************
	//*************************************
	branchController.RegisterRoutes(r)

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
