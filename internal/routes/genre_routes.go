package routes

import (
	"itv_movies_web_server/internal/controllers"
	"itv_movies_web_server/internal/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterGenreRoutes(router *gin.Engine, genreController *controllers.GenreController, db *gorm.DB) {
	adminGroup := router.Group("/admin/")
	// userGroup := router.Group("/api/")
	{
		// admin routes
		adminGroup.POST("/create/genre", middlewares.JWTAdminMiddleware(db), genreController.CreateGenre)

		// user routes
		// userGroup.POST("/load/genres", genreController.LoadGenres)
	}
}
