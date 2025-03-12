package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Favorite struct {
	UserID  uuid.UUID `gorm:"type:uuid;not null"`
	MovieID uuid.UUID `gorm:"type:uuid;not null"`

	User  User  `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Movie Movie `gorm:"foreignKey:MovieID;constraint:OnDelete:CASCADE"`
	gorm.Model
}
