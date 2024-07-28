package stfetch

import (
	"crawler/utils"
	"encoding/json"
)

func ListCompanies() (company, int) {
	url := "https://wifeed.vn/api/thong-tin-co-phieu/danh-sach-ma-chung-khoan"
	body, status := utils.Fetch(url)

	var data company
	err := json.Unmarshal(body, &data)
	utils.LogJsonError(err)

	return data, status
}
