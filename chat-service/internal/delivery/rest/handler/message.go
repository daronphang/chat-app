package handler

import (
	"chat-service/internal/domain"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func (h *RestHandler) HandleMediaContent(c echo.Context) error {
	p := new(domain.MediaMessage)
	if err := bindAndValidateRequestBody(c, p); err != nil {
		logger.Error("validation error", zap.String("trace", err.Error()))
		return newRequestValidationError(c, http.StatusBadRequest, err)
	}
	return c.String(http.StatusOK, "media received")
}