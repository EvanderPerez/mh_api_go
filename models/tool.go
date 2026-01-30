package models

import "time"

type Tool struct {
	ID             uint         `json:"id" gorm:"primaryKey"`
	Model          string       `json:"model" gorm:"not null"`
	ToolType       string       `json:"tool_type" gorm:"default:null"`
	Brand          string       `json:"brand" gorm:"not null"`
	Accesories     string       `json:"accesories" gorm:"default:null"`
	Location       string       `json:"location" gorm:"not null"`
	CreatedAt      *time.Time   `json:"-"`
	UpdatedAt      *time.Time   `json:"-"`
	ServiceOrderID uint         `json:"-" gorm:"default:null"`
	ServiceOrder   ServiceOrder `json:"-" gorm:"foreignKey:ServiceOrderID"`
}
