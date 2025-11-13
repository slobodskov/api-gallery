package repository

import (
	"api-gallery/internal/domain"
	"bytes"
	"database/sql"
	"image"
	"image/jpeg"
	"os"
	"path/filepath"
	"time"

	"github.com/nfnt/resize"
)

type PhotoRepository struct {
	data *photoDB
}

func NewPhotoRepository(db *sql.DB) *PhotoRepository {
	return &PhotoRepository{
		data: newPhotoDB(db),
	}
}

func (p *PhotoRepository) UploadPhoto(file []byte, filename string) (*domain.Photo, error) {
	img, format, err := image.Decode(bytes.NewReader(file))
	if err != nil {
		return nil, err
	}

	// Создаем директории если не существуют
	os.MkdirAll("uploads/original", 0755)
	os.MkdirAll("uploads/preview", 0755)

	// Генерируем уникальное имя файла
	ext := filepath.Ext(filename)
	if ext == "" {
		switch format {
		case "jpeg", "jpg":
			ext = ".jpg"
		case "png":
			ext = ".png"
		default:
			ext = ".jpg"
		}
	}
	baseName := time.Now().Format("20060102150405")
	originalFilename := baseName + ext
	previewFilename := baseName + "_preview.jpg"

	originalPath := filepath.Join("uploads", "original", originalFilename)
	originalFile, err := os.Create(originalPath)
	if err != nil {
		return nil, err
	}
	defer originalFile.Close()

	_, err = originalFile.Write(file)
	if err != nil {
		return nil, err
	}

	// Создаем превью
	preview := resize.Resize(300, 0, img, resize.Lanczos3)
	previewPath := filepath.Join("uploads", "preview", previewFilename)
	previewFile, err := os.Create(previewPath)
	if err != nil {
		return nil, err
	}
	defer previewFile.Close()

	err = jpeg.Encode(previewFile, preview, nil)
	if err != nil {
		return nil, err
	}

	// Получаем размеры изображения
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	// Сохраняем в БД
	photo := &domain.Photo{
		Filename:  filename,
		Original:  originalPath,
		Preview:   previewPath,
		Size:      int64(len(file)),
		Width:     width,
		Height:    height,
		CreatedAt: time.Now(),
	}

	err = p.data.Create(photo)
	if err != nil {
		return nil, err
	}

	response := &domain.Photo{
		ID:        photo.ID,
		Filename:  photo.Filename,
		Preview:   "/api/photos/preview/" + previewFilename,
		Size:      photo.Size,
		Width:     photo.Width,
		Height:    photo.Height,
		CreatedAt: photo.CreatedAt,
	}

	return response, nil
}

func (p *PhotoRepository) GetPhotos() ([]domain.Photo, error) {
	photos, err := p.data.FindAll()
	if err != nil {
		return nil, err
	}

	var responses []domain.Photo
	for _, photo := range photos {
		previewFilename := filepath.Base(photo.Preview)
		response := domain.Photo{
			ID:        photo.ID,
			Filename:  photo.Filename,
			Preview:   "/api/photos/preview/" + previewFilename,
			Size:      photo.Size,
			Width:     photo.Width,
			Height:    photo.Height,
			CreatedAt: photo.CreatedAt,
		}
		responses = append(responses, response)
	}

	return responses, nil
}

func (p *PhotoRepository) DeletePhoto(id int) error {
	photo, err := p.data.FindByID(id)
	if err != nil {
		return err
	}

	os.Remove(photo.Original)
	os.Remove(photo.Preview)

	return p.data.Delete(id)
}
