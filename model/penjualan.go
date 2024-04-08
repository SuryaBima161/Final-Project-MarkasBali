package model

import (
	"time"

	"gorm.io/gorm"
)

type Penjualan struct {
	ID           uint           `gorm:"not null" json:"id"`
	Kode_Invoice string         `gorm:"not null" json:"kode_invoice"`
	Nama_Pembeli string         `gorm:"not null" json:"nama_pembeli"`
	Subtotal     float64        `gorm:"not null" json:"subtotal"`
	Kode_Diskon  string         `gorm:"index" json:"kode_diskon"`
	Diskon       float64        `gorm:"not null" json:"diskon"`
	Total        float64        `gorm:"not null" json:"total"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	CreatedBy    string         `json:"created_by"`
}
