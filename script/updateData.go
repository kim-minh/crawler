package script

import (
	"context"
	"crawler/db"
	"crawler/stfetch"
	"crawler/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func UpdateCompanies(queries *db.Queries) {
	ctx := context.Background()

	companies, _ := stfetch.ListCompanies()
	for _, company := range companies.Data {
		err := queries.CreateCompany(ctx, db.CreateCompanyParams{
			Ticker:      company.Ticker,
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

func UpdateOverview(queries *db.Queries) {
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
		go func() {
			overview, status := stfetch.FetchOverview(company.Ticker)
			if status != http.StatusOK {
				return
			}

			establishedYear, err := strconv.Atoi(overview.EstablishedYear)
			industryIDV2, err := strconv.Atoi(overview.IndustryIDv2)

			err = queries.CreateOverview(ctx, db.CreateOverviewParams{
				CompanyID:            pgtype.Int4{Int32: company.ID, Valid: true},
				DeltaInMonth:         overview.DeltaInMonth,
				DeltaInWeek:          overview.DeltaInWeek,
				DeltaInYear:          overview.DeltaInYear,
				EstablishedYear:      pgtype.Int2{Int16: int16(establishedYear), Valid: true},
				ForeignPercent:       overview.ForeignPercent,
				IndustryID:           overview.IndustryID,
				IndustryIDV2:         pgtype.Int4{Int32: int32(industryIDV2), Valid: true},
				IssueShare:           overview.IssueShare,
				NumberOfEmployees:    overview.NoEmployees,
				NumberOfShareholders: overview.NoShareholders,
				OutstandingShare:     overview.OutstandingShare,
				StockRating:          overview.StockRating,
				CompanyType:          overview.CompanyType,
				Exchange:             overview.Exchange,
				Industry:             overview.Industry,
				IndustryEn:           overview.IndustryEn,
				ShortName:            overview.ShortName,
				Website:              overview.Website,
			})
			if err != nil {
				utils.LogError(err)
			}
		}()
		time.Sleep(700 * time.Millisecond)
	}

	defer utils.LogComplete(err, "overview")
}
