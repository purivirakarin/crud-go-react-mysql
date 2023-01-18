package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/purivirakarin/blogbackend/database"
	"github.com/purivirakarin/blogbackend/models"
)

func CreatePost(c *fiber.Ctx) error {
	var blogpost models.Blog 
	if err := c.BodyParser(&blogpost); err != nil {
		fmt.Println("Unable to parse body")
	}
	if err := database.DB.Create(&blogpost).Error; err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid payload",
		})
	}
	return c.JSON(fiber.Map {
		"message": "Congratulation!, You post is live!",
	})
}