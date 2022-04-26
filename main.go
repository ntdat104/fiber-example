package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World 👋!")
    })

	app.Get("/:name", handleName)

	app.Get("/:name/:age", handleNameAndAge)

	app.Get("/api/*", handleWildcard)

    app.Listen(":3000")
}

func handleName(c*fiber.Ctx) error {
	name := c.Params("name")
	return c.SendString("Hello, " + name + "👋!")
}

func handleNameAndAge(c*fiber.Ctx) error {
	name := c.Params("name")
	age := c.Params("age")
	return c.SendString("Hello, " + name + "👋! You are " + age + " years old.")
}

func handleWildcard(c*fiber.Ctx) error {
	param := c.Params("*")
	return c.SendString("You are on the api route. " + param)
}