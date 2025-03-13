package repositories

import (
	"itv_movies_web_server/internal/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	LoadTransactions() ([]models.Transaction, error)
	CreateTransaction(transaction *models.Transaction) error
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) LoadTransactions() ([]models.Transaction, error) {
	var transaction []models.Transaction
	err := r.db.Find(&transaction).Error
	return transaction, err
}

func (r *transactionRepository) CreateTransaction(transaction *models.Transaction) error {
	// create transaction
	tx := r.db.Begin()

	if err := tx.Create(transaction).Error; err != nil {
		tx.Rollback()
		return err
	}

	// check another ballance or actions
	// some actions

	// commit
	tx.Commit()
	return nil
}
