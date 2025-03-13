package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Movie struct {
	ID            uuid.UUID `gorm:"type:uuid;not null;primaryKey"`
	Title         string    `gorm:"type:varchar(255);not null"`
	VideoFilePath string    `gorm:"type:text" json:"video_file_path"`
	Director      string    `gorm:"type:varchar(255);not null"`
	Year          int       `gorm:"type:int;not null"`
	Plot          string    `gorm:"type:text"`
	TrailerURL    string    `gorm:"type:text" json:"trailer_url"`
	Genres        []Genre   `gorm:"many2many:movie_genres;"`
	gorm.Model
}
