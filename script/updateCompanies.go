package script

import (
	"context"
	"crawler/db"
	"crawler/stfetch"
	"crawler/utils"
)

func UpdateCompanies(queries *db.Queries) {
	ctx := context.Background()

	companies, _ := stfetch.ListCompanies()
	for _, company := range companies.Data {
		err := queries.CreateCompany(ctx, db.CreateCompanyParams{
			Ticker:      company.Ticker.String,
			FullnameVi:  company.FullnameVi,
			CompanyType: company.CompanyType,
			Exchange:    company.Exchange,
		})

		if err != nil {
			utils.LogError(err)
		}
		defer utils.LogComplete(err, "companies")
	}
}
