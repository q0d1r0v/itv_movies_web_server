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

func (c *MovieController) LoadMovies(ctx *gin.Context) {
	movies, err := c.movieService.LoadMovies()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch movies"})
		return
	}
	ctx.JSON(http.StatusOK, movies)
}

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
