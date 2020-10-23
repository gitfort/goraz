package gormext

import (
	"github.com/google/uuid"
	"time"
)

type Model struct {
	ID        uuid.UUID  `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time  `gorm:"not null"`
	UpdatedAt time.Time  `gorm:"not null"`
	DeletedAt *time.Time `gorm:"index"`
}
