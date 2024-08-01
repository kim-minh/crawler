package script

import (
	"context"
	"crawler/db"
	"crawler/utils"
	"sync"
	"time"
)

const (
	maxConcurrent = 50
	waitTime      = 30 * time.Second
)

func updateEach(queries *db.Queries, dataType string, update func(company db.Company, c chan<- utils.UpdateError)) {
	ctx := context.Background()

	companies, err := queries.ListCompanies(ctx)
	utils.LogFatal(err)

	var wg sync.WaitGroup
	c := make(chan utils.UpdateError)
	defer utils.LogComplete(c, dataType)

	totalCompanies := len(companies)
	for i := 0; i < totalCompanies; {
		batchSize := maxConcurrent
		if i+maxConcurrent > totalCompanies {
			batchSize = totalCompanies - i
		}

		for _, company := range companies[i : i+batchSize] {
			if company.Exchange.String != "UPCOM" &&
				company.Exchange.String != "HOSE" &&
				company.Exchange.String != "HNX" {
				continue
			}

			wg.Add(1)
			go func() {
				update(company, c)
				defer wg.Done()
			}()
		}
		i += batchSize

		wg.Wait()
		time.Sleep(waitTime)
	}
	close(c)
}
