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
	dataType := "overview"

	updateEach(queries, dataType, func(company db.Company) {
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
		utils.LogInsertError(company.Ticker, dataType, err)
	})
}

func UpdateProfile(queries *db.Queries) {
	ctx := context.Background()
	dataType := "profile"

	updateEach(queries, dataType, func(company db.Company) {
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
		utils.LogInsertError(company.Ticker, dataType, err)
	})
}

func UpdateShareholders(queries *db.Queries) {
	ctx := context.Background()
	var wg sync.WaitGroup
	dataType := "shareholders"

	updateEach(queries, dataType, func(company db.Company) {
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
				utils.LogInsertError(company.Ticker, dataType, err)
			}()
		}
		wg.Wait()
	})
}

func UpdateInsiderDeals(queries *db.Queries) {
	ctx := context.Background()
	var wg sync.WaitGroup
	dataType := "insider deals"

	actions := map[string]string{
		"0": "Mua",
		"1": "Bán",
	}
	methods := map[int]string{
		0: "Cổ đông nội bộ",
		1: "Cổ đông lớn",
		2: "Cổ đông sáng lập",
	}

	updateEach(queries, dataType, func(company db.Company) {
		insiderDeals, fetchErr := stfetch.FetchInsiderDeals(company.Ticker)
		if fetchErr != nil {
			return
		}

		for _, insiderDeal := range insiderDeals.Data {
			wg.Add(1)
			go func() {
				defer wg.Done()

				dealPrice := pgtype.Int4{Int32: int32(insiderDeal.Price), Valid: true}
				dealQuantity := pgtype.Int4{Int32: int32(insiderDeal.Quantity), Valid: true}
				dealAction := pgtype.Text{String: actions[insiderDeal.DealingAction], Valid: true}
				dealMethod := pgtype.Text{String: methods[insiderDeal.DealingMethod], Valid: true}

				err := queries.CreateInsiderDeal(ctx, db.CreateInsiderDealParams{
					CompanyID:        company.ID,
					DealPrice:        dealPrice,
					DealQuantity:     dealQuantity,
					DealRatio:        insiderDeal.Ratio,
					DealAnnounceDate: utils.FormatDate(insiderDeal.AnDate),
					DealAction:       dealAction,
					DealMethod:       dealMethod,
				})
				utils.LogInsertError(company.Ticker, dataType, err)
			}()
		}
		wg.Wait()
	})
}

func UpdateSubsidiaries(queries *db.Queries) {
	ctx := context.Background()
	var wg sync.WaitGroup
	dataType := "subsidiaries"

	updateEach(queries, dataType, func(company db.Company) {
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
				utils.LogInsertError(company.Ticker, dataType, err)
			}()
		}
		wg.Wait()
	})
}

func UpdateOfficers(queries *db.Queries) {
	ctx := context.Background()
	var wg sync.WaitGroup
	dataType := "officers"

	updateEach(queries, dataType, func(company db.Company) {
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
				utils.LogInsertError(company.Ticker, dataType, err)
			}()
		}
		wg.Wait()
	})
}

func UpdateEvents(queries *db.Queries) {
	ctx := context.Background()
	var wg sync.WaitGroup
	dataType := "events"

	updateEach(queries, dataType, func(company db.Company) {
		events, fetchErr := stfetch.FetchEvents(company.Ticker)
		if fetchErr != nil {
			return
		}

		for _, event := range events.Data {
			wg.Add(1)
			go func() {
				defer wg.Done()

				err := queries.CreateEvent(ctx, db.CreateEventParams{
					ID:                      event.ID,
					CompanyID:               company.ID,
					Price:                   event.Price,
					PriceChange:             event.PriceChange,
					PriceChangeRatio:        event.PriceChangeRatio,
					MonthlyPriceChangeRatio: event.PriceChangeRatio1M,
					ExRightsDate:            utils.FormatTime(event.ExRigthDate),
					ExerciseDate:            utils.FormatTime(event.ExerDate),
					NotifyDate:              utils.FormatTime(event.NotifyDate),
					RegistrationFinalDate:   utils.FormatTime(event.RegFinalDate),
					EventCode:               event.EventCode,
					EventName:               event.EventName,
					EventDescription:        utils.ExtractText(event.EventDesc),
					Rsi:                     event.Rsi,
					Rs:                      event.Rs,
				})
				utils.LogInsertError(company.Ticker, dataType, err)
			}()
		}
		wg.Wait()
	})
}
