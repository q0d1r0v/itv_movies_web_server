package main

import (
	_ "itv_movies_web_server/docs"
	"itv_movies_web_server/internal/config"
	"itv_movies_web_server/internal/controllers"
	"itv_movies_web_server/internal/models"
	"itv_movies_web_server/internal/repositories"
	"itv_movies_web_server/internal/routes"
	"itv_movies_web_server/internal/services"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			ITV Movies Web Server API
//	@version		1.0
//	@description	API Docs for ITV Movies Web Server.
//	@host			::host
//	@BasePath		/
func main() {
	config.InitDatabase()

	db := config.DB
	if err := models.CreateUserRoleEnum(db); err != nil {
		log.Fatal("Failed to create ENUM type:", err)
	}
	err := config.DB.AutoMigrate(
		&models.Favorite{},
		&models.Genre{},
		&models.Movie{},
		&models.MovieGenre{},
		&models.Review{},
		&models.User{},
		&models.Transaction{},
	)
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	// data of user
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	// data of genre
	genreRepo := repositories.NewGenreRepository(db)
	genreService := services.NewGenreService(genreRepo)
	genreController := controllers.NewGenreController(genreService)

	// data of movie
	movieRepo := repositories.NewMovieRepository(db)
	movieService := services.NewMovieService(movieRepo)
	movieController := controllers.NewMovieController(movieService)

	// data of transaction
	transactionRepo := repositories.NewTransactionRepository(db)
	transactionService := services.NewTransactionService(transactionRepo)
	transactionController := controllers.NewTransactionController(transactionService)

	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.RegisterUserRoutes(router, userController, db)
	routes.RegisterGenreRoutes(router, genreController, db)
	routes.RegisterMovieRoutes(router, movieController, db)
	routes.RegisterTransactionRoutes(router, transactionController, db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	router.Run(":" + port)
}
