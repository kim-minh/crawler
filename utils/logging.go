package utils

import (
	"log"
	"net/http"
)

func LogComplete(dataType string) {
	log.Printf("Updated %s data ", dataType)
}

func LogInsertError(ticker string, dataType string, err error) {
	if err != nil {
		log.Printf("Failed to update %s %s. %s", ticker, dataType, err)
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
