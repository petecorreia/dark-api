package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

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
	e := echo.New()

	e.GET("/", IndexHandler)

	e.GET("/characters", CharactersHandler)
	e.GET("/characters/:id", CharacterHandler)

	e.Logger.Fatal(e.Start(":8000"))
}
