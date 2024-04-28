package model

import (
	"time"

	"gorm.io/gorm"
)

type Penjualan struct {
	ID            uint            `gorm:"not null" json:"id"`
	Kode_Invoice  string          `gorm:"not null" json:"kode_invoice"`
	Nama_Pembeli  string          `gorm:"not null" json:"nama_pembeli"`
	Subtotal      float64         `gorm:"not null" json:"subtotal"`
	Kode_Diskon   string          `gorm:"index" json:"kode_diskon"`
	Diskon        float64         `gorm:"not null" json:"diskon"`
	Total         float64         `gorm:"not null" json:"total"`
	Kembalian     float64         `json:"kembalian"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
	DeletedAt     gorm.DeletedAt  `gorm:"index" json:"deleted_at"`
	CreatedBy     string          `json:"created_by"`
	ItemPenjualan []ItemPenjualan `json:"item_penjualan" gorm:"foreignKey:ID_Penjualan"`
}

func (cr *Penjualan) CreatePenjualan(db *gorm.DB) error {
	err := db.
		Model(Penjualan{}).
		Create(&cr).
		Error

	if err != nil {
		return err
	}

	return nil
}
func (cr *Penjualan) UpdateInvoicePenjualan(db *gorm.DB, id uint) error {
	err := db.
		Model(Penjualan{}).
		Where("id = ?", id).
		Update("kode_invoice", cr.Kode_Invoice).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (cr *ItemPenjualan) CreateItemPenjualan(db *gorm.DB) error {
	err := db.Model(&ItemPenjualan{}).Create(&cr).Error
	if err != nil {
		return err
	}
	return nil
}

func (cr *Penjualan) DeletePenjualan(db *gorm.DB) error {
	return db.Delete(cr).Error
}

func GetAllPenjualan(db *gorm.DB) ([]Penjualan, error) {
	res := []Penjualan{}

	err := db.
		Model(Penjualan{}).
		Order("created_at desc").
		Limit(50).
		Find(&res).
		Error

	if err != nil {
		return []Penjualan{}, err
	}

	return res, nil
}

func GetDetailPenjualan(db *gorm.DB, id uint) ([]Penjualan, error) {
	res := []Penjualan{}

	err := db.
		Model(Penjualan{}).Preload("ItemPenjualan").Where("id = ?", id).
		Find(&res).
		Error

	if err != nil {
		return []Penjualan{}, err
	}

	return res, nil
}

func (p *Penjualan) UpdateDiskonPenjualan(db *gorm.DB, id uint, diskon float64) error {
	if err := db.Model(&Penjualan{}).Where("id = ?", id).Update("diskon", diskon).Error; err != nil {
		return err
	}
	return nil
}

func (p *Penjualan) UpdateKembalianPenjualan(db *gorm.DB, id uint, kembalian float64) error {
	if err := db.Model(&Penjualan{}).Where("id = ?", id).Update("kembalian", kembalian).Error; err != nil {

		return err
	}
	return nil
}

func GetDetailtemPenjualan(db *gorm.DB, id uint) ([]ItemPenjualan, error) {
	res := []ItemPenjualan{}

	err := db.
		Model(ItemPenjualan{}).Where("id_penjualan = ?", id).Preload("Barang").
		Find(&res).
		Error

	if err != nil {
		return []ItemPenjualan{}, err
	}

	return res, nil
}
