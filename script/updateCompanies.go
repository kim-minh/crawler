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

	companies, _ := stfetch.FetchCompanies()

	var wg sync.WaitGroup
	c := make(chan utils.UpdateError)
	defer utils.LogComplete(c, "companies")

	for _, company := range companies.Data {
		wg.Add(1)
		go func() {
			err := queries.CreateCompany(ctx, db.CreateCompanyParams{
				Ticker:      company.Ticker,
				FullnameVi:  company.FullnameVi,
				CompanyType: company.CompanyType,
				Exchange:    company.Exchange,
			})
			if err != nil {
				c <- utils.UpdateError{Ticker: company.Ticker, Error: err}
			}
			defer wg.Done()
		}()
	}
	wg.Wait()
	close(c)
}
