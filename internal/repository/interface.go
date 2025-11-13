package repository

import "api-gallery/internal/domain"

// IPhoto defines the interface for photo repository operations
type IPhoto interface {
	UploadPhoto(file []byte, filename string) (*domain.Photo, error)
	GetPhotos() ([]domain.Photo, error)
	DeletePhoto(id int) error
}
