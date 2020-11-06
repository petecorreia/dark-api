package world

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// WorldsHandler returns all world information
func WorldsHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, GetAll())
}

// WorldHandler returns information about a specific world
func WorldHandler(c echo.Context) error {
	id := c.Param("id")

	world, err := GetById(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "World not found")
	}

	return c.JSON(http.StatusOK, world)
}
