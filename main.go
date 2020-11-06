package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"github.com/petecorreia/dark-api/character"
	"github.com/petecorreia/dark-api/graphql"
	"github.com/petecorreia/dark-api/graphql/generated"
	"github.com/petecorreia/dark-api/world"
)

// IndexHandler handles the index route and displays a helpful list of the
// available endpoints
func IndexHandler(c echo.Context) error {
	const info = `Dark API

	POST /graphql            GraphQL API
	GET  /graphql            GraphQL playground

	GET  /characters         Get all characters
	GET  /characters/:id     Get a specific character by ID

	GET  /worlds             Get all worlds
	GET  /worlds/:id         Get a specific world by ID`

	return c.String(http.StatusOK, info)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	e := echo.New()

	e.GET("/", IndexHandler)

	// REST

	e.GET("/characters", character.CharactersHandler)
	e.GET("/characters/:id", character.CharacterHandler)
	e.GET("/worlds", world.WorldsHandler)
	e.GET("/worlds/:id", world.WorldHandler)

	// GraphQL

	graphqlHandler := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graphql.Resolver{}}))

	e.POST("/graphql", func(c echo.Context) error {
		graphqlHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	playgroundHandler := playground.Handler("GraphQL", "/graphql")

	e.GET("/graphql", func(c echo.Context) error {
		playgroundHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.HideBanner = true

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
