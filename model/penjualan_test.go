package model_test

import (
	"finalproject/config"
	"finalproject/model"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertPenjualan(t *testing.T) {
	Init()

	dataBarang := model.Penjualan{
		ID:           3,
		Kode_Invoice: "INV-3",
		Nama_Pembeli: "Ahok",
		Subtotal:     10000,
		Kode_Diskon:  "dddkjf",
		Diskon:       2000,
		Total:        8000,
		Kembalian:    2000,
		CreatedBy:    "System",
		ItemPenjualan: []model.ItemPenjualan{
			{
				ID:           1,
				ID_Penjualan: 3,
				ID_Barang:    1,
				Jumlah:       2,
				Subtotal:     10000,
			},
		},
	}
	err := dataBarang.CreatePenjualan(config.Mysql.DB)
	assert.Nil(t, err)
	fmt.Println(dataBarang.ID)
}

func TestCreateItemPenjualan(t *testing.T) {
	Init()
	dataBarang := model.Penjualan{
		ID:           4,
		Kode_Invoice: "INV-4",
		Nama_Pembeli: "Ahok",
		Subtotal:     10000,
		Kode_Diskon:  "dddkjf",
		Diskon:       2000,
		Total:        8000,
		Kembalian:    2000,
		CreatedBy:    "System",
	}
	err := dataBarang.CreatePenjualan(config.Mysql.DB)
	assert.Nil(t, err)

	dataItem := &model.ItemPenjualan{
		ID:           2,
		ID_Penjualan: 4,
		ID_Barang:    1,
		Jumlah:       2,
		Subtotal:     20000,
	}
	err = dataItem.CreateItemPenjualan(config.Mysql.DB)
	assert.Nil(t, err)
	fmt.Println(dataItem.ID)
}

func TestGetAllPenjualan(t *testing.T) {
	Init()
	resp, err := model.GetAllPenjualan(config.Mysql.DB)
	assert.Nil(t, err)
	fmt.Println(resp)
}

func TestGetDetailPenjualan(t *testing.T) {
	Init()
	idPenjualan := uint(78) // Ganti dengan ID yang valid
	resp, err := model.GetDetailPenjualan(config.Mysql.DB, idPenjualan)
	assert.Nil(t, err)
	fmt.Println(resp)
}
