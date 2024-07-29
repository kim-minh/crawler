package script

import (
	"context"
	"crawler/db"
	"crawler/stfetch"
	"crawler/utils"
)

func UpdateCompanies(queries *db.Queries) {
	ctx := context.Background()

	companies, _ := stfetch.FetchCompanies()
	var err error
	for _, company := range companies.Data {
		err = queries.CreateCompany(ctx, db.CreateCompanyParams{
			Ticker:      company.Ticker.String,
			FullnameVi:  company.FullnameVi,
			CompanyType: company.CompanyType,
			Exchange:    company.Exchange,
		})
	}
	defer utils.LogComplete(err, "companies")
}
