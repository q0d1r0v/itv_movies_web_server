package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MovieGenre struct {
	MovieID uuid.UUID `gorm:"type:uuid;not null;primaryKey"`
	GenreID uuid.UUID `gorm:"type:uuid;not null;primaryKey"`
	gorm.Model
}
