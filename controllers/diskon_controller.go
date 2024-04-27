package controllers

import (
	"finalproject/config"
	"finalproject/model"
	"finalproject/model/payload"
	"finalproject/utils"
	"log"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func RoutDiskon(app *fiber.App) {
	penjualanGroup := app.Group("/kode-diskon")
	penjualanGroup.Get("/", GetAllDiskon)
	penjualanGroup.Get("/:id", GetSingleDiskon)
	penjualanGroup.Get("/get-by-code", GetDiskonByKodeDiskonController)
	penjualanGroup.Post("/", InsertDiskon)
	penjualanGroup.Get("/get-by-kode-barang", GetBarangByKodeBarangController)
}

func InsertDiskon(c *fiber.Ctx) error {

	req := new(payload.DiscountRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(map[string]any{
				"message": "Body Not Valid",
			})
	}

	isValid, err := govalidator.ValidateStruct(req)
	if !isValid {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]any{
			"message": err.Error(),
		})
	}

	diskon, errCreateDiskon := utils.InsertDiskon(model.KodeDiskon{
		Kode_Diskon: req.KodeDiskon,
		Amount:      req.Amount,
		Type:        req.Type,
	})

	if errCreateDiskon != nil {
		logrus.Printf("Terjadi Error : %s\n", errCreateDiskon.Error())
		return c.Status(fiber.StatusInternalServerError).
			JSON(map[string]any{
				"message": "Server error",
			})
	}
	return c.Status(fiber.StatusOK).JSON(map[string]any{
		"data":    diskon,
		"message": "Success Insert Data",
	})
}

func GetDiskonByKodeDiskonController(c *fiber.Ctx) error {
	// Menerima nilai kode diskon dari query parameter
	kodeDiskon := c.Query("kode_diskon")

	// Validasi kode diskon
	if kodeDiskon == "" {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]any{
			"message": "Kode diskon is required",
		})
	}

	// Memanggil fungsi GetDiskonByKodeDiskon untuk mendapatkan data kode diskon
	diskon, err := utils.GetDiskonByKodeDiskon(kodeDiskon)
	if err != nil {
		logrus.Error("Error on get diskon list: ", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(
			map[string]any{
				"message": "Server Error",
			},
		)
	}
	log.Printf("Kode Diskon: %s, Diskon: %v\n", kodeDiskon, diskon)
	return c.Status(fiber.StatusOK).JSON(
		map[string]any{
			"data":    diskon,
			"message": "Success Get all",
		},
	)
}

func GetAllDiskon(c *fiber.Ctx) error {
	diskonData, err := utils.GetAllDiskon()
	if err != nil {
		logrus.Error("Error on get diskon list: ", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(
			map[string]any{
				"message": "Server Error",
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		map[string]any{
			"data":    diskonData,
			"message": "Success Get all",
		},
	)
}

func GetSingleDiskon(c *fiber.Ctx) error {
	diskonId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			map[string]any{
				"message": "ID not valid",
			},
		)
	}

	diskonData, err := utils.GetSingleDiskon(uint(diskonId))
	if err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(
				map[string]any{
					"message": "record not found",
				},
			)
		}
		logrus.Error("Error on get penjualan data: ", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(
			map[string]any{
				"message": "Server Error",
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		map[string]any{
			"data":    diskonData,
			"message": "Sucess get data ",
		},
	)
}

// GetBarangByKodeBarangController mengembalikan semua barang berdasarkan kode barang
func GetBarangByKodeBarangController(c *fiber.Ctx) error {
	// Menerima nilai kode barang dari query parameter
	kodeBarang := c.Query("kode_barang")

	// Memanggil fungsi GetBarangByKodeBarangQuery untuk mendapatkan data barang berdasarkan kode barang
	barang, err := model.GetBarangByKodeBarangQuery(config.Mysql.DB, kodeBarang)

	if err != nil {
		// Tangani kesalahan jika data tidak ditemukan atau terjadi kesalahan lainnya
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{

			"message": "Failed to get barang by kode barang",
			"error":   err.Error(),
		})
	}

	// Mengembalikan data barang sebagai respons
	return c.JSON(fiber.Map{
		"data":    barang,
		"message": "Success Get barang by kode barang",
	})
}
