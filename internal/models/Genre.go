package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Genre struct {
	ID     uuid.UUID `gorm:"type:uuid;not null;primaryKey"`
	Name   string    `gorm:"type:varchar(255);not null;unique"`
	Movies []Movie   `gorm:"many2many:movie_genres;"`
	gorm.Model
}
