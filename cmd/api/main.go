package main

import (
	"fmt"
	"os"
	"shared-canvas/internal/canvas"
	"shared-canvas/internal/server"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

const (
	CANVAS_HEIGHT = 32
	CANVAS_WIDTH  = 32
)

func main() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "prod"
	}

	canvasRepo := canvas.New(CANVAS_HEIGHT, CANVAS_WIDTH)

	server := server.New(canvasRepo)

	server.RegisterFiberMiddleware()
	server.RegisterFiberRoutes(env)

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	err := server.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
