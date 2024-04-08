package utils

import (
	"finalproject/config"
	"finalproject/model"
	"time"
)

func GetListBarang() ([]model.Barang, error) {
	var barang model.Barang
	return barang.GetListBarang(config.Mysql.DB)
}

func GetListDetail(id uint) ([]model.Barang, error) {
	barang := model.Barang{
		ID: id,
	}
	return barang.GetListDetail(config.Mysql.DB, id)
}

func InsertCarData(data model.Barang) (model.Barang, error) {
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	err := data.Create(config.Mysql.DB)

	return data, err
}

func UpdateBarang(id uint, newData model.Barang) error {
	updatedBarang := model.Barang{
		ID:          id,
		Nama:        newData.Nama,
		Harga_Pokok: newData.Harga_Pokok,
		Harga_Jual:  newData.Harga_Jual,
		Tipe_Barang: newData.Tipe_Barang,
		Stok:        newData.Stok,
		CreatedBy:   newData.CreatedBy,
	}
	err := updatedBarang.UpdateBarang(config.Mysql.DB)
	if err != nil {
		return err
	}

	return nil
}

func DeleteBarangById(id uint) (err error) {
	model := &model.Barang{}
	if err := model.DeleteHistoriStokByBarangID(config.Mysql.DB, id); err != nil {
		return err
	}
	if err := model.DeleteBarangById(config.Mysql.DB, id); err != nil {
		return err
	}
	return nil
}
