package usecase

import "api-gallery/internal/domain"

type PhotoUseCase interface {
	UploadPhoto(file []byte, filename string) (*domain.Photo, error)
	GetPhotos() ([]domain.Photo, error)
	DeletePhoto(id int) error
}
