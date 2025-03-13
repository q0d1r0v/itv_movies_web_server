package models

import (
	"time"

	"github.com/google/uuid"
)

type Review struct {
	UserID  uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	MovieID uuid.UUID `gorm:"type:uuid;not null" json:"movie_id"`
	Rating  float64   `gorm:"type:decimal(3,1);not null;check:rating >= 0 AND rating <= 10"`
	Comment string    `gorm:"type:text"`

	User      User       `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Movie     Movie      `gorm:"foreignKey:MovieID;constraint:OnDelete:CASCADE"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}
