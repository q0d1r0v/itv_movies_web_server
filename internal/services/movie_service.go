package services

import (
	"itv_movies_web_server/internal/models"
	"itv_movies_web_server/internal/repositories"

	"github.com/google/uuid"
)

type MovieService interface {
	LoadMovies() ([]models.Movie, error)
	CreateMovie(movie *models.Movie) error
	UpdateMovie(movie *models.Movie) error
}

type movieService struct {
	movieRepo repositories.MovieRepository
}

func NewMovieService(movieRepo repositories.MovieRepository) MovieService {
	return &movieService{movieRepo}
}

func (s *movieService) LoadMovies() ([]models.Movie, error) {
	return s.movieRepo.LoadMovies()
}

func (s *movieService) CreateMovie(movie *models.Movie) error {
	movie.ID = uuid.New()
	return s.movieRepo.CreateMovie(movie)
}
func (s *movieService) UpdateMovie(movie *models.Movie) error {
	return s.movieRepo.UpdateMovie(movie)
}
