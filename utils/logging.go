package utils

import (
	"log"
	"net/http"
)

type UpdateError struct {
	Ticker string
	Error  error
}

func LogComplete(c <-chan UpdateError, dataType string) {
	if len(c) == 0 {
		log.Printf("Sucessfully updated %s data ", dataType)
	} else {
		for i := range c {
			log.Printf("Failed to update %s %s. Error: %s", i.Ticker, dataType, i.Error)
		}
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
