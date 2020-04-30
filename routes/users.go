package routes

import "github.com/gofiber/fiber"

func SetupUsersRouter(router *fiber.Group) {
	router.Get("/", func(c *fiber.Ctx) {
		c.Send("respond with a resource")
	})
}
