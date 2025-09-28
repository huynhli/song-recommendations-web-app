package handlers

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestInToOut(t *testing.T) {
	app := fiber.New()
	app.Get("/test", inToOut)

	req := httptest.NewRequest("GET", "/test?link=https://example.com/type/123", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)

	body, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "type", string(body))
}
