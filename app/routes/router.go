package routes

import (
	"api/app/controller"
	c "api/app/controller/api"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
)

// Handle all request to route to controller
func Handle(app *fiber.App) {
	app.Use(cors.New())
	services.InitDatabase()
	// services.InitRedis()

	api := app.Group(viper.GetString("ENDPOINT"))

	api.Get("/", controller.GetAPIIndex)
	api.Get("/info.json", controller.GetAPIInfo)
	api.Get("/user", c.GetAllUsers)
	api.Get("/user/:id", c.GetUserByID)
	api.Post("/user", c.CreateUser)
	api.Put("/user/:id", c.UpdateUser)
	api.Delete("/user/:id", c.DeleteUser)

}
