package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	SpotifyClientID     string
	SpotifyClientSecret string
	SpotifyRedirectURI  string
)

func LoadConfig() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set config variables from environment variables
	SpotifyClientID = os.Getenv("SPOTIFY_CLIENT_ID")
	SpotifyClientSecret = os.Getenv("SPOTIFY_CLIENT_SECRET")
	SpotifyRedirectURI = "http://localhost:8000/auth/callback" // Redirect URI for your app
}
