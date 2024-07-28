package stfetch

import "github.com/jackc/pgx/v5/pgtype"

type company struct {
	Data []struct {
		Ticker      string      `json:"code"`
		FullnameVi  pgtype.Text `json:"fullname_vi"`
		CompanyType pgtype.Int4 `json:"loaidn"`
		Exchange    pgtype.Text `json:"san"`
	} `json:"data"`
}

type overview struct {
	Exchange         pgtype.Text   `json:"exchange"`
	ShortName        pgtype.Text   `json:"shortName"`
	IndustryID       pgtype.Int4   `json:"industryID"`
	IndustryIDv2     string        `json:"industryIDv2"`
	IndustryIDLevel2 pgtype.Text   `json:"industryIdLevel2"`
	IndustryIDLevel4 pgtype.Text   `json:"industryIdLevel4"`
	Industry         pgtype.Text   `json:"industry"`
	IndustryEn       pgtype.Text   `json:"industryEn"`
	EstablishedYear  string        `json:"establishedYear"`
	NoEmployees      pgtype.Int4   `json:"noEmployees"`
	NoShareholders   pgtype.Int4   `json:"noShareholders"`
	ForeignPercent   pgtype.Float4 `json:"foreignPercent"`
	Website          pgtype.Text   `json:"website"`
	StockRating      pgtype.Float4 `json:"stockRating"`
	DeltaInWeek      pgtype.Float4 `json:"deltaInWeek"`
	DeltaInMonth     pgtype.Float4 `json:"deltaInMonth"`
	DeltaInYear      pgtype.Float4 `json:"deltaInYear"`
	OutstandingShare pgtype.Float4 `json:"outstandingShare"`
	IssueShare       pgtype.Float4 `json:"issueShare"`
	CompanyType      pgtype.Text   `json:"companyType"`
	Ticker           pgtype.Text   `json:"ticker"`
}
