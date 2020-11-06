package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

// Character represents one of the Dark characters
type Character struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	Worlds        []string `json:"worlds"`
	Aliases       []string `json:"aliases"`
	Alternates    []string `json:"alternates"`
	Parents       []string `json:"parents"`
	Children      []string `json:"children"`
	Relationships []string `json:"relationships"`
}

// IndexHandler handles the index route and displays a helpful list of the
// available endpoints
func IndexHandler(c echo.Context) error {
	const info = `Dark API

	/characters         gets all characters
	/characters/:id     get a specific character`

	return c.String(http.StatusOK, info)
}

// CharactersHandler returns all character information
func CharactersHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, Characters)
}

// CharacterHandler returns information about a specific character
func CharacterHandler(c echo.Context) error {
	id := c.Param("id")

	var character *Character

	for _, c := range Characters {
		if c.ID == id {
			character = &c
			break
		}
	}

	if character == nil {
		return echo.NewHTTPError(http.StatusNotFound, "Character not found")
	}

	return c.JSON(http.StatusOK, character)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	e := echo.New()

	e.GET("/", IndexHandler)

	e.GET("/characters", CharactersHandler)
	e.GET("/characters/:id", CharacterHandler)

	e.HideBanner = true

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
