package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./gallery.db")
	if err != nil {
		return nil, err
	}

	// Создаем таблицу
	query := `
    CREATE TABLE IF NOT EXISTS photos (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        filename TEXT NOT NULL,
        original_path TEXT NOT NULL,
        preview_path TEXT NOT NULL,
        size INTEGER NOT NULL,
        width INTEGER NOT NULL,
        height INTEGER NOT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP
    )
    `
	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}

	log.Println("Database initialized successfully")
	return db, nil
}
