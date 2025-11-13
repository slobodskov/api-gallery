package usecase

import (
	"api-gallery/internal/domain"
	"api-gallery/internal/repository"
)

// PhotoUseCase provides business logic for photo operations
type PhotoUseCase struct {
	repo repository.IPhoto
}

// NewPhotoUseCase creates a new instance of PhotoUseCase
func NewPhotoUseCase(repo repository.IPhoto) *PhotoUseCase {
	return &PhotoUseCase{repo: repo}
}

// UploadPhoto handles the business logic for uploading a new photo
func (uc *PhotoUseCase) UploadPhoto(file []byte, filename string) (*domain.Photo, error) {
	return uc.repo.UploadPhoto(file, filename)
}

// GetPhotos retrieves all photos from the repository
func (uc *PhotoUseCase) GetPhotos() ([]domain.Photo, error) {
	return uc.repo.GetPhotos()
}

// DeletePhoto handles the business logic for deleting a photo by ID
func (uc *PhotoUseCase) DeletePhoto(id int) error {
	return uc.repo.DeletePhoto(id)
}
