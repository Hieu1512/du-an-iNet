package router

import (
	"main/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Đăng ký các endpoint và các hàm xử lý tương ứng
	app.Get("/page", handler.GetPage)
	// app.Get("/page/get", handler.GetPage)
	app.Post("/page/create", handler.CreatePage)
	app.Patch("/page/update", handler.UpdatePage)
}
