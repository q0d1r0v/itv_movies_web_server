package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Review struct {
	UserID  uuid.UUID `gorm:"type:uuid;not null"`
	MovieID uuid.UUID `gorm:"type:uuid;not null"`
	Rating  float64   `gorm:"type:decimal(3,1);not null;check:rating >= 0 AND rating <= 10"`
	Comment string    `gorm:"type:text"`

	User  User  `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Movie Movie `gorm:"foreignKey:MovieID;constraint:OnDelete:CASCADE"`
	gorm.Model
}
