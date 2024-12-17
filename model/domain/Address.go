package domain

import "github.com/google/uuid"

type Address struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID     uuid.UUID `gorm:"type:uuid;not null"`
	Address    string    `gorm:"size:255;not null"`
	City       string    `gorm:"size:100;not null"`
	PostalCode string    `gorm:"size:20;not null"`
	User       User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
