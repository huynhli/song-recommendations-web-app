package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetSpotifyToken fetches an access token using Client Credentials Flow
func GetSpotifyToken(clientID, clientSecret string) (string, error) {
	url := "https://accounts.spotify.com/api/token"
	data := "grant_type=client_credentials"

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data)))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(clientID, clientSecret)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal(body, &result)

	if token, ok := result["access_token"].(string); ok {
		return token, nil
	}
	return "", fmt.Errorf("failed to get access token")
}

// GetGenreSeeds fetches available genre seeds from Spotify
func GetGenreSeeds(token string) ([]string, error) {
	url := "https://api.spotify.com/v1/recommendations/available-genre-seeds"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var result map[string][]string
	json.Unmarshal(body, &result)

	if genres, ok := result["genres"]; ok {
		return genres, nil
	}
	return nil, fmt.Errorf("failed to fetch genres")
}
