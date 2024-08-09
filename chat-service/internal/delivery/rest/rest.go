package rest

import (
	api "chat-service/internal/delivery/rest/api/v1"
	rh "chat-service/internal/delivery/rest/handler"
	cm "chat-service/internal/delivery/rest/middleware"
	"net/http"

	uc "chat-service/internal/usecase"
	cv "chat-service/internal/validator"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type RestServer struct {
	Echo *echo.Echo
}

func New(logger *zap.Logger, uc *uc.UseCaseService, wsHandler func(w http.ResponseWriter, r *http.Request)) *RestServer {
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

	eventGroup := baseGroup.Group("/event")
	api.RegisterEventRoutes(eventGroup, rh)

	// Register websocket.
	baseGroup.GET("/ws", func(c echo.Context) error {
		wsHandler(c.Response().Writer, c.Request())
		return nil
	})

	return &RestServer{
		Echo: e,
	}
}