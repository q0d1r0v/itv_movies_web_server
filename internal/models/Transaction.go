package models

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID            uuid.UUID  `gorm:"type:uuid;not null"`
	UserID        uuid.UUID  `gorm:"type:uuid;not null;index" json:"user_id"`
	User          User       `gorm:"foreignKey:UserID"`
	Amount        float64    `gorm:"type:decimal(10,2);not null"`
	Status        string     `gorm:"type:varchar(50);not null;default:'pending'"`
	PaymentMethod string     `gorm:"type:varchar(100);not null" json:"payment_method"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}
