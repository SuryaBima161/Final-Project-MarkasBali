package model_test

import (
	"finalproject/config"
	"finalproject/model"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateDiskon(t *testing.T) {
	Init()

	dataDiskon := model.KodeDiskon{
		ID:          1,
		Kode_Diskon: "agg",
		Amount:      10000,
		Type:        "FIXED",
	}
	err := dataDiskon.CreateDiskon(config.Mysql.DB)
	assert.Nil(t, err)
	fmt.Println(dataDiskon.ID)
}

func TestGetAllDiskon(t *testing.T) {
	Init()
	dataDiskon := model.KodeDiskon{}
	resp, err := dataDiskon.GetAllDiskon(config.Mysql.DB)
	assert.Nil(t, err)
	fmt.Println(resp)
}

func TestGetSingleDiskon(t *testing.T) {
	Init()
	dataDiskon := model.KodeDiskon{
		ID:          1,
		Kode_Diskon: "agg",
		Amount:      10000,
		Type:        "FIXED",
	}
	err := dataDiskon.CreateDiskon(config.Mysql.DB)
	assert.Nil(t, err)
	resp, err := dataDiskon.GetSingleDiskon(config.Mysql.DB, dataDiskon.ID)
	assert.Nil(t, err)
	fmt.Println(resp)
}

func TestGetDiskonByKodeDiskon(t *testing.T) {
	Init()
	dataDiskon := model.KodeDiskon{
		ID:          8,
		Kode_Diskon: "agg",
		Amount:      10.000,
		Type:        "FIXED",
	}
	err := dataDiskon.CreateDiskon(config.Mysql.DB)
	assert.Nil(t, err)
	resp, err := dataDiskon.GetDiskonByKodeDiskon(config.Mysql.DB, "agg")
	assert.Nil(t, err)
	fmt.Println(resp)

}
