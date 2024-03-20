package routes

import (
	"AstraTech/internal/handlers"

	"github.com/gofiber/fiber"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/upload-data", handlers.UploadData)
}
