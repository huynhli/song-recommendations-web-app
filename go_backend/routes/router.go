package routes

import (
	"go_backend/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Register routes for authentication
	app.Get("/", homePage)
	app.Get("/data", getGenres)
	app.Get("/api/data", handlers.GetRecommendationsAPI) //TODO change to GetRecommendationsAPI
}

func homePage(c *fiber.Ctx) error {

	return c.SendString("Hi this is the home page of the song recommendation web app. The Github repo can be found at: https://github.com/huynhli/song-recommendations-web-app")
}

func getGenres(c *fiber.Ctx) error {
	// This is the handler that will be called when the frontend sends a GET request to /data
	data := []string{"pop", "jazz", "blues", "rock"}
	return c.JSON(data)
}
