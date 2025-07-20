package handlers

import (
	"marketplace/internal/services"

	"github.com/labstack/echo/v4"
)

type adHandler struct {
	service services.AdService
}

func NewAdHandler(s services.AdService) *adHandler {
	return &adHandler{service: s}
}

func (h *adHandler) PostAd(c echo.Context) error {
	return nil
}

func (h *adHandler) GetAdList(c echo.Context) error {
	return nil
}
