package repository

import (
	"api-gallery/internal/domain"
	"database/sql"
)

// photoDB provides database operations for photos
type photoDB struct {
	db *sql.DB
}

// newPhotoDB creates a new instance of photoDB
func newPhotoDB(db *sql.DB) *photoDB {
	return &photoDB{db: db}
}

// Create inserts a new photo record into the database
func (p *photoDB) Create(photo *domain.Photo) error {
	query := `INSERT INTO photos (filename, original_path, preview_path, size, width, height) 
              VALUES (?, ?, ?, ?, ?, ?)`
	result, err := p.db.Exec(query, photo.Filename, photo.Original, photo.Preview,
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

// FindAll retrieves all photos from the database ordered by creation date
func (p *photoDB) FindAll() ([]domain.Photo, error) {
	query := `SELECT id, filename, original_path, preview_path, size, width, height, created_at 
              FROM photos ORDER BY created_at DESC`
	rows, err := p.db.Query(query)
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

// FindByID retrieves a specific photo by its ID
func (p *photoDB) FindByID(id int) (*domain.Photo, error) {
	query := `SELECT id, filename, original_path, preview_path, size, width, height, created_at 
              FROM photos WHERE id = ?`
	row := p.db.QueryRow(query, id)

	var photo domain.Photo
	err := row.Scan(&photo.ID, &photo.Filename, &photo.Original, &photo.Preview,
		&photo.Size, &photo.Width, &photo.Height, &photo.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &photo, nil
}

// Delete removes a photo from the database by ID
func (p *photoDB) Delete(id int) error {
	query := "DELETE FROM photos WHERE id = ?"
	_, err := p.db.Exec(query, id)
	return err
}
