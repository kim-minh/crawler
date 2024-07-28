BEGIN;

CREATE TABLE IF NOT EXISTS companies
(
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    fullname_vi TEXT,
    company_type INTEGER,
    exchange TEXT,
    ticker TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS overview
(
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    company_id INTEGER UNIQUE,
    delta_in_month REAL,
    delta_in_week REAL,
    delta_in_year REAL,
    established_year SMALLINT,
    foreign_percent REAL,
    industry_id INTEGER,
    industry_id_v2 INTEGER,
    issue_share REAL,
    number_of_employees INTEGER,
    number_of_shareholders INTEGER,
    outstanding_share REAL,
    stock_rating REAL,
    company_type TEXT,
    exchange TEXT,
    industry TEXT,
    industry_en TEXT,
    short_name TEXT,
    website TEXT
);

CREATE TABLE IF NOT EXISTS profile
(
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    company_id INTEGER UNIQUE,
    business_risk TEXT,
    business_strategies TEXT,
    company_name TEXT,
    history_dev TEXT,
    key_developments TEXT,
    profile TEXT,
    promise TEXT
);

CREATE TABLE IF NOT EXISTS large_shareholders
(
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    company_id INTEGER,
    share_own_percent REAL,
    shareholder TEXT
);

CREATE TABLE IF NOT EXISTS insider_deals
(
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    company_id INTEGER,
    deal_price INTEGER,
    deal_quantity INTEGER,
    deal_ratio REAL,
    deal_announce_date TIMESTAMPTZ,
    deal_action TEXT,
    deal_method TEXT
);

CREATE TABLE IF NOT EXISTS subsidiaries
(
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    company_id INTEGER,
    own_percent REAL,
    name TEXT
);

CREATE TABLE IF NOT EXISTS officers
(
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    company_id INTEGER,
    own_percent REAL,
    name TEXT,
    "position" TEXT
);

CREATE TABLE IF NOT EXISTS events
(
    id INTEGER PRIMARY KEY,
    company_id INTEGER,
    price INTEGER,
    price_change INTEGER,
    price_change_ratio REAL,
    monthly_price_change_ratio REAL,
    ex_rights_date TIMESTAMPTZ,
    exercise_date TIMESTAMPTZ,
    notify_date TIMESTAMPTZ,
    registration_final_date TIMESTAMPTZ,
    event_code TEXT,
    event_name TEXT,
    event_description TEXT,
    rsi REAL,
    rs REAL
);

CREATE TABLE IF NOT EXISTS news
(
    id INTEGER PRIMARY KEY,
    company_id INTEGER,
    price INTEGER,
    price_change INTEGER,
    price_change_ratio REAL,
    monthly_price_change_ratio REAL,
    publish_date TIMESTAMPTZ,
    source TEXT,
    title TEXT,
    rsi REAL,
    rs REAL
);

CREATE TABLE IF NOT EXISTS dividends
(
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    company_id INTEGER,
    exercise_date TIMESTAMPTZ,
    cash_year SMALLINT,
    cash_dividend_percentage REAL,
    issue_method TEXT
);

CREATE TABLE IF NOT EXISTS balance_sheet
(
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    company_id INTEGER,
    quarter SMALLINT,
    year SMALLINT,
    asset INTEGER,
    short_asset INTEGER,
    fixed_asset INTEGER,
    long_asset INTEGER,
    other_asset INTEGER,
    debt INTEGER,
    short_debt INTEGER,
    long_debt INTEGER,
    other_debt INTEGER,
    deposit INTEGER,
    central_bank_deposit INTEGER,
    other_bank_deposit INTEGER,
    customer_loan INTEGER,
    net_customer_loan INTEGER,
    bad_loan INTEGER,
    other_bank_loan INTEGER,
    cash INTEGER,
    short_invest INTEGER,
    short_receivable INTEGER,
    inventory INTEGER,
    equity INTEGER,
    capital INTEGER,
    stock_invest INTEGER,
    provision INTEGER,
    other_bank_credit INTEGER,
    owe_other_bank INTEGER,
    owe_central_bank INTEGER,
    valuable_paper INTEGER,
    payable_interest INTEGER,
    receivable_interest INTEGER,
    fund INTEGER,
    un_distributed_income INTEGER,
    minor_shareholder_profit INTEGER,
    payable INTEGER
);

CREATE TABLE IF NOT EXISTS income_statement
(
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    company_id INTEGER,
    quarter SMALLINT,
    year SMALLINT,
    revenue INTEGER,
    year_revenue_growth REAL,
    quarter_revenue_growth REAL,
    cost_of_good_sold INTEGER,
    gross_profit INTEGER,
    operation_expense INTEGER,
    operation_profit INTEGER,
    year_operation_profit_growth REAL,
    quarter_operation_profit_growth REAL,
    interest_expense INTEGER,
    pre_tax_profit INTEGER,
    post_tax_profit INTEGER,
    shareholder_income INTEGER,
    year_shareholder_income_growth REAL,
    quarter_shareholder_income_growth REAL,
    invest_profit INTEGER,
    service_profit INTEGER,
    other_profit INTEGER,
    provision_expense INTEGER,
    operation_income INTEGER,
    ebitda INTEGER
);

CREATE TABLE IF NOT EXISTS cash_flow
(
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    company_id INTEGER,
    quarter SMALLINT,
    year SMALLINT,
    invest_cost INTEGER,
    from_invest INTEGER,
    from_financial INTEGER,
    from_sale INTEGER,
    free_cash_flow INTEGER
);

CREATE TABLE IF NOT EXISTS ratio
(
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    company_id INTEGER,
    quarter SMALLINT,
    year SMALLINT,
    price_to_earning REAL,
    price_to_book REAL,
    value_before_ebitda REAL,
    dividend REAL,
    roe REAL,
    roa REAL,
    days_receivable INTEGER,
    days_inventory INTEGER,
    days_payable INTEGER,
    ebit_on_interest REAL,
    earning_per_share INTEGER,
    book_value_per_share INTEGER,
    interest_margin REAL,
    non_interest_on_toi REAL,
    bad_debt_percentage REAL,
    provision_on_bad_debt REAL,
    cost_of_financing REAL,
    equity_on_total_asset REAL,
    equity_on_loan REAL,
    cost_to_income REAL,
    equity_on_liability REAL,
    current_payment REAL,
    quick_payment REAL,
    eps_change REAL,
    ebitda_on_stock INTEGER,
    gross_profit_margin REAL,
    operating_profit_margin REAL,
    post_tax_margin REAL,
    debt_on_equity REAL,
    debt_on_asset REAL,
    debt_on_ebitda REAL,
    short_on_long_debt REAL,
    asset_on_equity REAL,
    capital_balance REAL,
    cash_on_equity REAL,
    cash_on_capitalize REAL,
    cash_circulation REAL,
    revenue_on_work_capital REAL,
    capex_on_fixed_asset REAL,
    revenue_on_asset REAL,
    post_tax_on_pre_tax REAL,
    ebit_on_revenue REAL,
    preTax_on_ebit REAL,
    pre_provision_on_toi REAL,
    post_tax_on_toi REAL,
    loan_on_earn_asset REAL,
    loan_on_asset REAL,
    loan_on_deposit REAL,
    deposit_on_earn_asset REAL,
    bad_debt_on_asset REAL,
    liquidity_on_liability REAL,
    payable_on_equity REAL,
    cancel_debt REAL,
    ebitda_on_stock_change REAL,
    book_value_per_share_change REAL,
    credit_growth REAL
);


CREATE TABLE IF NOT EXISTS stock_historical
(
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    company_id INTEGER,
    close INTEGER,
    high INTEGER,
    low INTEGER,
    open INTEGER,
    volume INTEGER,
    time TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS stock_intraday
(
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    company_id INTEGER,
    price REAL,
    volume INTEGER,
    order_type TEXT,
    previous_price_change REAL,
    time TIMESTAMPTZ
);

-- CREATE TABLE IF NOT EXISTS stock_indices
-- (
--     id serial NOT NULL,
--     index_name text,
--     CONSTRAINT stock_indices_pkey PRIMARY KEY (id)
-- );

-- CREATE TABLE IF NOT EXISTS stock_indices_companies
-- (
--     companies_id INTEGER NOT NULL,
--     stock_indices_id INTEGER NOT NULL,
--     CONSTRAINT stock_indices_companies_pkey PRIMARY KEY (companies_id, stock_indices_id)
-- );

-- CREATE TABLE IF NOT EXISTS ticker_price_volatility
-- (
--     company_id INTEGER,
--     id serial NOT NULL,
--     ticker_highest_price text,
--     ticker_highest_price_percent text,
--     ticker_lowest_price text,
--     ticker_lowest_price_percent text,
--     CONSTRAINT ticker_price_volatility_pkey PRIMARY KEY (id),
--     CONSTRAINT ticker_price_volatility_company_id_key UNIQUE (company_id)
-- );

ALTER TABLE IF EXISTS overview
    ADD CONSTRAINT fkkke2ivq536ntcq3uf971cinvs FOREIGN KEY (company_id)
    REFERENCES companies (id) MATCH SIMPLE
    ON UPDATE CASCADE
    ON DELETE NO ACTION;

ALTER TABLE IF EXISTS profile
    ADD CONSTRAINT fkq8yb3fsal2ki2kq0y1uaeanbv FOREIGN KEY (company_id)
    REFERENCES companies (id) MATCH SIMPLE
    ON UPDATE CASCADE
    ON DELETE NO ACTION;

ALTER TABLE IF EXISTS large_shareholders
    ADD CONSTRAINT fk68kwgcstm9agktuix7412f1nd FOREIGN KEY (company_id)
    REFERENCES companies (id) MATCH SIMPLE
    ON UPDATE CASCADE
    ON DELETE NO ACTION;

ALTER TABLE IF EXISTS insider_deals
    ADD CONSTRAINT fk2f327y1ev9x8skio8pt3eyntc FOREIGN KEY (company_id)
    REFERENCES companies (id) MATCH SIMPLE
    ON UPDATE CASCADE
    ON DELETE NO ACTION;

ALTER TABLE IF EXISTS subsidiaries
    ADD CONSTRAINT fki00iunoyfjax2o1f8it06o5gf FOREIGN KEY (company_id)
    REFERENCES companies (id) MATCH SIMPLE
    ON UPDATE CASCADE
    ON DELETE NO ACTION;

ALTER TABLE IF EXISTS officers
    ADD CONSTRAINT fklr36u6croi4qsm6o5natkrbi3 FOREIGN KEY (company_id)
    REFERENCES companies (id) MATCH SIMPLE
    ON UPDATE CASCADE
    ON DELETE NO ACTION;

ALTER TABLE IF EXISTS events
    ADD CONSTRAINT fkpndphgrrt2p3rr01e9ymfwx4k FOREIGN KEY (company_id)
    REFERENCES companies (id) MATCH SIMPLE
    ON UPDATE CASCADE
    ON DELETE NO ACTION;

ALTER TABLE IF EXISTS news
    ADD CONSTRAINT fkj0xjflx1i92l78pxehr0t3or0 FOREIGN KEY (company_id)
    REFERENCES companies (id) MATCH SIMPLE
    ON UPDATE CASCADE
    ON DELETE NO ACTION;

ALTER TABLE IF EXISTS dividends
    ADD CONSTRAINT fkmqw4w4pnq15r0vgudmilbnqis FOREIGN KEY (company_id)
    REFERENCES companies (id) MATCH SIMPLE
    ON UPDATE CASCADE
    ON DELETE NO ACTION;

ALTER TABLE IF EXISTS balance_sheet
    ADD CONSTRAINT fkk3vw9pj6rm6he5g36pjkenss1 FOREIGN KEY (company_id)
    REFERENCES companies (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION;
CREATE INDEX IF NOT EXISTS balance_sheet_company_id_key
    ON balance_sheet(company_id);

ALTER TABLE IF EXISTS income_statement
    ADD CONSTRAINT fkao43vi7gt01d1gl1knhikkyaa FOREIGN KEY (company_id)
    REFERENCES companies (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION;
CREATE INDEX IF NOT EXISTS income_statement_company_id_key
    ON income_statement(company_id);

ALTER TABLE IF EXISTS cash_flow
    ADD CONSTRAINT fkjgim8e6bj8m597jpmu0x5gn5c FOREIGN KEY (company_id)
    REFERENCES companies (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION;
CREATE INDEX IF NOT EXISTS cash_flow_company_id_key
    ON cash_flow(company_id);


ALTER TABLE IF EXISTS ratio
    ADD CONSTRAINT fkiru0von3m59nmdcu0e6ofpay3 FOREIGN KEY (company_id)
    REFERENCES companies (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION;
CREATE INDEX IF NOT EXISTS ratio_company_id_key
    ON ratio(company_id);


ALTER TABLE IF EXISTS stock_historical
    ADD CONSTRAINT fkfvi1rkoojsesrc9pc7khrfpth FOREIGN KEY (company_id)
    REFERENCES companies (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION;


ALTER TABLE IF EXISTS stock_intraday
    ADD CONSTRAINT fksu5dlmy8u3f7fwi7oi471jud7 FOREIGN KEY (company_id)
    REFERENCES companies (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION;


-- ALTER TABLE IF EXISTS stock_indices_companies
--     ADD CONSTRAINT fkd8347mlarhi4cdjkp96rpx1vv FOREIGN KEY (stock_indices_id)
--     REFERENCES stock_indices (id) MATCH SIMPLE
--     ON UPDATE NO ACTION
--     ON DELETE NO ACTION;


-- ALTER TABLE IF EXISTS stock_indices_companies
--     ADD CONSTRAINT fktbbkx4xhvomfm71ku7xtc5nn2 FOREIGN KEY (companies_id)
--     REFERENCES companies (id) MATCH SIMPLE
--     ON UPDATE NO ACTION
--     ON DELETE NO ACTION;
-- ALTER TABLE IF EXISTS ticker_price_volatility
--     ADD CONSTRAINT fke9rka887uybq1amwjfdhugc3v FOREIGN KEY (company_id)
--     REFERENCES companies (id) MATCH SIMPLE
--     ON UPDATE NO ACTION
--     ON DELETE NO ACTION;
-- CREATE INDEX IF NOT EXISTS ticker_price_volatility_company_id_key
--     ON ticker_price_volatility(company_id);
END;
