package script

import (
	"context"
	"crawler/db"
	"crawler/utils"
	"time"
)

func updateEach(queries *db.Queries, f func(c db.Company)) {
	ctx := context.Background()

	companies, err := queries.ListCompanies(ctx)
	if err != nil {
		utils.LogError(err)
	}

	for _, company := range companies {
		if company.Exchange.String != "UPCOM" &&
			company.Exchange.String != "HOSE" &&
			company.Exchange.String != "HNX" {
			continue
		}
		go f(company)
		time.Sleep(1 * time.Second)
	}

	defer utils.LogComplete(err, "overview")
}
