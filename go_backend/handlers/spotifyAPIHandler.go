package handlers

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// input a artist, album, or track. get an artist, album, or track.
func inToOut(c *fiber.Ctx) error {
	var link = c.Query("link")
	fmt.Println(link)
	var typeLink = strings.Split(link, "/")[3]
	fmt.Println(typeLink)
	return c.SendString(typeLink)
}
