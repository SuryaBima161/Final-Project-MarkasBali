package model

import (
	"time"

	"gorm.io/gorm"
)

type ItemPenjualan struct {
	ID           uint           `json:"id"`
	ID_Barang    uint           `json:"id_barang"`
	ID_Penjualan uint           `json:"id_penjualan"`
	Jumlah       uint           `json:"jumlah"`
	Subtotal     uint           `json:"subtotal"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
