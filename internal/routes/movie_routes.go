package routes

import (
	"itv_movies_web_server/internal/controllers"
	"itv_movies_web_server/internal/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterMovieRoutes(router *gin.Engine, movieController *controllers.MovieController, db *gorm.DB) {
	adminGroup := router.Group("/admin/")
	// userGroup := router.Group("/api/")
	{
		// admin routes
		adminGroup.GET("/load/movies", middlewares.JWTAdminMiddleware(db), movieController.LoadMovies)
		adminGroup.POST("/create/movie", middlewares.JWTAdminMiddleware(db), movieController.CreateMovie)
		adminGroup.PUT("/update/movie", middlewares.JWTAdminMiddleware(db), movieController.UpdateMovie)

		// user routes
		// userGroup.POST("/load/movies", genreController.LoadGenres)
	}
}
