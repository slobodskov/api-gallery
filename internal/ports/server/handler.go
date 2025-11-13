package server

import (
	"api-gallery/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PhotoHandler struct {
	photoUseCase usecase.PhotoUseCase
}

func NewPhotoHandler(uc usecase.PhotoUseCase) *PhotoHandler {
	return &PhotoHandler{photoUseCase: uc}
}

func (h *PhotoHandler) UploadPhoto(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	// Читаем файл
	uploadedFile, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot open file"})
		return
	}
	defer uploadedFile.Close()

	fileBytes := make([]byte, file.Size)
	_, err = uploadedFile.Read(fileBytes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot read file"})
		return
	}

	response, err := h.photoUseCase.UploadPhoto(fileBytes, file.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *PhotoHandler) GetPhotos(c *gin.Context) {
	photos, err := h.photoUseCase.GetAllPhotos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, photos)
}

func (h *PhotoHandler) DeletePhoto(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid photo ID"})
		return
	}

	err = h.photoUseCase.DeletePhoto(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Photo deleted successfully"})
}

func (h *PhotoHandler) ServePreview(c *gin.Context) {
	filename := c.Param("filename")
	filepath := "uploads/preview/" + filename

	c.File(filepath)
}
