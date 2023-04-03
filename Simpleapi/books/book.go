package main

import (
	
	"gitub.com/shashank/Restingapi/book"
	"github.com/gofiber/fiber"
)

func GetBooks(c *fiber.Ctx) {
	c.Send("All Books")
}

func GetBook(c *fiber.Ctx) {
	c.Send("A Singular Book")
}
func GetNewBook(c *fiber.Ctx) {
	c.Send("Adds a new Book")
}
func DeleteBook(c *fiber.Ctx) {
	c.Send("Deletes a Book")
}
