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

// LoadTransactions - Retrieve all transactions
//	@Summary		Retrieve all transactions
//	@Description	This endpoint returns a list of all transactions
//	@Tags			Transactions
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.Transaction
//	@Failure		500	{object}	map[string]string
//	@Router			/load/transactions [get]
func (c *TransactionController) LoadTransaction(ctx *gin.Context) {
	transaction, err := c.transactionService.LoadTransaction()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transaction"})
		return
	}
	ctx.JSON(http.StatusOK, transaction)
}

// CreateTransaction - Create a new transaction
//	@Summary		Create a new transaction
//	@Description	This endpoint is used to create a new transaction
//	@Tags			Transaction
//	@Accept			json
//	@Produce		json
//	@Param			transaction	body		models.Transaction	true	"Transaction details"
//	@Success		201			{object}	models.Transaction
//	@Failure		400			{object}	map[string]string
//	@Failure		500			{object}	map[string]string
//	@Router			/create/transactions [post]
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
