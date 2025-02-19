package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetRecommendationsAPI(c *fiber.Ctx) error {
	// if an error occured
	returnedGenresAndToken := GetGenreAPI(c)
	if strings.HasPrefix(returnedGenresAndToken[0][0], "This") {
		return c.JSON(returnedGenresAndToken[0])
	}

	// else, get recommendations!
	// TODO this is deprecated -> find a new way to do this
	// TODO before fixing because of deprecation -> get recommendations based on tracks, artists.
	// genres is kinda generalized and somehow not that supported for lots of spotify things
	recommendations := []string{}
	for _, eachGenre := range returnedGenresAndToken[0] {
		req, err := http.NewRequest("GET", "https://api.spotify.com/v1/recommendations?seed_genres="+strings.ReplaceAll(eachGenre, " ", "-"), nil)
		if err != nil {
			log.Fatalf("Error getting recommendations: ", err)
		}
		req.Header.Set("Authorization", "Bearer "+returnedGenresAndToken[1])

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalf("Error sending request: %v", err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Error reading response body: %v", err)
		}
		var Recommendations struct {
			Tracks []struct {
				Artists []struct {
					ArtistName string `json:"name"`
				} `json:"artists"`
				Name string `json:"name"`
			} `json:"tracks"`
		}
		err = json.Unmarshal(body, &Recommendations)
		if err != nil {
			log.Fatalf("Error unmarshalling: %v", err)
		}

		totalSongsMap := make(map[string]struct{})
		totalSongsList := []string{}
		for index, eachTrack := range Recommendations.Tracks {
			var ArtistList = ""
			for _, eachArtist := range Recommendations.Tracks[index].Artists {

			}
			addUnique(totalSongsMap, Recommendations.Tracks[index].Name+" "+ArtistList, &totalSongsList)
		}

	}

}
