package routes

import (
	"go_backend/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Register routes for authentication
	app.Get("/", homePage)

	api := app.Group("/api/v1")

	lastFMAPI := api.Group("/lastfm")
	lastFMAPI.Get("/track", handlers.LastFMRecs)
	lastFMAPI.Get("/lastFMRec	/artist", handlers.LastFMRecs)
	lastFMAPI.Get("/")

	// musicBrainzAPI = api.Group("/musicBrainz")

	// deezerAPI := api.Group("/deezer")
}

func homePage(c *fiber.Ctx) error {

	return c.SendString("Hi this is the home page of a song recommendation web app. The Github repo can be found at: https://github.com/huynhli/similarSongs")
}

// func getGenreAPI(c *fiber.Ctx) error {
// 	link := c.Query("link")
// 	tempList := []string{"This is a valid link.", "This is not a valid link. Try again."}
// 	if link == "" {
// 		return c.JSON(tempList[1:])
// 	}

// 	// authOptions = []string{}

// 	return c.JSON(tempList[:1])
// }
