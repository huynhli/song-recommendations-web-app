package main

import (
	"go_backend/config"
	"go_backend/handlers"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestInToOutArtist(t *testing.T) {
	config.LoadConfig()
	app := fiber.New()
	app.Get("/test", handlers.InToOut)

	// req := httptest.NewRequest("GET", "/test?link=https://open.spotify.com/track/0TwBtDAWpkpM3srywFVOV5?si=2c1aa537c2fe4340", nil)
	req := httptest.NewRequest("GET", "/test?link=https://open.spotify.com/artist/6vWDO969PvNqNYHIOW5v0m?si=i8pS1ntxTF-tCTV1rItksw", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)

	// Check that the JSON contains the expected fields
	assert.Contains(t, string(body), `"artist"`)
	// assert.Contains(t, string(body), `"track"`)
	// assert.Contains(t, string(body), `"album"`)
}

func TestInToOutAlbum(t *testing.T) {
	config.LoadConfig()
	app := fiber.New()
	app.Get("/test", handlers.InToOut)

	// req := httptest.NewRequest("GET", "/test?link=https://open.spotify.com/track/0TwBtDAWpkpM3srywFVOV5?si=2c1aa537c2fe4340", nil)
	req := httptest.NewRequest("GET", "/test?link=https://open.spotify.com/album/6BzxX6zkDsYKFJ04ziU5xQ?si=cchby_lyTES2MnA1L2kleg", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)

	// Check that the JSON contains the expected fields
	assert.Contains(t, string(body), `"artist"`)
	// assert.Contains(t, string(body), `"track"`)
	assert.Contains(t, string(body), `"album"`)
}

func TestInToOutTrack(t *testing.T) {
	config.LoadConfig()
	app := fiber.New()
	app.Get("/test", handlers.InToOut)

	req := httptest.NewRequest("GET", "/test?link=https://open.spotify.com/track/0TwBtDAWpkpM3srywFVOV5?si=2c1aa537c2fe4340", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)

	// Check that the JSON contains the expected fields
	assert.Contains(t, string(body), `"artist"`)
	assert.Contains(t, string(body), `"track"`)
	assert.Contains(t, string(body), `"album"`)
}

// TODO more tests
