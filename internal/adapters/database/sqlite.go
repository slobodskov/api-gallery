package database

import (
	"database/sql"
	"log/slog"

	_ "modernc.org/sqlite"
)

// InitDB initializes the SQLite database connection and creates necessary tables
// Returns database connection instance or error if initialization fails
func InitDB(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

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

	slog.Info("Database initialized successfully", "path", dbPath)
	return db, nil
}
