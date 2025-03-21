package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role string

const (
	Admin        Role = "admin"
	PlatformUser Role = "user"
)

func (r *Role) GormDataType() string {
	return "user_role"
}

type User struct {
	ID        uuid.UUID  `gorm:"type:uuid;not null;primaryKey"`
	Name      string     `gorm:"type:varchar(255);not null"`
	Email     string     `gorm:"type:varchar(255);uniqueIndex;not null"`
	Password  string     `gorm:"type:text;not null"`
	Role      Role       `gorm:"type:user_role;not null;default:'user'"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

func CreateUserRoleEnum(db *gorm.DB) error {
	return db.Exec(`DO $$ 
	BEGIN 
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_role') THEN 
			CREATE TYPE user_role AS ENUM ('admin', 'user'); 
		END IF; 
	END $$;`).Error
}
