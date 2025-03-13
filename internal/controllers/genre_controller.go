package controllers

import (
	"itv_movies_web_server/internal/models"
	"itv_movies_web_server/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GenreController struct {
	genreService services.GenreService
}

func NewGenreController(genreService services.GenreService) *GenreController {
	return &GenreController{genreService}
}

// LoadGenres - Retrieve all genres
//	@Summary		Retrieve all genres
//	@Description	Returns a list of all available genres
//	@Tags			Genre
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.Genre
//	@Failure		500	{object}	map[string]string
//	@Router			/load/genres [get]
func (c *GenreController) LoadGenres(ctx *gin.Context) {
	genres, err := c.genreService.LoadGenres()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch genres"})
		return
	}
	ctx.JSON(http.StatusOK, genres)
}

// CreateGenre - Add a new genre
//	@Summary		Add a new genre
//	@Description	This endpoint is used to add a new genre
//	@Tags			Genre
//	@Accept			json
//	@Produce		json
//	@Param			genre	body		models.Genre	true	"New genre details"
//	@Success		201		{object}	models.Genre
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/create/genre [post]
func (c *GenreController) CreateGenre(ctx *gin.Context) {
	var genre models.Genre
	if err := ctx.ShouldBindJSON(&genre); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if genre.Name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
		return
	}

	err := c.genreService.CreateGenre(&genre)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create genre"})
		return
	}
	ctx.JSON(http.StatusCreated, genre)
}
