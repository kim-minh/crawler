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

func updateEach(queries *db.Queries, update func(c db.Company)) {
	ctx := context.Background()

	companies, err := queries.ListCompanies(ctx)
	utils.LogFatal(err)

	var wg sync.WaitGroup
	totalGoroutines := len(companies)

	for i := 0; i < totalGoroutines; {
		batchSize := maxConcurrent
		if i+maxConcurrent > totalGoroutines {
			batchSize = totalGoroutines - i
		}

		for _, company := range companies[i : i+batchSize] {
			if company.Exchange.String != "UPCOM" &&
				company.Exchange.String != "HOSE" &&
				company.Exchange.String != "HNX" {
				continue
			}

			wg.Add(1)
			go func() {
				update(company)
				defer wg.Done()
			}()
		}

		i += batchSize
		wg.Wait()
		time.Sleep(waitTime)
	}
}
