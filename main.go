package main

import (
	"itv_movies_web_server/internal/config"
	"itv_movies_web_server/internal/controllers"
	"itv_movies_web_server/internal/models"
	"itv_movies_web_server/internal/repositories"
	"itv_movies_web_server/internal/routes"
	"itv_movies_web_server/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDatabase()

	db := config.DB
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

	router := gin.Default()

	routes.RegisterUserRoutes(router, userController, db)

	router.Run(":3001")
}
