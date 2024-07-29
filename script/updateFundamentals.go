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

	var err error
	updateEach(queries, func(company db.Company) {
		overview, fetchErr := stfetch.FetchOverview(company.Ticker)
		if fetchErr != nil {
			return
		}

		establishedYear, atoiErr := strconv.Atoi(overview.EstablishedYear.String)
		utils.LogError(atoiErr)
		industryIDV2, atoiErr := strconv.Atoi(overview.IndustryIDv2.String)
		utils.LogError(atoiErr)

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
	})

	defer utils.LogComplete(err, "overview")
}

func UpdateProfile(queries *db.Queries) {
	ctx := context.Background()

	var err error
	updateEach(queries, func(company db.Company) {
		profile, err := stfetch.FetchProfile(company.Ticker)
		if err != nil {
			return
		}

		err = queries.CreateProfile(ctx, db.CreateProfileParams{
			CompanyID:          pgtype.Int4{Int32: company.ID, Valid: true},
			BusinessRisk:       utils.ExtractText(profile.BusinessRisk),
			BusinessStrategies: utils.ExtractText(profile.BusinessStrategies),
			CompanyName:        profile.CompanyName,
			HistoryDev:         utils.ExtractText(profile.HistoryDev),
			KeyDevelopments:    utils.ExtractText(profile.KeyDevelopments),
			Profile:            utils.ExtractText(profile.CompanyProfile),
			Promise:            utils.ExtractText(profile.CompanyPromise),
		})
	})

	defer utils.LogComplete(err, "profile")
}
