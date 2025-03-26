package models

import "time"

type Client struct {
	ID             uint           `gorm:"primaryKey"`
	FirstName      string         `gorm:"not null"`
	SecondName     string         `gorm:"default:null"`
	LastName       string         `gorm:"not null"`
	SecondLastName string         `gorm:"default:null"`
	PhoneNumber    string         `gorm:"not null"`
	Metadata       map[string]any `gorm:"type:jsonb"`
	ServiceOrders  []ServiceOrder `gorm:"foreignKey:ClientID"`
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
}
