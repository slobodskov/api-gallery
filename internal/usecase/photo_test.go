package usecase

import (
	"api-gallery/internal/domain"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock репозитория
type MockPhotoRepository struct {
	mock.Mock
}

func (m *MockPhotoRepository) UploadPhoto(file []byte, filename string) (*domain.Photo, error) {
	args := m.Called(file, filename)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Photo), args.Error(1)
}

func (m *MockPhotoRepository) GetPhotos() ([]domain.Photo, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]domain.Photo), args.Error(1)
}

func (m *MockPhotoRepository) DeletePhoto(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestUploadPhoto(t *testing.T) {
	mockRepo := new(MockPhotoRepository)
	uc := NewPhotoUseCase(mockRepo)

	file := []byte("fake image data")
	filename := "test.jpg"
	expectedPhoto := &domain.Photo{
		ID:        1,
		Filename:  filename,
		Original:  "uploads/original/test.jpg",
		Preview:   "uploads/preview/test_preview.jpg",
		Size:      12345,
		Width:     800,
		Height:    600,
		CreatedAt: time.Now(),
	}

	mockRepo.On("UploadPhoto", file, filename).Return(expectedPhoto, nil)

	photo, err := uc.UploadPhoto(file, filename)

	assert.NoError(t, err)
	assert.Equal(t, expectedPhoto, photo)
	mockRepo.AssertExpectations(t)
}

func TestUploadPhoto_Error(t *testing.T) {
	mockRepo := new(MockPhotoRepository)
	uc := NewPhotoUseCase(mockRepo)

	file := []byte("fake image data")
	filename := "test.jpg"
	expectedError := errors.New("upload failed")

	mockRepo.On("UploadPhoto", file, filename).Return(nil, expectedError)

	photo, err := uc.UploadPhoto(file, filename)

	assert.Error(t, err)
	assert.Nil(t, photo)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

func TestGetPhotos(t *testing.T) {
	mockRepo := new(MockPhotoRepository)
	uc := NewPhotoUseCase(mockRepo)

	expectedPhotos := []domain.Photo{
		{
			ID:        1,
			Filename:  "test1.jpg",
			Original:  "uploads/original/test1.jpg",
			Preview:   "uploads/preview/test1_preview.jpg",
			Size:      12345,
			Width:     800,
			Height:    600,
			CreatedAt: time.Now(),
		},
		{
			ID:        2,
			Filename:  "test2.jpg",
			Original:  "uploads/original/test2.jpg",
			Preview:   "uploads/preview/test2_preview.jpg",
			Size:      54321,
			Width:     1024,
			Height:    768,
			CreatedAt: time.Now(),
		},
	}

	mockRepo.On("GetPhotos").Return(expectedPhotos, nil)

	photos, err := uc.GetPhotos()

	assert.NoError(t, err)
	assert.Equal(t, expectedPhotos, photos)
	mockRepo.AssertExpectations(t)
}

func TestDeletePhoto(t *testing.T) {
	mockRepo := new(MockPhotoRepository)
	uc := NewPhotoUseCase(mockRepo)

	id := 1
	mockRepo.On("DeletePhoto", id).Return(nil)

	err := uc.DeletePhoto(id)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
