package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Register routes for authentication
	app.Get("/", homePage)
	app.Get("/data", getGenres)
}

func homePage(c *fiber.Ctx) error {

	return c.SendString("Hi this is the home page of the song recommendation web app. The Github repo can be found at: https://github.com/huynhli/song-recommendations-web-app")
}

func getGenres(c *fiber.Ctx) error {
	// This is the handler that will be called when the frontend sends a GET request to /data
	data := []string{"pop", "jazz", "blues", "rock"}
	return c.JSON(data)
}
