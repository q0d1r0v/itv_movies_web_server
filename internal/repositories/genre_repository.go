package repositories

import (
	"itv_movies_web_server/internal/models"

	"gorm.io/gorm"
)

type GenreRepository interface {
	LoadGenres() ([]models.Genre, error)
	CreateGenre(genre *models.Genre) error
}

type genreRepository struct {
	db *gorm.DB
}

func NewGenreRepository(db *gorm.DB) GenreRepository {
	return &genreRepository{db}
}

func (r *genreRepository) LoadGenres() ([]models.Genre, error) {
	var genres []models.Genre
	err := r.db.Find(&genres).Error
	return genres, err
}

func (r *genreRepository) CreateGenre(genre *models.Genre) error {
	return r.db.Create(genre).Error
}
