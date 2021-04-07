package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Todo struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"conpleted"`
}

var todos = []Todo{
	{
		ID: 1, Name: "Learn GO", Completed: true,
	},
	{
		ID: 2, Name: "Learn Rust", Completed: false,
	},
	{
		ID: 3, Name: "Learn Typescript", Completed: true,
	},
	{
		ID: 4, Name: "Build Serverless APP", Completed: false,
	},
}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!!!!!!!")
	})

	app.Get("/todos", getTodos)

	app.Listen(":5000")
}

func getTodos(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(todos)
}
