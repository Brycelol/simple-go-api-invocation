package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"simple-go-api-invocation/model"
)

const baseURL = "https://jsonmock.hackerrank.com/api/tvseries?page="

func main() {
	log.Println("Launching Simple Go API Invocation...")

	var singlePage model.TVSeriesPage

	log.Printf("Invoking API at URL : [%s]\n", baseURL)
	resp, err := http.Get(baseURL)

	if err != nil {
		log.Fatalf("An error occured attempting to invoke URL [%s]", baseURL)
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatalf("An error occurred attempting to close response body [%s]", err)
		}
	}()

	if resp != nil {
		log.Printf("Got Response Code : [%d]\n", resp.StatusCode)
		if resp.StatusCode != 200 {
			log.Fatalf("Unexpected response code from API. We got status code [%d]", resp.StatusCode)
		}

		body, ioErr := io.ReadAll(resp.Body)
		if ioErr != nil {
			log.Fatalf("An error occurred attempting to read response body [%d]", ioErr)
		}

		umErr := json.Unmarshal(body, &singlePage)
		if umErr != nil {
			log.Fatalf("An error occurred attempting to unmarshall response data [%d]", ioErr)
		}

		log.Printf("We successfully downloaded a single page and found %d TV series.\n", len(singlePage.Data))
	}
}
