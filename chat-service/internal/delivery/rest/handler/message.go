package handler

import (
	"chat-service/internal/domain"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func (h *RestHandler) SendMsgToClient(c echo.Context) error {
	p := new(domain.ReceiverMessage)
	if err := bindAndValidateRequestBody(c, p); err != nil {
		logger.Error("validation error", zap.String("trace", err.Error()))
		return newRequestValidationError(c, http.StatusBadRequest, err)
	}
	if err := h.UseCase.EventBroker.PublishMessage(c.Request().Context(), p.Message.ChannelID, p.ReceiverID, p.Message); err != nil {
		logger.Error("error sending message to client", zap.String("trace", err.Error()))
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "message received")
}
