package server

import (
	"api-gallery/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// PhotoHandler handles HTTP requests for photo operations
type PhotoHandler struct {
	photoUseCase usecase.PhotoUseCase
}

// NewPhotoHandler creates a new instance of PhotoHandler
func NewPhotoHandler(uc usecase.PhotoUseCase) *PhotoHandler {
	return &PhotoHandler{photoUseCase: uc}
}

// UploadPhoto godoc
// @Summary Upload a photo
// @Description Upload a photo and create a preview
// @Tags photos
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Photo file"
// @Success 200 {object} domain.Photo
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/photos [post]
func (h *PhotoHandler) UploadPhoto(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

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

// GetPhotos godoc
// @Summary Get all photos
// @Description Get all photos with pagination
// @Tags photos
// @Produce json
// @Success 200 {array} domain.Photo
// @Failure 500 {object} map[string]string
// @Router /api/photos [get]
func (h *PhotoHandler) GetPhotos(c *gin.Context) {
	photos, err := h.photoUseCase.GetPhotos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, photos)
}

// DeletePhoto godoc
// @Summary Delete photo
// @Description Delete photo by ID
// @Tags photos
// @Produce json
// @Param id path int true "Photo ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /photos/{id} [delete]
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

// ServePreview godoc
// @Summary Get photo preview
// @Description Returns preview image
// @Tags photos
// @Produce image/jpeg
// @Param filename path string true "Preview filename"
// @Success 200
// @Failure 404 {object} map[string]string
// @Router /photos/preview/{filename} [get]
func (h *PhotoHandler) ServePreview(c *gin.Context) {
	filename := c.Param("filename")
	filepath := "uploads/preview/" + filename

	c.File(filepath)
}
