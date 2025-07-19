package handlers

import "github.com/labstack/echo/v4"

func SetAPI(e *echo.Echo) {
	e.GET("/", HelloHandler)
}
