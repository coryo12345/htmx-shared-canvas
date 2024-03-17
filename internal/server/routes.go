package server

import (
	"bytes"
	"context"
	"encoding/json"
	"shared-canvas/cmd/web"
	"shared-canvas/cmd/web/components"
	"shared-canvas/internal/canvas"
	"shared-canvas/internal/utils"
	"strconv"
	"strings"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func (s *FiberServer) RegisterFiberRoutes(environment string) {
	s.App.Get("/", func(c *fiber.Ctx) error {
		return web.RenderHomePage(c, s.Canvas)
	})

	s.App.Get("/clear", func(c *fiber.Ctx) error {
		s.Canvas.Clear()
		c.RedirectBack("/")
		return nil
	})

	s.App.Get("/ws", websocket.New(s.websocketUpdateHandler))

	if !strings.Contains(strings.ToLower(environment), "prod") {
		s.App.Use(monitor.New(monitor.Config{
			Title: "shared-canvas Metrics",
		}))
	}
}

type updateData struct {
	Color string `json:"color"`
	Pos   string `json:"pos"`
}

func (s *FiberServer) websocketUpdateHandler(c *websocket.Conn) {
	key := s.AddConnection(c)
	defer func() {
		s.RemoveConnection(key)
		c.Close()
	}()

	var (
		mt  int
		msg []byte
		err error
	)
	for {
		if mt, msg, err = c.ReadMessage(); err != nil {
			log.Infof("closing socket %d\n", key)
			break
		}

		// read updates
		data := updateData{}
		if err = json.Unmarshal(msg, &data); err != nil {
			log.Infof("failed to parse websocket data")
			break
		}

		// convert data formats
		pos, err := strconv.Atoi(data.Pos)
		if err != nil {
			break
		}
		rgb, err := utils.HexColor(data.Color)
		if err != nil {
			break
		}

		// update model
		err = s.Canvas.SetPixel(pos, canvas.Pixel{
			R: rgb.R,
			G: rgb.G,
			B: rgb.B,
		})
		if err != nil {
			break
		}

		// write update to all connections
		newPixel := s.Canvas.GetPixel(pos)
		pixelComponent := components.CanvasPixel(pos, newPixel)
		buf := new(bytes.Buffer)
		err = pixelComponent.Render(context.Background(), buf)
		if err != nil {
			break
		}

		for i, conn := range s.connections {
			if conn == nil {
				continue
			}
			if err = conn.WriteMessage(mt, buf.Bytes()); err != nil {
				log.Infof("failed to write new pixel to connection %d \n", i)
			}
		}
	}
}
