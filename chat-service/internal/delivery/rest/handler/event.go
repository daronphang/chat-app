package handler

import (
	"chat-service/internal/domain"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)


func (h *RestHandler) HandleUserPresenceEvent(c echo.Context) error {
	p := new(domain.UserPresenceEvent)
	if err := bindAndValidateRequestBody(c, p); err != nil {
		logger.Error("validation error", zap.String("trace", err.Error()))
		return newRequestValidationError(c, http.StatusBadRequest, err)
	}

	event := domain.BaseEvent{
		Event: domain.EventUserPresence,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Data: *p,
	}

	if err := h.usecase.SendEventToClient(c.Request().Context(), p.ClientID, event); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, "success")
}
