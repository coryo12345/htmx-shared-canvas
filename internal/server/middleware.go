package server

import (
	"net/http"
	"shared-canvas/cmd/web"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func (s *FiberServer) RegisterFiberMiddleware() {
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
