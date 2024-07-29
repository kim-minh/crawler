-- name: CreateCompany :exec
INSERT INTO companies (
fullname_vi, company_type, exchange, ticker
) VALUES (
  $1, $2, $3, $4
) ON CONFLICT (ticker) DO UPDATE
SET (fullname_vi, company_type, exchange)
= ($1, $2, $3);

-- name: ListCompanies :many
SELECT * FROM companies;

-- name: CreateOverview :exec
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
;
