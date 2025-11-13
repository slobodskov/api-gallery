package usecase

import (
	"api-gallery/internal/domain"
	"api-gallery/internal/repository"
)

type PhotoUseCase struct {
	repo repository.IPhoto
}

func NewPhotoUseCase(repo repository.IPhoto) *PhotoUseCase {
	return &PhotoUseCase{repo: repo}
}

func (uc *PhotoUseCase) UploadPhoto(file []byte, filename string) (*domain.Photo, error) {
	return uc.repo.UploadPhoto(file, filename)
}

func (uc *PhotoUseCase) GetPhotos() ([]domain.Photo, error) {
	return uc.repo.GetPhotos()
}

func (uc *PhotoUseCase) DeletePhoto(id int) error {
	return uc.repo.DeletePhoto(id)
}
