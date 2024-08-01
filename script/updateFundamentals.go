package script

import (
	"context"
	"crawler/db"
	"crawler/stfetch"
	"crawler/utils"
	"strconv"
	"sync"

	"github.com/jackc/pgx/v5/pgtype"
)

func UpdateOverview(queries *db.Queries) {
	ctx := context.Background()

	updateEach(queries, "overview", func(company db.Company, c chan<- utils.UpdateError) {
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

		err := queries.CreateOverview(ctx, db.CreateOverviewParams{
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

		if err != nil {
			c <- utils.UpdateError{Ticker: company.Ticker, Error: err}
		}
	})
}

func UpdateProfile(queries *db.Queries) {
	ctx := context.Background()

	updateEach(queries, "profile", func(company db.Company, c chan<- utils.UpdateError) {
		profile, fetchErr := stfetch.FetchProfile(company.Ticker)
		if fetchErr != nil {
			return
		}

		err := queries.CreateProfile(ctx, db.CreateProfileParams{
			CompanyID:          company.ID,
			BusinessRisk:       utils.ExtractText(profile.BusinessRisk),
			BusinessStrategies: utils.ExtractText(profile.BusinessStrategies),
			CompanyName:        profile.CompanyName,
			HistoryDev:         utils.ExtractText(profile.HistoryDev),
			KeyDevelopments:    utils.ExtractText(profile.KeyDevelopments),
			Profile:            utils.ExtractText(profile.CompanyProfile),
			Promise:            utils.ExtractText(profile.CompanyPromise),
		})

		if err != nil {
			c <- utils.UpdateError{Ticker: company.Ticker, Error: err}
		}
	})
}

func UpdateShareholders(queries *db.Queries) {
	ctx := context.Background()
	var wg sync.WaitGroup

	updateEach(queries, "shareholders", func(company db.Company, c chan<- utils.UpdateError) {
		shareholders, fetchErr := stfetch.FetchShareholders(company.Ticker)
		if fetchErr != nil {
			return
		}

		for _, shareholder := range shareholders.Data {
			wg.Add(1)
			go func() {
				defer wg.Done()

				err := queries.CreateShareholder(ctx, db.CreateShareholderParams{
					No:              shareholder.No,
					CompanyID:       company.ID,
					ShareOwnPercent: shareholder.OwnPercent,
					Shareholder:     shareholder.Name,
				})

				if err != nil {
					c <- utils.UpdateError{Ticker: company.Ticker, Error: err}
				}
			}()
		}
		wg.Wait()
	})
}

func UpdateInsiderDeals(queries *db.Queries) {
	ctx := context.Background()
	var wg sync.WaitGroup

	actions := map[string]string{
		"0": "Mua",
		"1": "Bán",
	}
	methods := map[int]string{
		0: "Cổ đông nội bộ",
		1: "Cổ đông lớn",
		2: "Cổ đông sáng lập",
	}

	updateEach(queries, "insider deals", func(company db.Company, c chan<- utils.UpdateError) {
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

			wg.Add(1)
			go func() {
				defer wg.Done()

				err := queries.CreateInsiderDeal(ctx, db.CreateInsiderDealParams{
					CompanyID:        company.ID,
					DealPrice:        dealPrice,
					DealQuantity:     dealQuantity,
					DealRatio:        insiderDeal.Ratio,
					DealAnnounceDate: dealAnnounceDate,
					DealAction:       dealAction,
					DealMethod:       dealMethod,
				})

				if err != nil {
					c <- utils.UpdateError{Ticker: company.Ticker, Error: err}
				}
			}()
		}
		wg.Wait()
	})
}

func UpdateSubsidiaries(queries *db.Queries) {
	ctx := context.Background()
	var wg sync.WaitGroup

	updateEach(queries, "subsidiaries", func(company db.Company, c chan<- utils.UpdateError) {
		subsidiaries, fetchErr := stfetch.FetchSubsidiares(company.Ticker)
		if fetchErr != nil {
			return
		}

		for _, subsidiary := range subsidiaries.Data {
			wg.Add(1)
			go func() {
				defer wg.Done()

				err := queries.CreateSubsidiary(ctx, db.CreateSubsidiaryParams{
					No:         subsidiary.No,
					CompanyID:  company.ID,
					OwnPercent: subsidiary.OwnPercent,
					Name:       subsidiary.CompanyName,
				})

				if err != nil {
					c <- utils.UpdateError{Ticker: company.Ticker, Error: err}
				}
			}()
		}
		wg.Wait()
	})
}

func UpdateOfficers(queries *db.Queries) {
	ctx := context.Background()
	var wg sync.WaitGroup

	updateEach(queries, "officers", func(company db.Company, c chan<- utils.UpdateError) {
		officers, fetchErr := stfetch.FetchOfficers(company.Ticker)
		if fetchErr != nil {
			return
		}

		for _, officer := range officers.Data {
			wg.Add(1)
			go func() {
				defer wg.Done()

				err := queries.CreateOfficer(ctx, db.CreateOfficerParams{
					No:         officer.No,
					CompanyID:  company.ID,
					OwnPercent: officer.OwnPercent,
					Name:       officer.Name,
					Position:   officer.Position,
				})

				if err != nil {
					c <- utils.UpdateError{Ticker: company.Ticker, Error: err}
				}
			}()
		}
		wg.Wait()
	})
}
