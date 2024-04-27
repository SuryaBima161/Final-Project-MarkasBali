package model

import (
	"time"

	"gorm.io/gorm"
)

type Barang struct {
	ID             uint            `json:"id"`
	Kode_Barang    string          `json:"kode_barang"`
	Nama           string          `json:"nama"`
	Harga_Pokok    float64         `json:"harga_pokok"`
	Harga_Jual     float64         `json:"harga_jual"`
	Tipe_Barang    string          `json:"tipe_barang"`
	Stok           uint            `json:"stok"`
	HistoriStok    []HistoriStok   `json:"histori_stok,omitempty" gorm:"foreignKey:ID_Barang"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
	DeletedAt      gorm.DeletedAt  `gorm:"index" json:"deleted_at"`
	CreatedBy      string          `json:"created_by"`
	Item_Penjualan []ItemPenjualan `json:"item_penjualan" gorm:"foreignKey:ID_Barang"`
}

func (cr *Barang) Create(db *gorm.DB) error {
	err := db.
		Model(Barang{}).
		Create(&cr).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (cr *Barang) GetListBarang(db *gorm.DB) ([]Barang, error) {
	res := []Barang{}

	err := db.
		Model(Barang{}).
		Find(&res).
		Error

	if err != nil {
		return []Barang{}, err
	}

	return res, nil
}

func (cr *Barang) GetBarangSpecific(db *gorm.DB) (Barang, error) {
	res := Barang{}

	err := db.
		Model(Barang{}).
		First(&res).
		Error

	if err != nil {
		return Barang{}, err
	}

	return res, nil
}

func (cr *Barang) GetBarangByKodeBarang(db *gorm.DB, kodebarang string) (Barang, error) {
	res := Barang{}

	err := db.
		Model(Barang{}).
		Where("kode_barang = ?", kodebarang).
		Take(&res).
		Error

	if err != nil {
		return Barang{}, err
	}

	return res, nil
}

func GetListDetail(db *gorm.DB, id uint) (Barang, error) {
	res := Barang{}

	err := db.
		Model(Barang{}).Preload("HistoriStok").Where("id = ?", id).
		Find(&res).
		Error

	if err != nil {
		return Barang{}, err
	}

	return res, nil
}

func (cr *Barang) UpdateBarang(db *gorm.DB, barangId uint) error {
	err := db.
		Model(&Barang{}).
		Select("nama", "harga_pokok", "harga_jual", "created_by").
		Where("id = ?", barangId).
		Updates(map[string]interface{}{
			"nama":        cr.Nama,
			"harga_pokok": cr.Harga_Pokok,
			"harga_jual":  cr.Harga_Jual,
			"created_by":  cr.CreatedBy,
		}).
		Error

	if err != nil {
		return err
	}

	return nil
}
func (cr *Barang) UpdateKodeBarang(db *gorm.DB) error {
	err := db.
		Model(Barang{}).
		Select("kode_barang").
		Where("id = ?", cr.ID).
		Updates(map[string]any{
			"kode_barang": cr.Kode_Barang,
		}).
		Error

	if err != nil {
		return err
	}

	return nil
}
func (cr *Barang) UpdateStokBarang(db *gorm.DB, barangId uint) error {
	err := db.
		Model(&Barang{}).
		Where("id = ?", barangId).
		Updates(map[string]interface{}{
			"stok": cr.Stok,
		}).
		Error
	if err != nil {
		return err
	}
	return nil
}

func (cr *HistoriStok) CreateHistoriStok(db *gorm.DB) error {
	err := db.Model(HistoriStok{}).Create(&cr).Error
	if err != nil {
		return err
	}
	return nil
}
func (cr *Barang) DeleteBarangById(db *gorm.DB, id uint) error {
	err := db.
		Model(Barang{}).
		Where("id = ?", id).
		Delete(&Barang{ID: id}).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (cr *Barang) DeleteHistoriStokByBarangID(db *gorm.DB, id uint) error {
	err := db.Model(HistoriStok{}).Where("id_barang = ?", id).Delete(&HistoriStok{}).Error
	if err != nil {
		return err
	}

	return nil
}
