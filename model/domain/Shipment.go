package domain

import (
	"time"

	"github.com/google/uuid"
)

type Shipment struct {
	ID                   uuid.UUID        `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	CustomerID           uuid.UUID        `gorm:"type:uuid;not null"`
	OriginAddressID      uuid.UUID        `gorm:"type:uuid;not null"`
	DestinationAddressID uuid.UUID        `gorm:"type:uuid;not null"`
	StatusID             uuid.UUID        `gorm:"type:uuid;not null"`
	ServiceType          string           `gorm:"size:50;not null"`
	CreatedAt            time.Time        `gorm:"autoCreateTime"`
	Customer             User             `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	OriginAddress        Address          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DestinationAddress   Address          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Status               string           `gorm:"size:100;not null"`
	Details              []ShipmentDetail `gorm:"foreignKey:ShipmentID"`
}

type ShipmentDetail struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ShipmentID  uuid.UUID `gorm:"type:uuid;not null"`
	Weight      float64   `gorm:"not null"`
	Description string    `gorm:"size:255"`
	Shipment    Shipment  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
