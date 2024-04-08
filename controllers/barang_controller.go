package controllers

import (
	"finalproject/model"
	"finalproject/utils"
	"fmt"
	"strconv"

	// "github.com/asaskevich/govalidator"
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func RouteCars(app *fiber.App) {
	barangGroup := app.Group("/barang", CheckClient)
	barangGroup.Get("/", GetListBarang)
	barangGroup.Get("/:id", GetListDetail)
	// barangGroup.Get("/by-id/:id", GetCarByID)
	barangGroup.Post("/", InsertCarData)
	barangGroup.Put("/:id", UpdateBarang)
	// barangGroup.Put("/", InsertCarData)
	barangGroup.Delete("/:id", DeleteBarangById)
	// barangGroup.Post("import-csv", CheckRole, ImportCsvFile)

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
func InsertCarData(c *fiber.Ctx) error {
	type AddCarRequest struct {
		Kode_Barang  string              `json:"kode_barang" valid:"required, type(string)"`
		Nama         string              `json:"nama" valid:"required, type(string)"`
		Harga_Pokok  float64             `json:"harga_pokok" valid:"optional , type(float64)"`
		Harga_Jual   float64             `json:"harga_jual" valid:"optional , type(float64)"`
		Tipe_Barang  string              `json:"tipe_barang" valid:"required, type(string)"`
		Stok         uint                `json:"stok" valid:"required, type(uint)"`
		History_Stok []model.HistoriStok `json:"histori_stok" valid:"required"`
		CreatedBy    string              `json:"created_by" valid:"required, type(string)"`
	}
	req := new(AddCarRequest)
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

	car, errCreateCar := utils.InsertCarData(model.Barang{
		Kode_Barang: req.Kode_Barang,
		Nama:        fmt.Sprintf("%v", req.Nama),
		Harga_Pokok: req.Harga_Pokok,
		Harga_Jual:  req.Harga_Jual,
		Tipe_Barang: req.Tipe_Barang,
		Stok:        req.Stok,
		HistoriStok: req.History_Stok,
		CreatedBy:   req.CreatedBy,
	})

	if errCreateCar != nil {
		logrus.Printf("Terjadi Error : %s\n", errCreateCar.Error())
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
	barangId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			map[string]any{
				"message": "ID not valid",
			},
		)
	}
	type requestBody struct {
		Nama        string  `json:"nama"`
		Harga_Pokok float64 `json:"harga_pokok"`
		Harga_Jual  float64 `json:"harga_jual"`
		CreatedBy   string  `json:"created_by"`
	}
	req := new(requestBody)
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
	err = utils.UpdateBarang(uint(barangId), model.Barang{
		Nama:        req.Nama,
		Harga_Pokok: req.Harga_Pokok,
		Harga_Jual:  req.Harga_Jual,
		CreatedBy:   req.CreatedBy,
	})
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
			"message": "Sucess Update Barang ",
		},
	)
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
