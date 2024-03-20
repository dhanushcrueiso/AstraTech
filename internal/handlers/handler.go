package handlers

import (
	"AstraTech/internal/database/daos"
	"AstraTech/internal/dtos"
	"AstraTech/internal/services"
	"fmt"

	"github.com/gofiber/fiber"
)

func UploadData(c *fiber.Ctx) {
	fmt.Println("here:")
	req := dtos.Req{}
	err := c.BodyParser(&req)
	if err != nil {
		fmt.Println("", err)
		return
	}
	fmt.Println("checking the parser", req)

	name, err := services.SaveData(c, req)
	if err != nil {
		return
	}
	fmt.Println("here:")
	daos.ProcessData(name)
	c.JSON("successfull")
}
