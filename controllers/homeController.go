package controllers

import "github.com/gofiber/fiber/v2"

func Home(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
}
