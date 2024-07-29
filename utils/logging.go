package utils

import (
	"log"
	"net/http"
)

func LogComplete(err error, dataType string) {
	if err == nil {
		log.Printf("Update %s data sucessfully", dataType)
	} else {
		log.Printf("Failed to update %s data. Error: %s", dataType, err)
	}
}

func LogFetchError(res *http.Response, err error) {
	if err != nil {
		log.Printf("Can't fetch %s. Error: %s", res.Request.URL, err)
	} else if res.StatusCode != http.StatusOK {
		log.Printf("Request %s failed with %s", res.Request.URL.Path, res.Status)
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

func LogFatal(err error) {
	if err != nil {
		log.Fatalf("Exited. Fatal error: %s", err)
	}
}
