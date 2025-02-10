package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	//Set port from env var
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to 8080 if PORT is not set
	}

	//create new Fiber app
	app := fiber.New()

	//enable CORS to allow requests to backend
	app.Use(cors.New())

	//routing
	app.Get("/", homePage)

	//port listen and serve
	err := app.Listen(":" + port)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}

}

func homePage(c *fiber.Ctx) error {
	return c.SendString("Hi this is the home page of the song recommendation web app. The Github repo can be found at: https://github.com/huynhli/song-recommendations-web-app")
}
