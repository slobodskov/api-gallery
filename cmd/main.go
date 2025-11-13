package main

import (
	"api-gallery/internal/app"
	"log/slog"
	"os"
)

// @title Photo Gallery API
// @version 1.0
// @description This is a sample server for a photo gallery.

// @host localhost:8080
// @BasePath /api
func main() {
	if err := app.Run(); err != nil {
		slog.Error("Application failed to start", "error", err)
		os.Exit(1)
	}
}
