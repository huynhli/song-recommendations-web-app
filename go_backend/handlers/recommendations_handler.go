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
			// https://open.spotify.com/playlist/73rtVYCu7hFAUJVrXMdFb3?si=5776c11ba60d4868
			// https://open.spotify.com/playlist/37i9dQZF1DX0XUsuxWHRQd?si=c66a76c2239f40de
			playlistIDs = append(playlistIDs, "73rtVYCu7hFAUJVrXMdFb3")
			supportedGenresList = append(supportedGenresList, "rap")
		case "hip hop":
			// https://open.spotify.com/playlist/73rtVYCu7hFAUJVrXMdFb3?si=0fba70e46e2742b9
			playlistIDs = append(playlistIDs, "73rtVYCu7hFAUJVrXMdFb3")
			supportedGenresList = append(supportedGenresList, "hip hop")
		case "pop":
			playlistIDs = append(playlistIDs, "")
			supportedGenresList = append(supportedGenresList, "pop")
		case "rock":
			playlistIDs = append(playlistIDs, "")
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
