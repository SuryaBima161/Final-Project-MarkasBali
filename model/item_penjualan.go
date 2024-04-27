package model

import (
	"time"

	"gorm.io/gorm"
)

type ItemPenjualan struct {
	ID           uint           `gorm:"not null" json:"id"`
	ID_Penjualan uint           `json:"id_penjualan"`
	Penjualan    Penjualan      `json:"penjualan" gorm:"foreignKey:ID_Penjualan"`
	ID_Barang    uint           `json:"id_barang"`
	Barang       Barang         `json:"barang" gorm:"foreignKey:ID_Barang"`
	Jumlah       uint           `gorm:"not null" json:"jumlah"`
	Subtotal     float64        `gorm:"not null" json:"subtotal"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
