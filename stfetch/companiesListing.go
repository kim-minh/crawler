package stfetch

func FetchCompanies() (company, error) {
	url := "https://wifeed.vn/api/thong-tin-co-phieu/danh-sach-ma-chung-khoan"
	return fetch[company](url)
}
