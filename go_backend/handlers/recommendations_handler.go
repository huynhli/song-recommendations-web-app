package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetRecommendationsAPI(c *fiber.Ctx) error {
	// if an error occured
	returnedGenres, token := GetGenreFromAPI(c)

	if strings.HasPrefix(returnedGenres[0], "This") {
		return c.JSON(returnedGenres[0])
	}
	// else, get recommendations!
	// TODO this is deprecated -> find a new way to do this

	totalSongsMap := make(map[string]struct{})
	totalSongsList := []string{}

	supportedGenresList := []string{}
	playlistIDs := []string{}
	for _, eachGenre := range returnedGenres {
		switch eachGenre {
		case "rap":
			// https://open.spotify.com/playlist/37i9dQZF1DX0XUsuxWHRQd?si=c66a76c2239f40de
			playlistIDs = append(playlistIDs, "37i9dQZF1DX0XUsuxWHRQd")
			supportedGenresList = append(supportedGenresList, "rap")
		case "hip hop":
			playlistIDs = append(playlistIDs, "")
			supportedGenresList = append(supportedGenresList, "hip hop")
		case "pop":
			playlistIDs = append(playlistIDs, "")
			supportedGenresList = append(supportedGenresList, "pop")
		case "rock":
			playlistIDs = append(playlistIDs, "")
			supportedGenresList = append(supportedGenresList, "rock")
		}
	}
	// req, err := http.NewRequest("GET", "https://api.spotify.com/v1/recommendations?seed_genres="+strings.ReplaceAll(eachGenre, " ", "-"), nil)
	// if err != nil {
	// 	log.Fatalf("Error getting recommendations: %v", err)
	// }
	// req.Header.Set("Authorization", "Bearer "+token)

	// client := &http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	// 	log.Fatalf("Error sending request: %v", err)
	// }
	// defer resp.Body.Close()

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatalf("Error reading response body: %v", err)
	// }
	// var Recommendations struct {
	// 	Tracks []struct {
	// 		Artists []struct {
	// 			Name string `json:"name"`
	// 		} `json:"artists"`
	// 		Name string `json:"name"`
	// 	} `json:"tracks"`
	// }

	// err = json.Unmarshal(body, &Recommendations)
	// if err != nil {
	// 	log.Fatalf("Error unmarshalling: %v", err)
	// }

	// for index, eachTrack := range Recommendations.Tracks {
	// 	var ArtistList = ""
	// 	for indexArtist, eachArtist := range Recommendations.Tracks[index].Artists {
	// 		if indexArtist == 0 {
	// 			ArtistList = ArtistList + " by " + eachArtist.Name
	// 			continue
	// 		}
	// 		ArtistList = ArtistList + " + " + eachArtist.Name
	// 	}
	// 	addUnique(totalSongsMap, eachTrack.Name+" "+ArtistList, &totalSongsList)
	// }
	for _, eachPlaylistId := range playlistIDs {
		req, err := http.NewRequest("GET", "https://api.spotify.com/v1/playlists/"+eachPlaylistId+"/tracks", nil)
		fmt.Println("url is: ", "https://api.spotify.com/v1/playlists/"+eachPlaylistId+"/tracks?limit=5")
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
		fmt.Println("body is:", string(body))

		for _, eachItem := range Playlist.Items {
			var ArtistList = ""
			for indexArtist, eachArtist := range eachItem.TrackObject.Artists {
				if indexArtist == 0 {
					ArtistList = ArtistList + " by " + eachArtist.Name
					continue
				}
				ArtistList = ArtistList + " + " + eachArtist.Name
			}
			addUnique(totalSongsMap, eachItem.Name+" "+ArtistList, &totalSongsList)
		}
	}

	returnedLength := len(returnedGenres)
	returnedGenres = append(returnedGenres, supportedGenresList...)
	returnedGenres = append(returnedGenres, strconv.Itoa(returnedLength))
	fmt.Println(totalSongsList)
	fmt.Println(returnedGenres)
	return c.JSON(returnedGenres)

}
