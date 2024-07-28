package stfetch

import (
	"crawler/utils"
	"encoding/json"
	"fmt"
)

func FetchOverview(ticker string) (overview, int) {
	url := fmt.Sprintf("https://apipubaws.tcbs.com.vn/tcanalysis/v1/ticker/%s/overview", ticker)
	body, status := utils.Fetch(url)

	var data overview
	err := json.Unmarshal(body, &data)
	utils.LogJsonError(err)

	return data, status
}
