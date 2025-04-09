package models

import "time"

type Wallet struct {
	ID        uint    `gorm:"primaryKey"`
	UserID    uint    `gorm:"not null"`
	Balance   float64 `gorm:"not null;default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
