package model

import (
	"time"

	"gorm.io/gorm"
)

type HistoriStok struct {
	ID         uint           `json:"id"`
	ID_Barang  uint       `json:"id_barang"`
	Amount     int            `json:"amount"`
	Status     string         `json:"status"`
	Keterangan string         `json:"keterangan"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
