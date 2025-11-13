package repository

import "api-gallery/internal/domain"

type IPhoto interface {
	UploadPhoto(file []byte, filename string) (*domain.Photo, error)
	GetPhotos() ([]domain.Photo, error)
	DeletePhoto(id int) error
}
