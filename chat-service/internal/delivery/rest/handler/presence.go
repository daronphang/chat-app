package handler

import (
	"chat-service/internal/domain"
	"net/http"

	cv "chat-service/internal/validator"

	"github.com/labstack/echo/v4"
)


func (h *RestHandler) BroadcastUserStatus(c echo.Context) error {
	clientId := c.QueryParam("clientId")
	status := c.QueryParam("status")
	targetId := c.QueryParam("targetId")

	arg := domain.PresenceStatus{
		ClientID: clientId,
		Status: status,
		TargetID: targetId,
	}
	if err := cv.ProvideValidator().Validate(arg); err != nil {
		return newRequestValidationError(c, http.StatusBadRequest, err)
	}

	if err := h.UseCase.BroadcastPresenceStatus(c.Request().Context(), arg); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, "broadcast success")
}