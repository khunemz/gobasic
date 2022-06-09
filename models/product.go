package models

import "time"

type Product struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Name         string    `json:"product_name"`
	SerialNumber string    `json:"serial_number"`
	CreatedAt    time.Time `json:"created_at"`
}
