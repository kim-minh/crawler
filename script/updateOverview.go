package script

import (
	"context"
	"crawler/db"
	"crawler/stfetch"
	"crawler/utils"
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
)

func UpdateOverview(queries *db.Queries) {
	ctx := context.Background()

	update := func(company db.Company) {
		overview, err := stfetch.FetchOverview(company.Ticker)
		if err != nil {
			return
		}

		establishedYear, err := strconv.Atoi(overview.EstablishedYear.String)
		industryIDV2, err := strconv.Atoi(overview.IndustryIDv2.String)

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

		utils.LogError(err)
	}

	updateEach(queries, update)
}
