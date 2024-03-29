package main

import "github.com/gofiber/fiber"

func helloworld(c *fiber.Ctx) {
	c.Send("Hello World")
}
func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.Getbook)
	app.Post("/api/v1/book", book.GetNewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func main() {
	app := fiber.New()

	setupRoutes(app)
	app.Listen(3000)

}
