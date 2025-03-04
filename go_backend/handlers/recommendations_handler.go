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
	returnedGenres, token := GetGenreFromAPI(c)
	if len(returnedGenres) == 0 {
		temp := []string{"This artist/track/album has not been categorized by Spotify. Or, the playlist you entered does not have any tracks."}
		return c.JSON(temp)
	} else if strings.HasPrefix(returnedGenres[0], "This") {
		return c.JSON(returnedGenres)
	}
	// else, get recommendations!
	// TODO this is deprecated -> find a new way to do this

	totalSongsMap := make(map[string]struct{})
	totalSongsList := []string{}

	supportedGenresList := []string{}
	playlistIDs := []string{}
	for _, eachGenre := range returnedGenres {
		// TODO throw playlist links in a file and call dynamically
		switch eachGenre {
		case "rap":
			playlistIDs = append(playlistIDs, "7k8E4tmsRVxLSQitB9bYPQ")
			playlistIDs = append(playlistIDs, "01MRi9jFGeSEEttKOk7VgR")
			playlistIDs = append(playlistIDs, "23bBJex1i75b9oqUo6jUCC")
			playlistIDs = append(playlistIDs, "7JIGfa0KkCTDxUPOQySODP")
			supportedGenresList = append(supportedGenresList, "rap")
		case "hip hop":
			playlistIDs = append(playlistIDs, "62y3BHKehWnb1hlaPclDAA")
			playlistIDs = append(playlistIDs, "5oN7X3cPTUkOJFPmVx5wCE")
			playlistIDs = append(playlistIDs, "0dMexqq0XIWS3QJ74z3ZhD")
			playlistIDs = append(playlistIDs, "6WnTjZBYrYOdPl4pK7PyZg")
			supportedGenresList = append(supportedGenresList, "hip hop")
		case "pop":
			playlistIDs = append(playlistIDs, "1WH6WVBwPBz35ZbWsgCpgr")
			playlistIDs = append(playlistIDs, "3JoHkM90TXzfIS1RMN0Cgd")
			playlistIDs = append(playlistIDs, "2L2HwKRvUgBv1YetudaRI3")
			playlistIDs = append(playlistIDs, "2Sxd5BovYwLgRg6KyZOTer")
			supportedGenresList = append(supportedGenresList, "pop")
		case "rock":
			playlistIDs = append(playlistIDs, "1Kgkkup7w9qvGxGJGa75PS")
			playlistIDs = append(playlistIDs, "77RvyLiqmUimojxq3vg6mY")
			playlistIDs = append(playlistIDs, "23hD5D7bvXtkJGz2ni7s9e")
			playlistIDs = append(playlistIDs, "7KahejEXeb27cRJCFt3VFO")
			supportedGenresList = append(supportedGenresList, "rock")
		}
	}

	for _, eachPlaylistId := range playlistIDs {
		req, err := http.NewRequest("GET", "https://api.spotify.com/v1/playlists/"+eachPlaylistId+"/tracks?offset=0&limit=5", nil)
		if err != nil {
			log.Fatalf("Error getting recommendations: %v", err)
		}
		req.Header.Set("Authorization", "Bearer "+token)

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
		var Playlist struct {
			Items []struct {
				TrackObject struct {
					Artists []struct {
						Name string `json:"name"`
					} `json:"artists"`
					Name string `json:"name"`
				} `json:"track"`
				Name string `json:"name"`
			} `json:"items"`
		}

		err = json.Unmarshal(body, &Playlist)
		if err != nil {
			log.Fatalf("Error unmarshalling: %v", err)
		}

		for _, eachItem := range Playlist.Items {
			var ArtistList = ""
			for indexArtist, eachArtist := range eachItem.TrackObject.Artists {
				if indexArtist == 0 {
					ArtistList = ArtistList + " by " + eachArtist.Name
					continue
				}
				ArtistList = ArtistList + " + " + eachArtist.Name
			}
			addUnique(totalSongsMap, eachItem.TrackObject.Name+" "+ArtistList+" ", &totalSongsList)
		}
	}
	totalSongsList = append(totalSongsList, strings.Join(returnedGenres, ", ")+".", strings.Join(supportedGenresList, ", ")+".")
	return c.JSON(totalSongsList)

}
