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
	app.Post("/todo", createTodo)

	app.Listen(":5000")
}

func getTodos(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(todos)
}

func createTodo(c *fiber.Ctx) error {
	type request struct {
		Name      string `json:"name"`
		Completed bool   `json:"completed"`
	}

	var body request

	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad Request",
		})
	}

	todo := Todo{
		ID:        len(todos) + 1,
		Name:      body.Name,
		Completed: body.Completed,
	}

	todos = append(todos, todo)

	return c.Status(fiber.StatusCreated).JSON(todo)
}
