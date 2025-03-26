package models

import "time"

type ServiceOrder struct {
	ID        uint   `gorm:"primaryKey"`
	ClientID  uint   `gorm:"not null"`
	Client    Client `gorm:"foreignKey:ClientID"`
	Details   string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Status    int    `gorm:"default:0"`
	Tools     []Tool `gorm:"foreignKey:ServiceOrderID"` // One-to-many
}
