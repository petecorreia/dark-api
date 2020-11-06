package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func IndexHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func CharactersHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, Characters)
}

func main() {
	e := echo.New()

	e.GET("/", IndexHandler)
	e.GET("/characters", CharactersHandler)

	e.Logger.Fatal(e.Start(":8000"))
}
