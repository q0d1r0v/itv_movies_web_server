package routes

import (
	"itv_movies_web_server/internal/controllers"
	"itv_movies_web_server/internal/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUserRoutes(router *gin.Engine, userController *controllers.UserController, db *gorm.DB) {
	authGroup := router.Group("/auth/")
	adminGroup := router.Group("/admin/")
	// userGroup := router.Group("/api/")
	{
		// admin routes
		adminGroup.GET("/load/users", middlewares.JWTAdminMiddleware(db), userController.GetUsers)

		// auth routes
		authGroup.POST("/register", userController.CreateUser)
		authGroup.POST("/login", userController.Login)
	}
}
