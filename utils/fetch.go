package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func Fetch[T any](url string) (T, int) {
	res, err := http.Get(url)

	LogFetchError(res, err, url)

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	LogRequestBodyError(err)

	var data T
	err = json.Unmarshal(body, &data)
	LogJsonError(err)

	return data, res.StatusCode
}
