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
) = ($2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19);

-- name: CreateProfile :exec
INSERT INTO profile (
    company_id,
    business_risk,
    business_strategies,
    company_name,
    history_dev,
    key_developments,
    profile,
    promise
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
ON CONFLICT (company_id) DO UPDATE
SET (
    business_risk,
    business_strategies,
    company_name,
    history_dev,
    key_developments,
    profile,
    promise
) = ($2, $3, $4, $5, $6, $7, $8);

-- name: CreateShareholder :exec
INSERT INTO large_shareholders (
    no,
    company_id,
    share_own_percent,
    shareholder
) VALUES ($1, $2, $3, $4)
ON CONFLICT (no, company_id) DO UPDATE
SET (
    share_own_percent,
    shareholder
) = ($3, $4);

-- name: CreateInsiderDeal :exec
INSERT INTO insider_deals (
    company_id,
    deal_price,
    deal_quantity,
    deal_ratio,
    deal_announce_date,
    deal_action,
    deal_method
) VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: CreateSubsidiary :exec
INSERT INTO subsidiaries (
    no,
    company_id,
    own_percent,
    name
) VALUES ($1, $2, $3, $4)
ON CONFLICT (no, company_id) DO UPDATE
SET (
    own_percent,
    name
) = ($3, $4);

-- name: CreateOfficer :exec
INSERT INTO officers (
    no,
    company_id,
    own_percent,
    name,
    position
) VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (no, company_id) DO UPDATE
SET (
own_percent,
    name,
    position
) = ($3, $4, $5);

-- name: CreateEvent :exec
INSERT INTO events (
id,
company_id,
price,
price_change,
price_change_ratio,
monthly_price_change_ratio,
ex_rights_date,
exercise_date,
notify_date ,
registration_final_date,
event_code,
event_name,
event_description,
rsi,
rs
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15);
