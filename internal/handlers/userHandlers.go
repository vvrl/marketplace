package handlers

import (
	"marketplace/internal/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HelloHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"title": "Graduating",
		"body":  "Hello, world!",
	})
}

type userHandler struct {
	service services.UserService
}

func NewUserHandler(s services.UserService) *userHandler {
	return &userHandler{service: s}
}

func (h *userHandler) Register(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "regitser is done"})
}

func (h *userHandler) Login(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "login is done"})
}
