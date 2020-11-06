package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

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

func IndexHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Dark API")
}

func CharactersHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, Characters)
}

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
	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "8000"
	}

	e := echo.New()

	e.GET("/", IndexHandler)

	e.GET("/characters", CharactersHandler)
	e.GET("/characters/:id", CharacterHandler)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
