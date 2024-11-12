package main

import (
	"log"
	"os"

	"github.com/TechBuilder-360/business-directory-backend/graph"
	"github.com/TechBuilder-360/business-directory-backend/graph/generated"
	"github.com/arsmn/fastgql/graphql/handler"
	"github.com/arsmn/fastgql/graphql/playground"
	fiber "github.com/gofiber/fiber/v2"
)

const defaultPort = "8080"

func main2() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	app := fiber.New()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	gqlHandler := srv.Handler()
	playground := playground.Handler("GraphQL playground", "/query")

	app.All("/query", func(c *fiber.Ctx) error {
		gqlHandler(c.Context())
		return nil
	})

	app.All("/", func(c *fiber.Ctx) error {
		playground(c.Context())
		return nil
	})

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(app.Listen(":" + port))
}
