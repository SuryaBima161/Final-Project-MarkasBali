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

func (cr *KodeDiskon) CreateDiskon(db *gorm.DB) error {
	err := db.
		Model(KodeDiskon{}).
		Create(&cr).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (cr *KodeDiskon) GetAllDiskon(db *gorm.DB) ([]KodeDiskon, error) {
	res := []KodeDiskon{}

	err := db.
		Model(KodeDiskon{}).
		Order("created_at desc"). // Mengurutkan berdasarkan tanggal pembuatan secara descending
		Limit(50).                // Batasan 50 data terakhir
		Find(&res).
		Error

	if err != nil {
		return []KodeDiskon{}, err
	}

	return res, nil
}

func (cr *KodeDiskon) GetSingleDiskon(db *gorm.DB, id uint) ([]KodeDiskon, error) {
	res := []KodeDiskon{}

	err := db.
		Model(KodeDiskon{}).Where("id = ?", id).
		Find(&res).
		Error

	if err != nil {
		return []KodeDiskon{}, err
	}

	return res, nil
}

func (cr *KodeDiskon) GetDiskonByKodeDiskon(db *gorm.DB, kodediskon string) (KodeDiskon, error) {
	res := KodeDiskon{}

	err := db.
		Model(KodeDiskon{}).Where("kode_diskon = ?", kodediskon).
		Find(&res).
		Error

	if err != nil {
		return KodeDiskon{}, err
	}

	return res, nil
}

func GetBarangByKodeBarangQuery(db *gorm.DB, kodeBarang string) ([]Barang, error) {
	var barang []Barang
	if err := db.Where("kode_barang = ?", kodeBarang).Find(&barang).Error; err != nil {
		return nil, err
	}
	return barang, nil
}

func (cr *KodeDiskon) DeleteKodeDiskon(db *gorm.DB, id uint) error {
	if err := db.Where("id = ?", id).Delete(&KodeDiskon{}).Error; err != nil {
		return err
	}
	return nil
}
