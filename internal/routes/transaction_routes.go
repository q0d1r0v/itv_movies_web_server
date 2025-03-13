package routes

import (
	"itv_movies_web_server/internal/controllers"
	"itv_movies_web_server/internal/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterTransactionRoutes(router *gin.Engine, transactionController *controllers.TransactionController, db *gorm.DB) {
	adminGroup := router.Group("/admin/")
	// userGroup := router.Group("/api/")
	{
		// admin routes
		adminGroup.GET("/load/transactions", middlewares.JWTAdminMiddleware(db), transactionController.LoadTransaction)
		adminGroup.POST("/create/transaction", middlewares.JWTAdminMiddleware(db), transactionController.CreateTransaction)

		// user routes
		// userGroup.POST("/load/transaction", transactionController.LoadTransaction)
	}
}
