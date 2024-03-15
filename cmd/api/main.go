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
	CANVAS_HEIGHT = 10
	CANVAS_WIDTH  = 10
)

func main() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "prod"
	}

	canvasRepo := canvas.New(CANVAS_HEIGHT, CANVAS_WIDTH)

	server := server.New(canvasRepo)

	server.RegisterFiberMiddleware(env)
	server.RegisterFiberRoutes()

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	err := server.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
