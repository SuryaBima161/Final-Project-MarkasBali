package model

import (
	"time"

	"gorm.io/gorm"
)

type KodeDiskon struct {
	ID          uint           `json:"id"`
	Kode_Diskon string         `json:"kode_diskon"`
	Amount      float64        `json:"amount"`
	Type        string         `json:"type"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
