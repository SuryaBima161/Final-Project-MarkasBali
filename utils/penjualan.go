package utils

import (
	"errors"
	"finalproject/config"
	"finalproject/model"
	"finalproject/model/payload"
	"fmt"
	"sync"
	"time"
)

func GenerateInvoice(id uint) string {
	invoice := fmt.Sprintf("INV - %d", id)
	return invoice
}
func GetBarangByKodeBarang(kodebarang string) (model.Barang, error) {
	barang := &model.Barang{}
	return barang.GetBarangByKodeBarang(config.Mysql.DB, kodebarang)
}

func InsertPenjualan(kodebarang []payload.ItemPenjualanRequest, req payload.AddPenjualanRequest) (model.Penjualan, error) {
	currentTime := time.Now()

	// Buat entitas penjualan baru
	penjualan := model.Penjualan{
		CreatedAt:    currentTime,
		UpdatedAt:    currentTime,
		Nama_Pembeli: req.Nama_Pembeli,
		Subtotal:     req.Subtotal,
		Kode_Diskon:  req.Kode_Diskon,
		Total:        req.Total,
		CreatedBy:    req.CreatedBy,
	}

	if penjualan.Kode_Diskon != "" {
		discount, err := GetDiskonByKodeDiskon(penjualan.Kode_Diskon)
		if err != nil {
			return penjualan, fmt.Errorf("error getting discount by code %s: %v", penjualan.Kode_Diskon, err)
		}
		if discount.ID != 0 {
			if discount.Type == "FIXED" {
				penjualan.Total -= discount.Amount
			} else if discount.Type == "PERCENT" {
				discountAmount := (discount.Amount / 100) * penjualan.Subtotal
				penjualan.Total -= discountAmount
			}
			if err := discount.DeleteKodeDiskon(config.Mysql.DB, discount.ID); err != nil {
				return penjualan, err
			}
		}
	}

	if penjualan.Total < penjualan.Subtotal {
		return penjualan, errors.New("transaksi gagal: total pembayaran kurang dari subtotal")
	}

	// Hitung nilai diskon
	diskon := penjualan.Subtotal - penjualan.Total
	// Pastikan nilai diskon tidak negatif
	if diskon < 0 {
		diskon = 0
	}
	penjualan.Diskon = diskon

	if err := penjualan.CreatePenjualan(config.Mysql.DB); err != nil {
		return penjualan, fmt.Errorf("error creating penjualan: %v", err)
	}
	if err := penjualan.UpdateDiskonPenjualan(config.Mysql.DB, penjualan.ID); err != nil {
		return penjualan, fmt.Errorf("error updating penjualan: %v", err)
	}

	penjualan.Kode_Invoice = GenerateInvoice(penjualan.ID)
	if err := penjualan.UpdateInvoicePenjualan(config.Mysql.DB, penjualan.ID); err != nil {
		return penjualan, fmt.Errorf("error updating kode invoice: %v", err)
	}

	var wg sync.WaitGroup
	wg.Add(len(kodebarang))

	for _, item := range kodebarang {
		barang, err := GetBarangByKodeBarang(item.Kode_Barang)
		if err != nil {
			return penjualan, fmt.Errorf("error getting barang by kode barang %s: %v", item.Kode_Barang, err)
		}
		if barang.Stok < item.Jumlah {
			return penjualan, fmt.Errorf("stok barang %s tidak mencukupi", item.Kode_Barang)
		}
		newItem := model.ItemPenjualan{
			ID_Penjualan: penjualan.ID,
			ID_Barang:    barang.ID,
			Jumlah:       item.Jumlah,
			Subtotal:     float64(item.Subtotal),
		}
		if err := newItem.CreateItemPenjualan(config.Mysql.DB); err != nil {
			return penjualan, fmt.Errorf("error creating ItemPenjualan: %v", err)
		}

		stokBaru := barang.Stok - item.Jumlah
		historiUpdate := model.HistoriStok{
			ID_Barang:  barang.ID,
			Amount:     float64(item.Jumlah),
			Status:     "OUT",
			Keterangan: "JUAL",
		}
		barangUpdate := model.Barang{
			ID:   barang.ID,
			Stok: stokBaru,
		}

		go func(barangUpdate model.Barang, historiUpdate model.HistoriStok) {
			defer wg.Done()
			if err := UpdateStokBarangArray(barang.ID, []model.Barang{barangUpdate}, []model.HistoriStok{historiUpdate}); err != nil {
				fmt.Println("Error updating stok barang:", err)
			}
		}(barangUpdate, historiUpdate)
	}

	wg.Wait()
	return penjualan, nil
}

func GetAllPenjualan() ([]model.Penjualan, error) {
	var penjualan model.Penjualan
	return penjualan.GetAllPenjualan(config.Mysql.DB)
}
func GetDetailPenjualan(id uint) ([]model.Penjualan, error) {
	penjualan := model.Penjualan{
		ID: id,
	}
	return penjualan.GetDetailPenjualan(config.Mysql.DB, id)
}
