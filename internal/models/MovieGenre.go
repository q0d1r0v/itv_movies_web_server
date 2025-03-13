package models

import (
	"time"

	"github.com/google/uuid"
)

type MovieGenre struct {
	MovieID   uuid.UUID  `gorm:"type:uuid;not null;primaryKey"`
	GenreID   uuid.UUID  `gorm:"type:uuid;not null;primaryKey"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}
