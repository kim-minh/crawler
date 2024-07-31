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
	defer utils.LogComplete(err, "overview")

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
}

func UpdateProfile(queries *db.Queries) {
	ctx := context.Background()

	var err error
	defer utils.LogComplete(err, "profile")

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
}

func UpdateShareholders(queries *db.Queries) {
	ctx := context.Background()

	var err error
	defer utils.LogComplete(err, "shareholders")

	updateEach(queries, func(company db.Company) {
		shareholders, fetchErr := stfetch.FetchShareholders(company.Ticker)
		if fetchErr != nil {
			return
		}

		for _, shareholder := range shareholders.Data {
			err = queries.CreateShareholder(ctx, db.CreateShareholderParams{
				No:              shareholder.No,
				CompanyID:       company.ID,
				ShareOwnPercent: shareholder.OwnPercent,
				Shareholder:     shareholder.Name,
			})
		}
	})
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
	defer utils.LogComplete(err, "insider deals")

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

			err = queries.CreateInsiderDeal(ctx, db.CreateInsiderDealParams{
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
}

func UpdateSubsidiaries(queries *db.Queries) {
	ctx := context.Background()

	var err error
	defer utils.LogComplete(err, "subsidiaries")

	updateEach(queries, func(company db.Company) {
		subsidiaries, fetchErr := stfetch.FetchSubsidiares(company.Ticker)
		if fetchErr != nil {
			return
		}

		for _, subsidiary := range subsidiaries.Data {
			err = queries.CreateSubsidiary(ctx, db.CreateSubsidiaryParams{
				No:         subsidiary.No,
				CompanyID:  company.ID,
				OwnPercent: subsidiary.OwnPercent,
				Name:       subsidiary.CompanyName,
			})
		}
	})
}

func UpdateOfficers(queries *db.Queries) {
	ctx := context.Background()

	var err error
	defer utils.LogComplete(err, "officers")

	updateEach(queries, func(company db.Company) {
		officers, fetchErr := stfetch.FetchOfficers(company.Ticker)
		if fetchErr != nil {
			return
		}

		for _, officer := range officers.Data {
			err = queries.CreateOfficer(ctx, db.CreateOfficerParams{
				No:         officer.No,
				CompanyID:  company.ID,
				OwnPercent: officer.OwnPercent,
				Name:       officer.Name,
				Position:   officer.Position,
			})
		}
	})
}
