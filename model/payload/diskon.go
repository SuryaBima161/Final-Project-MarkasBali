package payload

import "time"

type DiscountResponse struct {
	ID         uint       `json:"id"`
	KodeDiskon string     `json:"kode_diskon"`
	Amount     float64    `json:"amount"`
	Type       string     `json:"type"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
}

type DiscountRequest struct {
	KodeDiskon string  `json:"kode_diskon"`
	Amount     float64 `json:"amount"`
	Type       string  `json:"type"`
}
