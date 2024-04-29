package model_test

import (
	"finalproject/config"
	"finalproject/model"
	"fmt"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func Init() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("env not found, using global env")
	}
	config.OpenDB()
}

func TestCreateBarang(t *testing.T) {
	Init()

	dataBarang := model.Barang{
		ID:          3,
		Kode_Barang: "",
		Nama:        "",
		Harga_Pokok: 10.000,
		Harga_Jual:  12.000,
		Tipe_Barang: "",
		Stok:        33,
	}

	err := dataBarang.Create(config.Mysql.DB)
	assert.Nil(t, err)

	fmt.Println(dataBarang.ID)
}

func TestGetListBarang(t *testing.T) {
	Init()

	dataBarang := model.Barang{
		ID:          2,
		Kode_Barang: "MK01",
		Nama:        "Nasi Goreng",
		Harga_Pokok: 8000,
		Harga_Jual:  10000,
		Tipe_Barang: "Makanan",
		Stok:        121,
		CreatedBy:   "Admin",
	}

	err := dataBarang.Create(config.Mysql.DB)
	assert.Nil(t, err)

	res, err := dataBarang.GetListBarang(config.Mysql.DB)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(res), 1)

	fmt.Println(res)
}

func TestGetBarangSpesific(t *testing.T) {
	Init()

	dataBarang := model.Barang{
		ID: 1,
	}

	data, err := dataBarang.GetBarangSpecific(config.Mysql.DB)
	assert.Nil(t, err)

	fmt.Println(data)
}

func TestGetBarangByKodeBarang(t *testing.T) {
	Init()

	carData := model.Barang{}

	data, err := carData.GetBarangByKodeBarang(config.Mysql.DB,
		//Masukan Kode barang di bawah ini "example"
		"MK01")
	assert.Nil(t, err)

	fmt.Println(data)
}

func TestGetListDetail(t *testing.T) {
	Init()

	barang := model.Barang{
		ID: 3,
	}
	data, err := model.GetListDetail(config.Mysql.DB, barang.ID)
	assert.Nil(t, err)

	fmt.Println(data)
}

// Entahlah ga paham masuk ke history_stok

func TestUpdateBarang(t *testing.T) {
	Init()

	barang := model.Barang{
		ID:          3,
		Nama:        "Nasi Kotak",
		Harga_Pokok: 8500,
		Harga_Jual:  10000,
		CreatedBy:   "Admin",
	}

	err := barang.UpdateBarang(config.Mysql.DB, barang.ID)
	assert.Nil(t, err)
}

func TestUpdateKodeBarang(t *testing.T) {
	Init()

	barang := model.Barang{
		ID:          1,
		Kode_Barang: "MI02",
	}

	err := barang.UpdateKodeBarang(config.Mysql.DB)
	assert.Nil(t, err)
}

func TestUpdateStokBarang(t *testing.T) {
	Init()

	barang := model.Barang{
		ID:   3,
		Stok: 222,
	}

	err := barang.UpdateStokBarang(config.Mysql.DB, barang.ID)
	assert.Nil(t, err)
}

func TestCreateHistoriStok(t *testing.T) {
	Init()

	historiStok := model.HistoriStok{
		ID:         1,
		ID_Barang:  2,
		Amount:     121,
		Status:     "IN",
		Keterangan: "Initial Stok",
	}

	err := historiStok.CreateHistoriStok(config.Mysql.DB)
	assert.Nil(t, err)

	fmt.Println(historiStok.ID)
}

func TestDeleteBarangByID(t *testing.T) {
	Init()

	dataBarang := model.Barang{
		ID: 3,
	}

	err := dataBarang.DeleteBarangById(config.Mysql.DB, dataBarang.ID)
	assert.Nil(t, err)
}

func TestDeleteHistoriStokByBarangID(t *testing.T) {
	Init()
	barang := new(model.Barang)
	historiStok := model.HistoriStok{
		ID: 1,
	}

	err := barang.DeleteHistoriStokByBarangID(config.Mysql.DB, historiStok.ID)
	assert.Nil(t, err)
}
