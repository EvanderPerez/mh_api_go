package models

import "time"

type Location struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Name      string `json:"name" gorm:"index:idx_locations_name"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
