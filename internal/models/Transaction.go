package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	ID            uuid.UUID `gorm:"type:uuid;not null"`
	UserID        uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
	User          User      `gorm:"foreignKey:UserID"`
	Amount        float64   `gorm:"type:decimal(10,2);not null"`
	Status        string    `gorm:"type:varchar(50);not null;default:'pending'"`
	PaymentMethod string    `gorm:"type:varchar(100);not null" json:"payment_method"`
	gorm.Model
}
