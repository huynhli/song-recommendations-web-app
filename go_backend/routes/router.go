package routes

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"go_backend/config"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Register routes for authentication
	app.Get("/", homePage)
	app.Get("/data", getGenres)
	app.Get("/api/data", getGenreAPI)
}

func homePage(c *fiber.Ctx) error {

	return c.SendString("Hi this is the home page of the song recommendation web app. The Github repo can be found at: https://github.com/huynhli/song-recommendations-web-app")
}

func getGenres(c *fiber.Ctx) error {
	// This is the handler that will be called when the frontend sends a GET request to /data
	data := []string{"pop", "jazz", "blues", "rock"}
	return c.JSON(data)
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func getGenreAPI(c *fiber.Ctx) error {
	link := c.Query("link")

	typeOf := ""
	var spotifyID strings.Builder
	tempList := []string{"This is a valid link.", "This is not a valid link. Try again.", "Not a valid spotify link. Try again.", "API process"}
	if link == "" {
		return c.JSON(tempList[1:2])
	} else {

		if strings.Contains(link, "https://open.spotify.com/") {
			if strings.Contains(link[25:], "artist") {
				typeOf = "artist"
				for _, each := range(link[25:]) {
					if each == '?' {
						break
					}
					spotifyID.WriteString(string(each))
				}
			} else if strings.Contains(link[25:], "track") {
				typeOf = "track"
				for _, each := range(link[25:]) {
					if each == '?' {
						break
					}
					spotifyID.WriteString(string(each))
				}
			} else if strings.Contains(link[25:], "album") {
				typeOf = "album"
				for _, each := range(link[25:]) {
					if each == '?' {
						break
					}
					spotifyID.WriteString(string(each))
				}
			} else if strings.Contains(link[25:], "playlist") {
				typeOf = "playlist"
				for _, each := range(link[25:]) {
					if each == '?' {
						break
					}
					spotifyID.WriteString(string(each))
				}
			}
		}

		// if no assignment, link is broken, return so
		if (typeOf == "" || spotifyID.Len() == 0) {
			return c.JSON(tempList[2:3])
		}

		// if given artist: https://open.spotify.com/artist/6Xktu0x9IXB4ghFSPw6Jqv?si=vpU3HtylQTWYXDF9wZ95DA
		// 	//    find genres of artist, save max 3
		// 	// elif given song or album:https://open.spotify.com/track/0gNpXNiopu6nXKRPnfQ89E?si=1ca3023bb7da475d
		// https://open.spotify.com/album/5FFviHXLHrtM8bPkklaXrD?si=RLEszY8eQNuy4LP7-kSl_A
		// 	//    find artist(s), then genre of artist(s), save 3 each
		// 	// else given playlist:
		// https://open.spotify.com/playlist/18Mk8tJFfZdj1XvtDL9Bom?si=2b2d554c96a44807
		// check is playlist, check !is_local, use artists -> 3 genres each max
		// 	//    find artists most frequently occuring -> dict, save top frequent 10 genres -> dict
	}

	// TODO make new request, send to loading screen

	// link is good, so make token for api
	accessToken := generateToken()

	// make api request, return json data -> genres, artists, artists top tracks
	apiRequest(accessToken, spotifyID.String(), typeOf)
	return c.JSON(tempList[3:], accessToken)
}

func generateToken() string {
	clientID := config.SpotifyClientID
	clientSecret := config.SpotifyClientSecret

	authString := clientID + ":" + clientSecret
	authBase64 := base64.StdEncoding.EncodeToString([]byte(authString))

	// Set the URL for token request
	url := "https://accounts.spotify.com/api/token"

	// Prepare the POST request body (form data)
	data := "grant_type=client_credentials"

	// Create a new POST request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data)))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// Set the required headers
	req.Header.Set("Authorization", "Basic "+authBase64)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Initialize the HTTP client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	// Parse the JSON response into the TokenResponse struct
	var tokenResponse TokenResponse
	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	// Get the access token from the response
	return tokenResponse.AccessToken
}

func apiRequest(accessToken string, spotifyID string, typeOf string) {
	// API request
	// Description:
	// "genres" is deprecated, so the solution is by cases
	// if given artist:
	//    find genres of artist, save max 3
	// elif given song or album:
	//    find artist(s), then genre of artist(s), save 3 each
	// else given playlist:
	//    find artists most frequently occuring -> dict, save top frequent 10 genres -> dict
	// then, find genres of browse playlist if exist

}
