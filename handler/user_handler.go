package handler

import "github.com/gofiber/fiber/v2"

type UserHandlerInterface interface {
	FindAll(c *fiber.Ctx) error
	FindByID(c *fiber.Ctx) error
	Save(c *fiber.Ctx) error
	Edit(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}
