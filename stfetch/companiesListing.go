package stfetch

import (
	"crawler/utils"
)

func ListCompanies() (company, int) {
	url := "https://wifeed.vn/api/thong-tin-co-phieu/danh-sach-ma-chung-khoan"
	data, status := utils.Fetch[company](url)
	return data, status
}
