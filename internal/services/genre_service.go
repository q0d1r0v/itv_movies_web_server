package services

import (
	"itv_movies_web_server/internal/models"
	"itv_movies_web_server/internal/repositories"

	"github.com/google/uuid"
)

type GenreService interface {
	LoadGenres() ([]models.Genre, error)
	CreateGenre(genre *models.Genre) error
}

type genreService struct {
	genreRepo repositories.GenreRepository
}

func NewGenreService(genreRepo repositories.GenreRepository) GenreService {
	return &genreService{genreRepo}
}

func (s *genreService) LoadGenres() ([]models.Genre, error) {
	return s.genreRepo.LoadGenres()
}

func (s *genreService) CreateGenre(genre *models.Genre) error {
	genre.ID = uuid.New()
	return s.genreRepo.CreateGenre(genre)
}
