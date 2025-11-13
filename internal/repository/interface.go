package repository

import "api-gallery/internal/domain"

type PhotoRepository interface {
	Create(photo *domain.Photo) error
	FindAll() ([]domain.Photo, error)
	FindByID(id int) (*domain.Photo, error)
	Delete(id int) error
}
