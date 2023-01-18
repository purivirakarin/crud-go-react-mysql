package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/purivirakarin/blogbackend/controller"
	//"github.com/purivirakarin/blogbackend/middleware"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)
	//app.Use(middleware.IsAuthenticate)
}