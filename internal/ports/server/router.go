package server

import (
	"api-gallery/internal/usecase"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(photoUC usecase.PhotoUseCase) *gin.Engine {
	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	photoHandler := NewPhotoHandler(photoUC)

	api := router.Group("/api")
	{
		photos := api.Group("/photos")
		{
			photos.POST("", photoHandler.UploadPhoto)
			photos.GET("", photoHandler.GetPhotos)
			photos.DELETE("/:id", photoHandler.DeletePhoto)
			photos.GET("/preview/:filename", photoHandler.ServePreview)
		}
	}

	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
