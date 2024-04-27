package payload

import (
	"finalproject/model"
	"time"

	"gorm.io/gorm"
)

type AddBarangRequest struct {
	Nama         string              `json:"nama_barang" valid:"required, type(string)"`
	Harga_Pokok  float64             `json:"harga_pokok" valid:"optional , type(float64)"`
	Harga_Jual   float64             `json:"harga_jual" valid:"optional , type(float64)"`
	Tipe_Barang  string              `json:"tipe_barang" valid:"required, type(string)"`
	Stok         uint                `json:"stok" valid:"required, type(uint)"`
	History_Stok []model.HistoriStok `json:"histori_stok" valid:"required"`
	CreatedBy    string              `json:"created_by" valid:"required, type(string)"`
}
type GetListBarangRespone struct {
	ID           uint           `json:"id"`
	Kode_Barang  string         `json:"kode_barang"`
	Nama         string         `json:"nama_barang"`
	Harga_Pokok  float64        `json:"harga_pokok"`
	Harga_Jual   float64        `json:"harga_jual"`
	Tipe_Barang  string         `json:"tipe_barang"`
	Stok         uint           `json:"stok"`
	Created_At   time.Time      `json:"created_at"`
	Updated_At   time.Time      `json:"updated_at"`
	Deleted_at   gorm.DeletedAt `json:"deleted_at"`
	CreatedBy    string         `json:"created_by"`
	Histori_Stok []Histori_Stok `json:"histori_stok"`
}
type Histori_Stok struct {
	Amount     float64        `json:"amount"`
	Status     string         `json:"status"`
	Keterangan string         `json:"keterangan"`
	Created_At time.Time      `json:"created_at"`
	Updated_At time.Time      `json:"updated_at"`
	Deleted_at gorm.DeletedAt `json:"deleted_at"`
}

type UpdateBarang struct {
	Nama        string  `json:"nama_barang"`
	Harga_Pokok float64 `json:"harga_pokok"`
	Harga_Jual  float64 `json:"harga_jual"`
	CreatedBy   string  `json:"created_by"`
}
type UpdateStokBarangRequest struct {
	Stok         uint                `json:"stok"`
	Histori_Stok []model.HistoriStok `json:"histori_stok"`
}
