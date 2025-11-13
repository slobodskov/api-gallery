package repository

import (
	"api-gallery/internal/domain"
	"database/sql"
)

type photoRepository struct {
	db *sql.DB
}

func NewPhotoRepository(db *sql.DB) PhotoRepository {
	return &photoRepository{db: db}
}

func (r *photoRepository) Create(photo *domain.Photo) error {
	query := `INSERT INTO photos (filename, original_path, preview_path, size, width, height) 
              VALUES (?, ?, ?, ?, ?, ?)`
	result, err := r.db.Exec(query, photo.Filename, photo.Original, photo.Preview,
		photo.Size, photo.Width, photo.Height)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	photo.ID = int(id)
	return nil
}

func (r *photoRepository) FindAll() ([]domain.Photo, error) {
	query := `SELECT id, filename, original_path, preview_path, size, width, height, created_at 
              FROM photos ORDER BY created_at DESC`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var photos []domain.Photo
	for rows.Next() {
		var photo domain.Photo
		err := rows.Scan(&photo.ID, &photo.Filename, &photo.Original, &photo.Preview,
			&photo.Size, &photo.Width, &photo.Height, &photo.CreatedAt)
		if err != nil {
			return nil, err
		}
		photos = append(photos, photo)
	}
	return photos, nil
}

func (r *photoRepository) FindByID(id int) (*domain.Photo, error) {
	query := `SELECT id, filename, original_path, preview_path, size, width, height, created_at 
              FROM photos WHERE id = ?`
	row := r.db.QueryRow(query, id)

	var photo domain.Photo
	err := row.Scan(&photo.ID, &photo.Filename, &photo.Original, &photo.Preview,
		&photo.Size, &photo.Width, &photo.Height, &photo.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &photo, nil
}

func (r *photoRepository) Delete(id int) error {
	query := "DELETE FROM photos WHERE id = ?"
	_, err := r.db.Exec(query, id)
	return err
}
