package domain

import "time"

type Photo struct {
	ID        int       `json:"id" db:"id"`
	Filename  string    `json:"filename" db:"filename"`
	Original  string    `json:"original_path" db:"original_path"`
	Preview   string    `json:"preview_path" db:"preview_path"`
	Size      int64     `json:"size" db:"size"`
	Width     int       `json:"width" db:"width"`
	Height    int       `json:"height" db:"height"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type PhotoUploadRequest struct {
	File []byte `json:"file" binding:"required"`
}
