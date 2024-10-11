package main

import (
	// third-party libraries
	"log"
	"net/http"
	"tester/telegram/tools"
)

func main() {
	urlAdd := "https://www.google.com"

	contentType := "application/x-www-form-urlencoded"

	client := &http.Client{}
	res, err := tools.Get(client, urlAdd, contentType, nil)
	if err != nil {
		// handle error
		panic(err)
	}

	log.Printf("ntework response body: %s", res)
}
