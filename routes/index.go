package routes

import "github.com/gofiber/fiber"

func SetupIndexRouter(router *fiber.Group) {
	router.Get("/", func(c *fiber.Ctx) {
		bind := fiber.Map{
			"title": "Express",
		}
		if err := c.Render("./index", bind); err != nil {
			c.Status(500).Send(err.Error())
		}
	})
}
