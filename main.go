package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("start the server from here ")
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	
	app.Listen("localhost:5000")
}
