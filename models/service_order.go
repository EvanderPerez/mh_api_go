package models

import "time"

type ServiceOrder struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	ClientID  uint       `json:"-" gorm:"not null"`
	Client    Client     `json:"client" gorm:"foreignKey:ClientID"`
	Details   string     `json:"details"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"udpated_at"`
	Status    int        `json:"status" gorm:"default:0"`
	Tools     []Tool     `json:"tools" gorm:"foreignKey:ServiceOrderID"` // One-to-many
}
