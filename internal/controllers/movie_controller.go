package controllers

import (
	"fmt"
	"itv_movies_web_server/internal/models"
	"itv_movies_web_server/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MovieController struct {
	movieService services.MovieService
}

func NewMovieController(movieService services.MovieService) *MovieController {
	return &MovieController{movieService}
}

// LoadMovies - Retrieve all movies
//	@Summary		Retrieve all movies
//	@Description	Returns a list of all available movies
//	@Tags			Movie
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.Movie
//	@Failure		500	{object}	map[string]string
//	@Router			/load/movies [get]
func (c *MovieController) LoadMovies(ctx *gin.Context) {
	movies, err := c.movieService.LoadMovies()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch movies"})
		return
	}
	ctx.JSON(http.StatusOK, movies)
}

// CreateMovie - Add a new movie
//	@Summary		Add a new movie
//	@Description	This endpoint is used to add a new movie
//	@Tags			Movie
//	@Accept			json
//	@Produce		json
//	@Param			movie	body		models.Movie	true	"New movie details"
//	@Success		201		{object}	models.Movie
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/create/movie [post]
func (c *MovieController) CreateMovie(ctx *gin.Context) {
	var movie models.Movie
	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	fmt.Printf("Parsed movie: %+v\n", movie)

	if movie.Title == "" || movie.VideoFilePath == "" || movie.Director == "" || movie.Year == 0 || movie.Plot == "" || movie.TrailerURL == "" || len(movie.Genres) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Genres, Title, VideoFilePath, Director, Year(int), Plot and TrailerURL are required!"})
		return
	}

	err := c.movieService.CreateMovie(&movie)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create movie"})
		return
	}
	ctx.JSON(http.StatusCreated, movie)
}

// UpdateMovie - Update an existing movie
//	@Summary		Update an existing movie
//	@Description	This endpoint updates the details of an existing movie
//	@Tags			Movie
//	@Accept			json
//	@Produce		json
//	@Param			movie	body		models.Movie	true	"Updated movie details"
//	@Success		200		{object}	map[string]string
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/update/movie [put]
func (c *MovieController) UpdateMovie(ctx *gin.Context) {
	var movie models.Movie
	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if movie.ID == uuid.Nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Movie ID is required!"})
		return
	}

	err := c.movieService.UpdateMovie(&movie)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update movie"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Movie updated successfully"})
}
