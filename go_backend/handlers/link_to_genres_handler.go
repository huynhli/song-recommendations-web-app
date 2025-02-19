package handlers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"go_backend/config"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func GetGenreAPI(c *fiber.Ctx) ([]string, string) {
	link := c.Query("link")

	typeOf := ""
	var spotifyID strings.Builder
	tempList := []string{"This is a valid link.", "This is not a valid link. Try again.", "This is not a valid spotify link. Try again.", "API process"}
	if link == "" {
		return tempList[1:2], ""
	} else {

		if strings.Contains(link, "https://open.spotify.com/") {
			typeOf = decideTypeOfAPICall(link, typeOf, spotifyID)
		}

		// if no assignment, link is broken, return so
		if typeOf == "" || spotifyID.Len() == 0 {
			return tempList[2:3], ""
		}
		// link is good, so make token for api
		accessToken := generateToken()

		// make api request, return json data -> genres, artists, artists top tracks
		returningData := apiRequest(accessToken, spotifyID.String(), typeOf)
		return returningData, accessToken
	}

	// TODO make new request, send to loading screen
}

func decideTypeOfAPICall(link string, typeOf string, spotifyID strings.Builder) string {
	if strings.Contains(link[25:], "artist") {
		typeOf = "artist"
		for _, each := range link[32:] {
			if each == '?' {
				break
			}
			spotifyID.WriteString(string(each))
		}
	} else if strings.Contains(link[25:], "track") {
		typeOf = "track"
		for _, each := range link[31:] {
			if each == '?' {
				break
			}
			spotifyID.WriteString(string(each))
		}
	} else if strings.Contains(link[25:], "album") {
		typeOf = "album"
		for _, each := range link[31:] {
			if each == '?' {
				break
			}
			spotifyID.WriteString(string(each))
		}
	} else if strings.Contains(link[25:], "playlist") {
		typeOf = "playlist"
		for _, each := range link[34:] {
			if each == '?' {
				break
			}
			spotifyID.WriteString(string(each))
		}
	}
	return typeOf
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
	body := initHTTPandReadResponse(req)

	// Parse the JSON response into the TokenResponse struct
	var tokenResponse TokenResponse
	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	// Get the access token from the response
	return tokenResponse.AccessToken
}

func initHTTPandReadResponse(req *http.Request) []byte {
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
	return body
}

func apiRequest(accessToken string, spotifyID string, typeOf string) []string {
	// API request
	// Description:
	// "genres" is deprecated, so the solution is by cases
	// if given artist:
	//    find genres of artist
	// elif given song or album:
	//    find artist(s), then genre of artist(s)
	// else given playlist:
	//    find artists most frequently occuring -> dict, save top frequent 10 genres -> dict
	// then, find genres of browse playlist if exist

	switch typeOf {
	case "artist":
		var Artist struct {
			Genres []string `json:"genres"`
		}
		return artistCase(spotifyID, accessToken, Artist)

	case "track":
		var Track struct {
			Artists []struct {
				ArtistID string `json:"id"`
			} `json:"artists"`
			IsLocal bool `json:"is_local"`
		}
		shouldReturn, result := trackCase(spotifyID, accessToken, Track)
		if shouldReturn {
			return result
		}

		var totalGenresList []string
		for _, eachArtist := range Track.Artists {
			var Artist struct {
				Genres []string `json:"genres"`
			}
			totalGenresList = trackCaseNested(eachArtist, accessToken, Artist, totalGenresList)
		}
		return totalGenresList

	case "album":
		var Album struct {
			Items []struct {
				TrackID string `json:"id"`
			} `json:"items"`
		}
		albumCase(spotifyID, accessToken, Album)

		totalGenresMap := make(map[string]struct{})
		totalGenresList := []string{}
		for _, eachTrack := range Album.Items {
			var Track struct {
				Artists []struct {
					ArtistID string `json:"id"`
				} `json:"artists"`
				IsLocal bool `json:"is_local"`
			}
			shouldReturn, result := validateTrack(eachTrack, accessToken, Track)
			if shouldReturn {
				return result
			}

			for _, eachArtist := range Track.Artists {
				var Artist struct {
					Genres []string `json:"genres"`
				}
				albumCaseNestedTwice(eachArtist, accessToken, Artist, totalGenresMap, totalGenresList)
			}
		}
		return totalGenresList

	case "playlist":
		var Playlist struct {
			TrackList []struct {
				Track struct {
					TrackID string `json:"id"`
				} `json:"track"`
			} `json:"items"`
		}
		playlistCase(spotifyID, accessToken, Playlist)

		totalGenresMap := make(map[string]struct{})
		totalGenresList := []string{}
		for _, eachTrack := range Playlist.TrackList {
			var Track struct {
				Artists []struct {
					ArtistID string `json:"id"`
				} `json:"artists"`
				IsLocal bool `json:"is_local"`
			}
			shouldReturn, result := validateTrack(eachTrack.Track, accessToken, Track)
			if shouldReturn {
				return result
			}

			for _, eachArtist := range Track.Artists {
				// create get request to link
				var Artist struct {
					Genres []string `json:"genres"`
				}
				body := apiCall("artists", eachArtist.ArtistID, accessToken)
				artistToGenres(body, Artist)
				for _, eachItem := range Artist.Genres {
					addUnique(totalGenresMap, eachItem, &totalGenresList)
				}
			}
		}
		return totalGenresList
	}
	tempReturn := []string{"This has an issue."}
	return tempReturn
}

func playlistCase(spotifyID string, accessToken string, Playlist struct {
	TrackList []struct {
		Track struct {
			TrackID string "json:\"id\""
		} "json:\"track\""
	} "json:\"items\""
}) {
	body := apiCall("playlists", spotifyID, accessToken)
	var err = json.Unmarshal(body, &Playlist)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}
}

func albumCaseNestedTwice(eachArtist struct {
	ArtistID string "json:\"id\""
}, accessToken string, Artist struct {
	Genres []string "json:\"genres\""
}, totalGenresMap map[string]struct{}, totalGenresList []string) {
	body := apiCall("artists", eachArtist.ArtistID, accessToken)
	artistToGenres(body, Artist)
	for _, eachItem := range Artist.Genres {
		addUnique(totalGenresMap, eachItem, &totalGenresList)
	}
}

func validateTrack(eachTrack struct {
	TrackID string "json:\"id\""
}, accessToken string, Track struct {
	Artists []struct {
		ArtistID string "json:\"id\""
	} "json:\"artists\""
	IsLocal bool "json:\"is_local\""
}) (bool, []string) {
	body := apiCall("tracks", eachTrack.TrackID, accessToken)
	var err = json.Unmarshal(body, &Track)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	localReturn := []string{"This is a local track. Try a non-local track."}
	if Track.IsLocal {
		return true, localReturn
	}
	return false, nil
}

func albumCase(spotifyID string, accessToken string, Album struct {
	Items []struct {
		TrackID string "json:\"id\""
	} "json:\"items\""
}) {
	body := apiCall("albums", spotifyID, accessToken)
	var err = json.Unmarshal(body, &Album)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}
}

func trackCaseNested(eachArtist struct {
	ArtistID string "json:\"id\""
}, accessToken string, Artist struct {
	Genres []string "json:\"genres\""
}, totalGenresList []string) []string {
	body := apiCall("artists", eachArtist.ArtistID, accessToken)
	artistToGenres(body, Artist)
	totalGenresList = append(totalGenresList, Artist.Genres...)
	return totalGenresList
}

func trackCase(spotifyID string, accessToken string, Track struct {
	Artists []struct {
		ArtistID string "json:\"id\""
	} "json:\"artists\""
	IsLocal bool "json:\"is_local\""
}) (bool, []string) {
	body := apiCall("tracks", spotifyID, accessToken)

	var err = json.Unmarshal(body, &Track)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	localReturn := []string{"This is a local track. Try a non-local track."}
	if Track.IsLocal {
		return true, localReturn
	}
	return false, nil
}

func artistCase(spotifyID string, accessToken string, Artist struct {
	Genres []string "json:\"genres\""
}) []string {
	body := apiCall("artists", spotifyID, accessToken)
	artistToGenres(body, Artist)
	return Artist.Genres
}

func artistToGenres(body []byte, Artist struct {
	Genres []string "json:\"genres\""
}) {
	var err = json.Unmarshal(body, &Artist)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}
}

func apiCall(decider string, spotifyID string, accessToken string) []byte {
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/"+decider+"/"+spotifyID, nil)
	if err != nil {
		log.Fatalf("Error creating GET: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	// send request to HTTP client instance
	body := initHTTPandReadResponse(req)
	return body
}

func addUnique(uniqueSet map[string]struct{}, item string, finalList *[]string) {
	if _, exists := uniqueSet[item]; !exists {
		uniqueSet[item] = struct{}{}
		*finalList = append(*finalList, item)
	}
}
