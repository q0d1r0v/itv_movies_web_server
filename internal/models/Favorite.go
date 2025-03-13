package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Favorite struct {
	UserID  uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	MovieID uuid.UUID `gorm:"type:uuid;not null" json:"movie_id"`

	User  User  `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Movie Movie `gorm:"foreignKey:MovieID;constraint:OnDelete:CASCADE"`
	gorm.Model
}
