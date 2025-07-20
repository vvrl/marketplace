package handlers

import (
	"marketplace/internal/logger"
	"marketplace/internal/models"
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
	type RegisterRequest struct {
		Login    string `json:"login" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	var req RegisterRequest

	if err := c.Bind(&req); err != nil {
		logger.Logger.Error("failed to bind register request body")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	if req.Login == "" || req.Password == "" {
		logger.Logger.Error("invalid register request")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "login and password required"})
	}

	var user *models.User
	user, err := h.service.Register(c.Request().Context(), req.Login, req.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}

func (h *userHandler) Login(c echo.Context) error {

	return c.JSON(http.StatusOK, map[string]string{"status": "login is done"})
}
