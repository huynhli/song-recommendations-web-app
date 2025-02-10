package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to 8080 if PORT is not set
	}

	print("Welcome to the song recommendation web app!")

	//handler -> (url, function to handle)
	http.HandleFunc("/", homePage)

	//port listen and serve
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		print(err.Error())
	}

}

func homePage(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Hi this is the home page of the song recommendation web app. The Github repo can be found at: https://github.com/huynhli/song-recommendations-web-app")
}
