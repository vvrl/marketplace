package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HelloHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"title": "Graduating",
		"body":  "Hello, world!",
	})
}
