package handler

import (
	"main/database"
	"main/model"

	"github.com/gofiber/fiber/v2"
)

func GetPage(c *fiber.Ctx) error {
	var page []model.Page2
	if err := database.DB.Find(&page).Error; err != nil {
		return err
	}
	return c.JSON(page)
}

// func GetPage(c *fiber.Ctx) error {
// 	// Lấy thông tin page từ cơ sở dữ liệu
// 	var page model.Page2
// 	if err := database.DB.First(&page, c.Params("id")).Error; err != nil {
// 		return err
// 	}
// 	return c.JSON(page)
// }

func CreatePage(c *fiber.Ctx) error {
	// Parse request body để tạo mới một page
	var page model.Page2
	if err := c.BodyParser(&page); err != nil {
		return err
	}

	// Tạo mới page trong cơ sở dữ liệu
	if err := database.DB.Create(&page).Error; err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(page)
}

func UpdatePage(c *fiber.Ctx) error {
	// Parse request body để cập nhật thông tin của page
	var updatedPage model.Page2
	if err := c.BodyParser(&updatedPage); err != nil {
		return err
	}

	// Tìm page cần cập nhật trong cơ sở dữ liệu dựa trên trường Key
	var existingPage model.Page2
	if err := database.DB.Where("`key` = ?", updatedPage.Key).First(&existingPage).Error; err != nil {
		return err
	}

	// Cập nhật thông tin của page
	existingPage.Label = updatedPage.Label
	existingPage.Link = updatedPage.Link
	existingPage.Metadata = updatedPage.Metadata

	// Lưu thay đổi vào cơ sở dữ liệu
	if err := database.DB.Save(&existingPage).Error; err != nil {
		return err
	}

	return c.JSON(existingPage)
}
