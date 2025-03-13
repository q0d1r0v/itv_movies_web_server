package services

import (
	"itv_movies_web_server/internal/models"
	"itv_movies_web_server/internal/repositories"

	"github.com/google/uuid"
)

type TransactionService interface {
	LoadTransaction() ([]models.Transaction, error)
	CreateTransaction(transaction *models.Transaction) error
}

type transactionService struct {
	transactionRepo repositories.TransactionRepository
}

func NewTransactionService(transactionRepo repositories.TransactionRepository) TransactionService {
	return &transactionService{transactionRepo}
}

func (s *transactionService) LoadTransaction() ([]models.Transaction, error) {
	return s.transactionRepo.LoadTransactions()
}

func (s *transactionService) CreateTransaction(transaction *models.Transaction) error {
	transaction.ID = uuid.New()
	return s.transactionRepo.CreateTransaction(transaction)
}
