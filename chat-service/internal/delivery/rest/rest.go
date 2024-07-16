package rest

import (
	api "chat-service/internal/delivery/rest/api/v1"
	rh "chat-service/internal/delivery/rest/handler"
	cm "chat-service/internal/delivery/rest/middleware"
	ws "chat-service/internal/delivery/websocket"
	"context"

	uc "chat-service/internal/usecase"
	cv "chat-service/internal/validator"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type Server struct {
	Echo *echo.Echo
}

func New(ctx context.Context, logger *zap.Logger, uc *uc.UseCaseService, hub *ws.Hub) *Server {
	// Create server.
	e := echo.New()

	// Register middlewares.
	e.Use(
		middleware.CORS(),
		cm.CustomRequestLogger(logger),
	)

	// Register custom handlers.
	e.Validator = cv.ProvideValidator()

	// Create handler.
	rh := rh.NewRestHandler(uc)

	// Register routes.
	baseGroup := e.Group("/api/v1")
	api.RegisterBaseRoutes(baseGroup, rh)

	msgGroup := baseGroup.Group("/message")
	api.RegisterMessageRoutes(msgGroup, rh)

	// Register websocket.
	baseGroup.GET("/ws", func(c echo.Context) error {
		ws.ServeWs(hub, c.Response().Writer, c.Request())
		return nil
	})

	return &Server{
		Echo: e,
	}
}