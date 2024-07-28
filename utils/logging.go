package utils

import (
	"log"
	"net/http"
)

func LogComplete(err error, data string) {
	if err == nil {
		log.Printf("Update %s data sucessfully", data)
	} else {
		log.Printf("Failed to update %s data", data)
	}
}

func LogFetchError(res *http.Response, err error, url string) {
	if err != nil {
		log.Printf("Can't fetch %s. Error: %s", url, err)
	} else if res.StatusCode != http.StatusOK {
		log.Printf("Request failed with %s", res.Status)
	}
}

func LogRequestBodyError(err error) {
	if err != nil {
		log.Printf("Error reading body: %s", err)
	}
}

func LogJsonError(err error) {
	if err != nil {
		log.Printf("Error creating json: %s", err)
	}
}

func LogError(err error) {
	if err != nil {
		log.Printf("Error: %s", err)
	}
}
