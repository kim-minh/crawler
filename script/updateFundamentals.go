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

		year, atoiErr := strconv.Atoi(overview.EstablishedYear)
		utils.LogError(atoiErr)
		establishYear := pgtype.Int2{Int16: int16(year), Valid: true}

		id, atoiErr := strconv.Atoi(overview.IndustryIDv2)
		utils.LogError(atoiErr)
		industryIDV2 := pgtype.Int4{Int32: int32(id), Valid: true}

		err = queries.CreateOverview(ctx, db.CreateOverviewParams{
			CompanyID:            company.ID,
			DeltaInMonth:         overview.DeltaInMonth,
			DeltaInWeek:          overview.DeltaInWeek,
			DeltaInYear:          overview.DeltaInYear,
			EstablishedYear:      establishYear,
			ForeignPercent:       overview.ForeignPercent,
			IndustryID:           overview.IndustryID,
			IndustryIDV2:         industryIDV2,
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
		profile, fetchErr := stfetch.FetchProfile(company.Ticker)
		if fetchErr != nil {
			return
		}

		err = queries.CreateProfile(ctx, db.CreateProfileParams{
			CompanyID:          company.ID,
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

func UpdateShareholders(queries *db.Queries) {
	ctx := context.Background()

	var err error
	updateEach(queries, func(company db.Company) {
		shareholders, fetchErr := stfetch.FetchShareholders(company.Ticker)
		if fetchErr != nil {
			return
		}

		for _, shareholder := range shareholders.Data {
			err = queries.CreateShareholders(ctx, db.CreateShareholdersParams{
				No:              shareholder.No,
				CompanyID:       company.ID,
				ShareOwnPercent: shareholder.OwnPercent,
				Shareholder:     shareholder.Name,
			})
		}
	})
	utils.LogComplete(err, "shareholders")
}

func UpdateInsiderDeals(queries *db.Queries) {
	ctx := context.Background()

	actions := map[string]string{
		"0": "Mua",
		"1": "Bán",
	}
	methods := map[int]string{
		0: "Cổ đông nội bộ",
		1: "Cổ đông lớn",
		2: "Cổ đông sáng lập",
	}

	var err error
	updateEach(queries, func(company db.Company) {
		insiderDeals, fetchErr := stfetch.FetchInsiderDeals(company.Ticker)
		if fetchErr != nil {
			return
		}

		for _, insiderDeal := range insiderDeals.Data {
			dealPrice := pgtype.Int4{Int32: int32(insiderDeal.Price.Float32), Valid: true}
			dealQuantity := pgtype.Int4{Int32: int32(insiderDeal.Quantity.Float32), Valid: true}
			dealAction := pgtype.Text{String: actions[insiderDeal.DealingAction], Valid: true}
			dealMethod := pgtype.Text{String: methods[insiderDeal.DealingMethod], Valid: true}
			dealAnnounceDate := pgtype.Timestamptz{Time: utils.FormatTime(insiderDeal.AnDate), Valid: true}

			err = queries.CreateInsiderDeals(ctx, db.CreateInsiderDealsParams{
				CompanyID:        company.ID,
				DealPrice:        dealPrice,
				DealQuantity:     dealQuantity,
				DealRatio:        insiderDeal.Ratio,
				DealAnnounceDate: dealAnnounceDate,
				DealAction:       dealAction,
				DealMethod:       dealMethod,
			})
		}
	})
	utils.LogComplete(err, "insider deals")
}
