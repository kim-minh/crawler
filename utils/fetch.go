package utils

import (
	"io"
	"net/http"
)

func Fetch(url string) ([]byte, int) {
	res, err := http.Get(url)

	LogFetchError(res, err, url)

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	LogRequestBodyError(err)

	return body, res.StatusCode
}
