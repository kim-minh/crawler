package stfetch

import "fmt"

const host = "https://apipubaws.tcbs.com.vn"

func FetchOverview(ticker string) (overview, error) {
	url := fmt.Sprintf("%s/tcanalysis/v1/ticker/%s/overview", host, ticker)
	return fetch[overview](url)
}

func FetchProfile(ticker string) (profile, error) {
	url := fmt.Sprintf("%s/tcanalysis/v1/company/%s/overview", host, ticker)
	return fetch[profile](url)
}

func FetchShareholders(ticker string) (shareholders, error) {
	url := fmt.Sprintf("%s/tcanalysis/v1/company/%s/large-share-holders", host, ticker)
	return fetch[shareholders](url)
}

func FetchInsiderDeals(ticker string) (insiderdeals, error) {
	url := fmt.Sprintf("%s/tcanalysis/v1/company/%s/insider-dealing?page=0&size=20", host, ticker)
	return fetch[insiderdeals](url)
}

func FetchSubsidiares(ticker string) (subsidiaries, error) {
	url := fmt.Sprintf("%s/tcanalysis/v1/company/%s/sub-companies?page=0&size=10", host, ticker)
	return fetch[subsidiaries](url)
}

func FetchOfficers(ticker string) (officers, error) {
	url := fmt.Sprintf("%s/tcanalysis/v1/company/%s/key-officers?page=0&size=10", host, ticker)
	return fetch[officers](url)
}

func FetchEvents(ticker string) (events, error) {
	url := fmt.Sprintf("%s/tcanalysis/v1/ticker/%s/events-news?page=0&size=20", host, ticker)
	return fetch[events](url)
}

func FetchNews(ticker string) (news, error) {
	url := fmt.Sprintf("%s/tcanalysis/v1/ticker/%s/activity-news?page=0&size=20", host, ticker)
	return fetch[news](url)
}

func FetchDividends(ticker string, page int) (dividends, error) {
	url := fmt.Sprintf("%s/tcanalysis/v1/company/%s/dividend-payment-histories?page=%d&size=1000", host, ticker, page)
	return fetch[dividends](url)
}
