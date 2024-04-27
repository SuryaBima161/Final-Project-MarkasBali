package controllers

import (
	"finalproject/model"
	"finalproject/model/payload"
	"finalproject/utils"
	"fmt"
	"strconv"

	// "github.com/asaskevich/govalidator"
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func RouteBarang(app *fiber.App) {
	barangGroup := app.Group("/barang", CheckClient)
	barangGroup.Get("/", GetListBarang)
	barangGroup.Get("/:id", GetListDetail)
	barangGroup.Post("/", InsertBarang)
	barangGroup.Put("/:id", UpdateBarang)
	barangGroup.Put("/stok/:id", UpdateStokBarang)
	barangGroup.Delete("/:id", DeleteBarangById)
}
func CheckClient(c *fiber.Ctx) error {
	client := string(c.Request().Header.Peek("Client"))
	if client == "Mobile" {
		return c.Next()
	}
	return c.Status(fiber.StatusUnauthorized).JSON(map[string]any{
		"message": "user not authorized",
	})
}

func CheckRole(c *fiber.Ctx) error {
	client := string(c.Request().Header.Peek("Role"))
	if client == "Admin" {
		return c.Next()
	}
	return c.Status(fiber.StatusUnauthorized).JSON(map[string]any{
		"message": "role is not authorized",
	})
}
func InsertBarang(c *fiber.Ctx) error {

	req := new(payload.AddBarangRequest)
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

	car, errCreateBarang := utils.InsertBarangData(model.Barang{
		Nama:        fmt.Sprintf("%v", req.Nama),
		Harga_Pokok: req.Harga_Pokok,
		Harga_Jual:  req.Harga_Jual,
		Tipe_Barang: req.Tipe_Barang,
		Stok:        req.Stok,
		HistoriStok: req.History_Stok,
		CreatedBy:   req.CreatedBy,
	})

	if errCreateBarang != nil {
		logrus.Printf("Terjadi Error : %s\n", errCreateBarang.Error())
		return c.Status(fiber.StatusInternalServerError).
			JSON(map[string]any{
				"message": "Server error",
			})
	}
	return c.Status(fiber.StatusOK).JSON(map[string]any{
		"data":    car,
		"message": "Success Insert Data",
	})
}

func GetListBarang(c *fiber.Ctx) error {
	barangData, err := utils.GetListBarang()
	if err != nil {
		logrus.Error("Error on get cars list: ", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(
			map[string]any{
				"message": "Server Error",
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		map[string]any{
			"data":    barangData,
			"message": "Success Get all",
		},
	)
}

func GetListDetail(c *fiber.Ctx) error {
	barangId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			map[string]any{
				"message": "ID not valid",
			},
		)
	}

	barangData, err := utils.GetListDetail(uint(barangId))
	if err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(
				map[string]any{
					"message": "ID not found",
				},
			)
		}
		logrus.Error("Error on get barang data: ", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(
			map[string]any{
				"message": "Server Error",
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		map[string]any{
			"data":    barangData,
			"message": "Sucess get data ",
		},
	)
}
func UpdateBarang(c *fiber.Ctx) error {
	// Ambil id barang dari parameter rute
	barangId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			map[string]interface{}{
				"message": "ID not valid",
			},
		)
	}

	// Parse body permintaan ke dalam struktur payload.UpdateBarang
	var req *payload.UpdateBarang
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			map[string]interface{}{
				"message": "Body Not Valid",
			},
		)
	}

	// Validasi struktur payload.UpdateBarang
	isValid, err := govalidator.ValidateStruct(req)
	if !isValid {
		return c.Status(fiber.StatusBadRequest).JSON(
			map[string]interface{}{
				"message": err.Error(),
			},
		)
	}

	// Update data barang di database
	err = utils.UpdateBarang(uint(barangId), model.Barang{
		Nama:        req.Nama,
		Harga_Pokok: req.Harga_Pokok,
		Harga_Jual:  req.Harga_Jual,
		CreatedBy:   req.CreatedBy,
	})
	if err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(
				map[string]interface{}{
					"message": "ID not found",
				},
			)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(
			map[string]interface{}{
				"message": "Server Error",
			},
		)
	}

	// Kirim respons sukses jika pembaruan berhasil
	return c.Status(fiber.StatusOK).JSON(
		map[string]interface{}{
			"message": "Success Update Barang",
		},
	)
}

func UpdateStokBarang(c *fiber.Ctx) error {
	barangId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			map[string]interface{}{
				"message": "ID not valid",
			},
		)
	}

	var req *payload.UpdateStokBarangRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"message": "Body Not Valid",
		})
	}

	if _, err := govalidator.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"message": err.Error(),
		})
	}

	for _, h := range req.Histori_Stok {
		// Validasi histori stok
		if _, err := govalidator.ValidateStruct(h); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
				"message": "Histori Stok Not Valid: " + err.Error(),
			})
		}
	}

	for _, h := range req.Histori_Stok {
		// Update stok barang dan histori stok
		if err := utils.UpdateStokBarang(uint(barangId), model.Barang{
			Stok:        req.Stok,
			HistoriStok: req.Histori_Stok,
		}, h); err != nil {
			if err.Error() == "record not found" {
				return c.Status(fiber.StatusNotFound).JSON(map[string]interface{}{
					"message": "ID not found",
				})
			}
			logrus.Error("Error on updating barang data: ", err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
				"message": "Server Error",
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"message": "Success Update Barang",
	})
}

func DeleteBarangById(c *fiber.Ctx) error {
	barangId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			map[string]any{
				"message": "ID not valid",
			},
		)
	}

	err = utils.DeleteBarangById(uint(barangId))
	if err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(
				map[string]any{
					"message": "ID not found",
				},
			)
		}
		logrus.Error("Error on get barang data: ", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(
			map[string]any{
				"message": "Server Error",
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		map[string]any{
			"message": "Sucess Delete Barang ",
		},
	)
}
