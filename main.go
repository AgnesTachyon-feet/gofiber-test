package main

import (
	"github.com/gofiber/fiber/v2"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author`
}

var books []Book

func main() {
	app := fiber.New()

	books = append(books, Book{ID: 1, Tile: "1984", Author: "George Orwell"})
	books = append(books, Book{ID: 2, Tile: "The Great Gatsby", Author: "F. Soctt Fitzgerald"})

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	app.Listen(":8080")
}
