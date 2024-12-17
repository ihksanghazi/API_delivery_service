package domain

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name      string     `gorm:"size:255;not null"`
	Email     string     `gorm:"size:255;unique;not null"`
	Password  string     `gorm:"not null"`
	Role      string     `gorm:"size:50;not null"`
	Addresses []Address  `gorm:"foreignKey:UserID"`
	Shipments []Shipment `gorm:"foreignKey:CustomerID"`
}
