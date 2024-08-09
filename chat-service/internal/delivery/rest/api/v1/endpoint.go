package v1

import (
	"chat-service/internal/delivery/rest/handler"

	"github.com/labstack/echo/v4"
)

func RegisterBaseRoutes(g *echo.Group, h *handler.RestHandler) {
	g.GET("/heartbeat", h.Heartbeat)
}

func RegisterMessageRoutes(g *echo.Group, h *handler.RestHandler) {
	g.POST("", h.SendMsgToClient)
	g.POST("/media", h.UploadMedia)
}

func RegisterEventRoutes(g *echo.Group, h *handler.RestHandler) {
	g.GET("/presence", h.HandleUserPresenceEvent)
}
