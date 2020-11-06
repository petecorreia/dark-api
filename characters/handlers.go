package characters

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// CharactersHandler returns all character information
func CharactersHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, GetAll())
}

// CharacterHandler returns information about a specific character
func CharacterHandler(c echo.Context) error {
	id := c.Param("id")

	character, err := GetById(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Character not found")
	}

	return c.JSON(http.StatusOK, character)
}
