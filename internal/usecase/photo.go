package usecase

import (
	"api-gallery/internal/domain"
	"api-gallery/internal/repository"
)

type photoUseCase struct {
	repo PhotoUseCase
}

func NewPhotoUseCase(repo *repository.PhotoRepository) *photoUseCase {
	return &photoUseCase{repo: repo}
}

func (uc *photoUseCase) UploadPhoto(file []byte, filename string) (*domain.Photo, error) {
	return uc.repo.UploadPhoto(file, filename)
}

func (uc *photoUseCase) GetPhotos() ([]domain.Photo, error) {
	return uc.repo.GetPhotos()
}

func (uc *photoUseCase) DeletePhoto(id int) error {
	return uc.repo.DeletePhoto(id)
}
