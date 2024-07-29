// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createCompany = `-- name: CreateCompany :exec
INSERT INTO companies (
fullname_vi, company_type, exchange, ticker
) VALUES (
  $1, $2, $3, $4
) ON CONFLICT (ticker) DO UPDATE
SET (fullname_vi, company_type, exchange)
= ($1, $2, $3)
`

type CreateCompanyParams struct {
	FullnameVi  pgtype.Text
	CompanyType pgtype.Int4
	Exchange    pgtype.Text
	Ticker      string
}

func (q *Queries) CreateCompany(ctx context.Context, arg CreateCompanyParams) error {
	_, err := q.db.Exec(ctx, createCompany,
		arg.FullnameVi,
		arg.CompanyType,
		arg.Exchange,
		arg.Ticker,
	)
	return err
}

const createOverview = `-- name: CreateOverview :exec
INSERT INTO overview (
company_id,
delta_in_month,
delta_in_week,
delta_in_year,
established_year,
foreign_percent,
industry_id,
industry_id_v2,
issue_share,
number_of_employees,
number_of_shareholders,
outstanding_share,
stock_rating,
company_type,
exchange,
industry,
industry_en,
short_name,
website
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19
) ON CONFLICT (company_id) DO UPDATE
SET (
delta_in_month,
delta_in_week,
delta_in_year,
established_year,
foreign_percent,
industry_id,
industry_id_v2,
issue_share,
number_of_employees,
number_of_shareholders,
outstanding_share,
stock_rating,
company_type,
exchange,
industry,
industry_en,
short_name,
website
) = ($2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)
`

type CreateOverviewParams struct {
	CompanyID            pgtype.Int4
	DeltaInMonth         pgtype.Float4
	DeltaInWeek          pgtype.Float4
	DeltaInYear          pgtype.Float4
	EstablishedYear      pgtype.Int2
	ForeignPercent       pgtype.Float4
	IndustryID           pgtype.Int4
	IndustryIDV2         pgtype.Int4
	IssueShare           pgtype.Float4
	NumberOfEmployees    pgtype.Int4
	NumberOfShareholders pgtype.Int4
	OutstandingShare     pgtype.Float4
	StockRating          pgtype.Float4
	CompanyType          pgtype.Text
	Exchange             pgtype.Text
	Industry             pgtype.Text
	IndustryEn           pgtype.Text
	ShortName            pgtype.Text
	Website              pgtype.Text
}

func (q *Queries) CreateOverview(ctx context.Context, arg CreateOverviewParams) error {
	_, err := q.db.Exec(ctx, createOverview,
		arg.CompanyID,
		arg.DeltaInMonth,
		arg.DeltaInWeek,
		arg.DeltaInYear,
		arg.EstablishedYear,
		arg.ForeignPercent,
		arg.IndustryID,
		arg.IndustryIDV2,
		arg.IssueShare,
		arg.NumberOfEmployees,
		arg.NumberOfShareholders,
		arg.OutstandingShare,
		arg.StockRating,
		arg.CompanyType,
		arg.Exchange,
		arg.Industry,
		arg.IndustryEn,
		arg.ShortName,
		arg.Website,
	)
	return err
}

const listCompanies = `-- name: ListCompanies :many
SELECT id, fullname_vi, company_type, exchange, ticker FROM companies
`

func (q *Queries) ListCompanies(ctx context.Context) ([]Company, error) {
	rows, err := q.db.Query(ctx, listCompanies)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Company
	for rows.Next() {
		var i Company
		if err := rows.Scan(
			&i.ID,
			&i.FullnameVi,
			&i.CompanyType,
			&i.Exchange,
			&i.Ticker,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
