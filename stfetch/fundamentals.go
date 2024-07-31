package stfetch

import (
	"crawler/utils"
	"fmt"
)

const host = "https://apipubaws.tcbs.com.vn"

func FetchOverview(ticker string) (overview, error) {
	url := fmt.Sprintf("%s/tcanalysis/v1/ticker/%s/overview", host, ticker)
	data, err := utils.Fetch[overview](url)
	return data, err
}

func FetchProfile(ticker string) (profile, error) {
	url := fmt.Sprintf("%s/tcanalysis/v1/company/%s/overview", host, ticker)
	data, err := utils.Fetch[profile](url)
	return data, err
}

func FetchShareholders(ticker string) (shareholders, error) {
	url := fmt.Sprintf("%s/tcanalysis/v1/company/%s/large-share-holders", host, ticker)
	data, err := utils.Fetch[shareholders](url)
	return data, err
}

func FetchInsiderDeals(ticker string) (insiderdeals, error) {
	url := fmt.Sprintf("%s/tcanalysis/v1/company/%s/insider-dealing?page=0&size=20", host, ticker)
	data, err := utils.Fetch[insiderdeals](url)
	return data, err
}

func FetchSubsidiares(ticker string) (subsidiaries, error) {
	url := fmt.Sprintf("%s/tcanalysis/v1/company/%s/sub-companies?page=0&size=20", host, ticker)
	data, err := utils.Fetch[subsidiaries](url)
	return data, err
}

func FetchOfficers(ticker string) (officers, error) {
	url := fmt.Sprintf("%s/tcanalysis/v1/company/%s/key-officers?page=0&size=10", host, ticker)
	data, err := utils.Fetch[officers](url)
	return data, err
}

func FetchEvents(ticker string) (events, error) {
	url := fmt.Sprintf("%s/tcanalysis/v1/ticker/%s/events-news?page=0&size=20", host, ticker)
	data, err := utils.Fetch[events](url)
	return data, err
}

func FetchNews(ticker string) (news, error) {
	url := fmt.Sprintf("%s/tcanalysis/v1/ticker/%s/activity-news?page=0&size=20", host, ticker)
	data, err := utils.Fetch[news](url)
	return data, err
}

func FetchDividends(ticker string, page int) (dividends, error) {
	url := fmt.Sprintf("%s/tcanalysis/v1/company/%s/dividend-payment-histories?page=%d&size=1000", host, ticker, page)
	data, err := utils.Fetch[dividends](url)
	return data, err
}
