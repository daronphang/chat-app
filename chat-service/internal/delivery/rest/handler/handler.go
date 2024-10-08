package handler

import (
	"chat-service/internal"
	"chat-service/internal/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

var logger, _ = internal.WireLogger()

type RestHandler struct {
	usecase *usecase.UseCaseService
}

func NewRestHandler(uc *usecase.UseCaseService) *RestHandler {
	return &RestHandler{usecase: uc}
}

func (h *RestHandler) Heartbeat(c echo.Context) error {
	return c.String(http.StatusOK, "chat-service is alive")
}
