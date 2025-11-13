package app

import (
	"api-gallery/internal/adapters/database"
	"api-gallery/internal/ports/server"
	"api-gallery/internal/repository"
	"api-gallery/internal/usecase"
	"log"
)

func Run() error {

	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	photoRepo := repository.NewPhotoRepository(db)
	photoUC := usecase.NewPhotoUseCase(photoRepo)

	// Настройка роутера
	router := server.SetupRouter(photoUC)

	// Запуск сервера
	log.Println("Server starting on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
	return nil
}
