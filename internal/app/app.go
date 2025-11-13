package app

import (
	"api-gallery/config"
	"api-gallery/internal/adapters/database"
	"api-gallery/internal/infrastructure/logger"
	"api-gallery/internal/ports/server"
	"api-gallery/internal/repository"
	"api-gallery/internal/usecase"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

// Run initializes and starts the application
// It sets up configuration, database, use cases, and HTTP server
// Returns error if any initialization step fails
func Run() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	log := logger.New()

	db, err := database.InitDB(cfg.DatabasePath)
	if err != nil {
		log.Error("Failed to initialize database", "error", err)
		return err
	}
	defer db.Close()

	photoRepo := repository.NewPhotoRepository(db)
	photoUC := usecase.NewPhotoUseCase(photoRepo)

	if cfg.GinMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := server.SetupRouter(*photoUC)

	srv := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: router,
	}

	go func() {
		log.Info("Server starting", "port", cfg.ServerPort, "mode", cfg.GinMode)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("Failed to start server", "error", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error("Server forced to shutdown", "error", err)
		return err
	}

	log.Info("Server exited successfully")
	return nil
}
