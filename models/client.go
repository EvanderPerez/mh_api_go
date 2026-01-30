package models

import "time"

type Client struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	Name          string         `json:"name" gorm:"not null"`
	PhoneNumber   string         `json:"phone_number" gorm:"not null"`
	Metadata      JSONB          `json:"metadata" gorm:"type:jsonb"`
	ServiceOrders []ServiceOrder `json:"service_orders" gorm:"foreignKey:ClientID"`
	CreatedAt     *time.Time     `json:"-"`
	UpdatedAt     *time.Time     `json:"-"`
}
