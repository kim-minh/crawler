package stfetch

import (
	"crawler/utils"
	"fmt"
)

func FetchOverview(ticker string) (overview, int) {
	url := fmt.Sprintf("https://apipubaws.tcbs.com.vn/tcanalysis/v1/ticker/%s/overview", ticker)
	data, status := utils.Fetch[overview](url)
	return data, status
}
