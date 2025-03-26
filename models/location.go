package models

import "time"

type Location struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"index:idx_locations_name"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
