package server

import (
	"net/http"
	"shared-canvas/cmd/web"
	"strings"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func (s *FiberServer) RegisterFiberMiddleware(environment string) {
	if !strings.Contains(strings.ToLower(environment), "prod") {
		s.App.Use(monitor.New(monitor.Config{
			Title: "shared-canvas Metrics",
		}))
	}

	s.App.Use(compress.New())

	s.App.Use("/static", filesystem.New(filesystem.Config{
		Root:       http.FS(web.StaticFiles),
		PathPrefix: "static",
	}))

	s.App.Use(logger.New())
	s.App.Use(recover.New())

	s.App.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
}
