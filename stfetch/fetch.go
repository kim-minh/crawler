package stfetch

import (
	"crawler/utils"
	"encoding/json"
	"io"
	"net/http"
)

func fetch[T any](url string) (T, error) {
	res, err := http.Get(url)
	utils.LogFetchError(res, err)

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	utils.LogRequestBodyError(err)

	var data T
	err = json.Unmarshal(body, &data)
	utils.LogJsonError(err)

	return data, err
}
