package utils

import (
	"finalproject/config"
	"finalproject/model"
	"time"
)

func InsertDiskon(data model.KodeDiskon) (model.KodeDiskon, error) {
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()

	err := data.CreateDiskon(config.Mysql.DB)
	if err != nil {
		return data, err
	}
	return data, nil
}

func GetAllDiskon() ([]model.KodeDiskon, error) {
	var kodediskon model.KodeDiskon
	return kodediskon.GetAllDiskon(config.Mysql.DB)
}
func GetSingleDiskon(id uint) ([]model.KodeDiskon, error) {
	kodediskon := model.KodeDiskon{
		ID: id,
	}
	return kodediskon.GetSingleDiskon(config.Mysql.DB, id)
}
func GetDiskonByKodeDiskon(kodeDiskon string) (model.KodeDiskon, error) {
	data := model.KodeDiskon{
		Kode_Diskon: kodeDiskon,
	}
	diskon, err := data.GetDiskonByKodeDiskon(config.Mysql.DB, kodeDiskon)
	if err != nil {
		return model.KodeDiskon{}, err
	}

	return diskon, nil
}
func DeleteKodeDiskon(id uint) (model.KodeDiskon, error) {
	kodedsikon := &model.KodeDiskon{}
	err := kodedsikon.DeleteKodeDiskon(config.Mysql.DB, id)
	if err != nil {
		return model.KodeDiskon{}, err
	}
	return *kodedsikon, err
}
