package controllers

import (
	"itv_movies_web_server/internal/models"
	"itv_movies_web_server/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TransactionController struct {
	transactionService services.TransactionService
}

func NewTransactionController(transactionService services.TransactionService) *TransactionController {
	return &TransactionController{transactionService}
}

func (c *TransactionController) LoadTransaction(ctx *gin.Context) {
	transaction, err := c.transactionService.LoadTransaction()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transaction"})
		return
	}
	ctx.JSON(http.StatusOK, transaction)
}

func (c *TransactionController) CreateTransaction(ctx *gin.Context) {
	var transaction models.Transaction
	if err := ctx.ShouldBindJSON(&transaction); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if transaction.UserID == uuid.Nil || transaction.Amount == 0 || transaction.PaymentMethod == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "UserID, Amount, PaymentMethod are required!"})
		return
	}

	err := c.transactionService.CreateTransaction(&transaction)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}
	ctx.JSON(http.StatusCreated, transaction)
}
