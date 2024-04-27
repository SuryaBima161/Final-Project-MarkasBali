package controllers

import (
	"finalproject/model/payload"
	"finalproject/utils"
	"fmt"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func RoutePenjualan(app *fiber.App) {
	penjualanGroup := app.Group("/penjualan")
	penjualanGroup.Get("/", GetAllPenjualan)
	penjualanGroup.Get("/:id", GetDetailPenjualan)
	penjualanGroup.Post("/", InsertPenjualanController)
}

func InsertPenjualanController(c *fiber.Ctx) error {
	req := new(payload.AddPenjualanRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(map[string]interface{}{
				"message": "Body not valid",
			})
	}

	isValid, err := govalidator.ValidateStruct(*req)
	if !isValid {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"message": err.Error(),
		})
	}

	if len(req.Item_Penjualan) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"message": "Item penjualan tidak boleh kosong",
		})
	}

	// Panggil fungsi InsertPenjualan dari utils dengan data yang telah disiapkan
	data, err := utils.InsertPenjualan(req.Item_Penjualan, *req)
	if err != nil {
		// Tangani kesalahan sesuai kebutuhan
		fmt.Println("Error InsertPenjualan:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]any{
			"message": "Internal server error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(map[string]any{
		"data":    data,
		"message": "Success Insert Data",
	})
}

func GetAllPenjualan(c *fiber.Ctx) error {
	penjualanData, err := utils.GetAllPenjualan()
	if err != nil {
		logrus.Error("Error on get penjualan list: ", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(
			map[string]any{
				"message": "Server Error",
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		map[string]any{
			"data":    penjualanData,
			"message": "Success Get all",
		},
	)
}

func GetDetailPenjualan(c *fiber.Ctx) error {
	penjualanId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			map[string]any{
				"message": "ID not valid",
			},
		)
	}

	penjualanData, err := utils.GetDetailPenjualan(uint(penjualanId))
	if err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(
				map[string]any{
					"message": "ID not found",
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
			"data":    penjualanData,
			"message": "Sucess get data ",
		},
	)
}
