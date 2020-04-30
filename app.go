package main

import (
	"gofiber/starter/routes"

	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
	"github.com/gofiber/template"
)

func main() {
	app := fiber.New()

	// view engine setup
	app.Settings.TemplateFolder = "./views"
	app.Settings.TemplateEngine = template.Pug()
	app.Settings.TemplateExtension = ".jade"

	app.Use(logger.New())
	app.Static("/public", "./public")

	routes.SetupIndexRouter(app.Group("/"))
	routes.SetupUsersRouter(app.Group("/users"))

	/*
		This is for router params example
	*/
	/*
	  Route path: /users/:userId/books/:bookId
	  Request URL: http://localhost:3000/users/34/books/8989
	  req.params: { "userId": "34", "bookId": "8989" }
	*/
	app.Get("/users/:userId/books/:bookId", func(c *fiber.Ctx) {
		c.JSON(fiber.Map{
			"userId": c.Params("userId"),
			"bookId": c.Params("bookId"),
		})
	})

	/*
		This is for chained middlewares example
	*/
	app.Get("/verify/:status/:role/:userId", authenticate, authorize, func(c *fiber.Ctx) {
		if c.Locals("isAuthenticated") == false {
			c.Status(403)
			c.Send("Unauthenticated. Please signup!")
			return
		}
		c.Send("Redirecting " + c.Locals("redirectRoute").(string))
	})

	// catch 404 error handler
	app.Use("/*", func(c *fiber.Ctx) {
		bind := fiber.Map{
			"title":   "Page not found",
			"message": "Page not found",
			"error": map[string]interface{}{
				"status": 400,
				"stack":  "",
			},
		}
		// render the error page
		if err := c.Render("./error", bind); err != nil {
			c.Status(500).Send(err.Error())
		}
	})

	app.Listen(3000)
}

/*
	This is for chained middlewares example
*/
func authenticate(c *fiber.Ctx) {
	if c.Params("status") == "authenticated" {
		c.Locals("isAuthenticated", true)
	} else {
		c.Locals("isAuthenticated", false)
	}
	c.Next()
}
func authorize(c *fiber.Ctx) {
	if c.Locals("isAuthenticated") == false {
		c.Next()
		return
	}
	if c.Params("role") == "admin" {
		c.Locals("redirectRoute", "dashboard")
	} else if c.Params("role") == "user" {
		c.Locals("redirectRoute", "homepage/"+c.Params("userId"))
	} else {
		c.Locals("redirectRoute", "contact-support")
	}
	c.Next()
}
