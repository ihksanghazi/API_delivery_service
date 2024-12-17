package domain

import (
	"time"

	"github.com/google/uuid"
)

type Pricing struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Origin      string    `gorm:"size:100;not null"`
	Destination string    `gorm:"size:100;not null"`
	Weight      float64   `gorm:"not null"`
	ServiceType string    `gorm:"size:50;not null"`
	Price       float64   `gorm:"not null"`
}

type Payment struct {
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ShipmentID    uuid.UUID `gorm:"type:uuid;not null"`
	Amount        float64   `gorm:"not null"`
	PaymentStatus string    `gorm:"size:50;not null"`
	PaidAt        *time.Time
	Shipment      Shipment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
