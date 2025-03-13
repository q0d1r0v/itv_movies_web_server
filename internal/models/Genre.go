package models

import (
	"time"

	"github.com/google/uuid"
)

type Genre struct {
	ID        uuid.UUID  `gorm:"type:uuid;not null;primaryKey"`
	Name      string     `gorm:"type:varchar(255);not null;unique"`
	Movies    []Movie    `gorm:"many2many:movie_genres;"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}
