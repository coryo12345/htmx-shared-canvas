package web

import (
	"shared-canvas/cmd/web/components"
	"shared-canvas/internal/canvas"

	"github.com/gofiber/fiber/v2"
)

func RenderHomePage(c *fiber.Ctx, canvas *canvas.Canvas) error {
	c.Set("Content-Type", "text/html")
	canvasComponent := components.Canvas(canvas)
	homepage := Home(canvasComponent)
	err := homepage.Render(c.Context(), c.Response().BodyWriter())
	return err
}
