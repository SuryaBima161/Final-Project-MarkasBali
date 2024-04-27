package utils

import (
	"finalproject/config"
	"finalproject/model"
	"finalproject/model/payload"
	"fmt"
	"sync"
	"time"
)

func GetListBarang() ([]model.Barang, error) {
	var barang model.Barang
	return barang.GetListBarang(config.Mysql.DB)
}

func GetBarangSpecifific() (model.Barang, error) {
	var barang model.Barang
	return barang.GetBarangSpecific(config.Mysql.DB)
}

func GetListDetail(id uint) (resp payload.GetListBarangRespone, err error) {
	barang, err := model.GetListDetail(config.Mysql.DB, id)
	if err != nil {
		return
	}
	historiStok := make([]payload.Histori_Stok, len(barang.HistoriStok))
	for i, h := range barang.HistoriStok {
		historiStok[i] = payload.Histori_Stok{
			Amount:     h.Amount,
			Status:     h.Status,
			Keterangan: h.Keterangan,
			Created_At: h.CreatedAt,
			Updated_At: h.UpdatedAt,
			Deleted_at: h.DeletedAt,
		}
	}
	resp = payload.GetListBarangRespone{
		ID:           barang.ID,
		Kode_Barang:  barang.Kode_Barang,
		Nama:         barang.Nama,
		Harga_Pokok:  barang.Harga_Pokok,
		Harga_Jual:   barang.Harga_Jual,
		Tipe_Barang:  barang.Tipe_Barang,
		Stok:         barang.Stok,
		Created_At:   barang.CreatedAt,
		Updated_At:   barang.UpdatedAt,
		Deleted_at:   barang.DeletedAt,
		CreatedBy:    barang.CreatedBy,
		Histori_Stok: historiStok,
	}
	return
}

func InsertBarangData(data model.Barang) (model.Barang, error) {
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()

	err := data.Create(config.Mysql.DB)
	if err != nil {
		return data, err
	}

	data.Kode_Barang = GenerateKodeBarang(data.Tipe_Barang, data.Kode_Barang, data.ID)

	err = data.UpdateKodeBarang(config.Mysql.DB)
	if err != nil {
		return data, err
	}

	return data, nil
}

func GenerateKodeBarang(tipeBarang, kodeBarang string, id uint) string {
	switch tipeBarang {
	case "Makanan":
		return fmt.Sprintf("MA-%d", id)
	case "Minuman":
		return fmt.Sprintf("MI-%d", id)
	case "Lainnya":
		return fmt.Sprintf("L-%d", id)
	default:
		return kodeBarang
	}
}
func UpdateBarang(id uint, newData model.Barang) error {
	updatedBarang := model.Barang{
		Nama:        newData.Nama,
		Harga_Pokok: newData.Harga_Pokok,
		Harga_Jual:  newData.Harga_Jual,
		CreatedBy:   newData.CreatedBy,
	}
	err := updatedBarang.UpdateBarang(config.Mysql.DB, id)
	if err != nil {
		return err
	}

	return nil
}
func UpdateStokBarang(id uint, newData model.Barang, dataH model.HistoriStok) error {
	updatedBarang := model.Barang{
		ID:   id,
		Stok: newData.Stok,
	}

	if err := updatedBarang.UpdateStokBarang(config.Mysql.DB, id); err != nil {
		return err
	}

	createHistory := model.HistoriStok{
		ID_Barang:  id,
		Amount:     dataH.Amount,
		Status:     dataH.Status,
		Keterangan: dataH.Keterangan,
	}

	if err := createHistory.CreateHistoriStok(config.Mysql.DB); err != nil {
		return err
	}

	return nil
}

func UpdateStokBarangArray(id uint, barangUpdates []model.Barang, historiUpdates []model.HistoriStok) error {
	// Inisialisasi sync.WaitGroup
	var wg sync.WaitGroup

	errCh := make(chan error, len(barangUpdates)+len(historiUpdates))

	for _, update := range barangUpdates {
		wg.Add(1)
		go func(u model.Barang) {
			defer wg.Done()
			updatedBarang := model.Barang{
				ID:   id,
				Stok: u.Stok,
			}
			if err := updatedBarang.UpdateStokBarang(config.Mysql.DB, u.ID); err != nil {
				errCh <- err
			}
		}(update)
	}

	for _, update := range historiUpdates {
		wg.Add(1)
		go func(u model.HistoriStok) {
			defer wg.Done()
			createHistory := model.HistoriStok{
				ID_Barang:  id,
				Amount:     u.Amount,
				Status:     u.Status,
				Keterangan: u.Keterangan,
			}
			if err := createHistory.CreateHistoriStok(config.Mysql.DB); err != nil {
				errCh <- err
			}
		}(update)
	}

	// Menunggu semua goroutine selesai
	go func() {
		wg.Wait()
		close(errCh)
	}()

	for err := range errCh {
		if err != nil {
			return err
		}
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
