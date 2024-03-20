package services

import (
	"AstraTech/internal/dtos"
	"encoding/json"
	"os"
	"time"

	"github.com/gofiber/fiber"
)

const directory = "/tmp/astra/files/"

func init() {
	if err := os.MkdirAll(directory, 0755); err != nil {
		panic("Error creating directory: " + err.Error())
	}
}
func SaveData(c *fiber.Ctx, req dtos.Req) (string, error) {
	filename := directory + time.Now().Format("20060102150405") + ".json"
	file, err := os.Create(filename)
	if err != nil {
		return "", nil
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(req)

	return filename, nil
}
