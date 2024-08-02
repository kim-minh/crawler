package script

import (
	"context"
	"crawler/db"
	"crawler/stfetch"
	"crawler/utils"
	"sync"
)

func UpdateCompanies(queries *db.Queries) {
	ctx := context.Background()
	dataType := "companies"

	companies, _ := stfetch.FetchCompanies()

	var wg sync.WaitGroup
	defer utils.LogComplete(dataType)

	for _, company := range companies.Data {
		wg.Add(1)
		go func() {
			err := queries.CreateCompany(ctx, db.CreateCompanyParams{
				Ticker:      company.Ticker,
				FullnameVi:  company.FullnameVi,
				CompanyType: company.CompanyType,
				Exchange:    company.Exchange,
			})
			utils.LogInsertError(company.Ticker, dataType, err)
			defer wg.Done()
		}()
	}
	wg.Wait()
}
