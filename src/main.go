package main

import (
	"authentication/rest-endpoint/src/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	// Sign Up Route
	app.Post("/signup", func(c *fiber.Ctx) error {
		var user services.User

		err := c.BodyParser(&user)

		if err != nil {
			return err
		}

		res := user.SignUp()

		return c.SendString(res)
	})

	// Login Route
	app.Post("/login", func(c *fiber.Ctx) error {
		var user services.User

		err := c.BodyParser(&user)

		if err != nil {
			return err
		}

		res := user.Login()

		return c.SendString(res)
	})

	// Logout Route
	app.Delete("/logout", func(c *fiber.Ctx) error {
		var user services.User

		err := c.BodyParser(&user)

		if err != nil {
			return err
		}

		res := user.LogOut()

		return c.SendString(res)
	})

	app.Listen(":3000")
}
