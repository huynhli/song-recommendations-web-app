package main

import (
	"os"
)

func main() {
	engine := html.New("./views", ".tmpl")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")

}