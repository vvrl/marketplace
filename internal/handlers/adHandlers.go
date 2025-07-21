package handlers

import (
	"marketplace/internal/logger"
	"marketplace/internal/services"
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo/v4"
)

type PostAdRequest struct {
	Title    string  `json:"title" validate:"required,min=5,max=100"`
	Text     string  `json:"text" validate:"required,min=10,max=1000"`
	ImageURL string  `json:"image_url" validate:"required,url"`
	Price    float64 `json:"price" validate:"required,gte=0.0,lte=1000000000.0"`
}

type adHandler struct {
	service services.AdService
}

func NewAdHandler(s services.AdService) *adHandler {
	return &adHandler{service: s}
}

func (h *adHandler) PostAd(c echo.Context) error {

	var req PostAdRequest
	if err := c.Bind(&req); err != nil {
		logger.Logger.Error("failed to bind post ad request")
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request"})
	}

	if err := c.Validate(&req); err != nil {
		logger.Logger.Error("failed to validate post advertisement request body")
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request"})
	}

	if !isValidImageURL(req.ImageURL) {
		logger.Logger.Error("invalid image format")
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid image format. Only jpg, jpeg, png are allowed."})
	}

	userID, ok := c.Get("userID").(int)
	if !ok || userID == 0 {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "invalid or missing token"})
	}

	ad, err := h.service.PostAd(c.Request().Context(), req.Title, req.Text, req.ImageURL, req.Price, userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, ad)
}

func (h *adHandler) GetAdList(c echo.Context) error {
	return nil
}

func isValidImageURL(imageURL string) bool {
	u, err := url.Parse(imageURL)
	if err != nil {
		return false
	}
	path := strings.ToLower(u.Path)
	return strings.HasSuffix(path, ".jpg") || strings.HasSuffix(path, ".jpeg") || strings.HasSuffix(path, ".png")
}
