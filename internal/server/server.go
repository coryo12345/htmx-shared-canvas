package server

import (
	"shared-canvas/internal/canvas"
	"sync"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type FiberServer struct {
	*fiber.App
	SessionStore   *session.Store
	Canvas         *canvas.Canvas
	connections    []*websocket.Conn
	connectionLock *sync.Mutex
}

func New(can *canvas.Canvas) *FiberServer {
	server := &FiberServer{
		App:            fiber.New(),
		SessionStore:   session.New(),
		Canvas:         can,
		connections:    make([]*websocket.Conn, 0),
		connectionLock: &sync.Mutex{},
	}

	return server
}

func (s *FiberServer) AddConnection(c *websocket.Conn) int {
	s.connectionLock.Lock()
	s.connections = append(s.connections, c)
	key := len(s.connections) - 1 // the index that was just added
	s.connectionLock.Unlock()
	return key
}

func (s *FiberServer) RemoveConnection(key int) {
	s.connectionLock.Lock()
	s.connections[key] = nil
	s.connectionLock.Unlock()
}
