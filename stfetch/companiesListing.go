package stfetch

import (
	"crawler/utils"
)

func FetchCompanies() (company, error) {
	url := "https://wifeed.vn/api/thong-tin-co-phieu/danh-sach-ma-chung-khoan"
	data, err := utils.Fetch[company](url)
	return data, err
}
