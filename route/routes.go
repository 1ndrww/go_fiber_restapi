package route

import (
	"example-go-restapi/exception"
	"example-go-restapi/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
)

// New create an instance of Book app route
func New(
	userHandler handler.UserHandlerInterface,
) *fiber.App {
	app := fiber.New(fiber.Config{ErrorHandler: exception.ErrorHandler})
	app.Use(recover2.New())
	app.Use(cors.New())

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/users", userHandler.FindAll)
	v1.Get("/users/:id", userHandler.FindByID)
	v1.Post("/users", userHandler.Save)
	v1.Put("/users/:id", userHandler.Edit)
	v1.Delete("/users/:id", userHandler.Delete)

	return app
}
