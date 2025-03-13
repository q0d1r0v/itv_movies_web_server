package repositories

import (
	"itv_movies_web_server/internal/models"

	"gorm.io/gorm"
)

type MovieRepository interface {
	LoadMovies() ([]models.Movie, error)
	CreateMovie(movie *models.Movie) error
	UpdateMovie(movie *models.Movie) error
}

type movieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) MovieRepository {
	return &movieRepository{db}
}

func (r *movieRepository) LoadMovies() ([]models.Movie, error) {
	var movies []models.Movie
	err := r.db.Find(&movies).Error
	return movies, err
}

func (r *movieRepository) CreateMovie(movie *models.Movie) error {
	return r.db.Create(movie).Error
}

func (r *movieRepository) UpdateMovie(movie *models.Movie) error {
	return r.db.Model(&models.Movie{}).Where("id = ?", movie.ID).Updates(movie).Error
}
