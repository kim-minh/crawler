// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type BalanceSheet struct {
	ID                     int32
	CompanyID              int32
	Quarter                pgtype.Int2
	Year                   pgtype.Int2
	Asset                  pgtype.Int4
	ShortAsset             pgtype.Int4
	FixedAsset             pgtype.Int4
	LongAsset              pgtype.Int4
	OtherAsset             pgtype.Int4
	Debt                   pgtype.Int4
	ShortDebt              pgtype.Int4
	LongDebt               pgtype.Int4
	OtherDebt              pgtype.Int4
	Deposit                pgtype.Int4
	CentralBankDeposit     pgtype.Int4
	OtherBankDeposit       pgtype.Int4
	CustomerLoan           pgtype.Int4
	NetCustomerLoan        pgtype.Int4
	BadLoan                pgtype.Int4
	OtherBankLoan          pgtype.Int4
	Cash                   pgtype.Int4
	ShortInvest            pgtype.Int4
	ShortReceivable        pgtype.Int4
	Inventory              pgtype.Int4
	Equity                 pgtype.Int4
	Capital                pgtype.Int4
	StockInvest            pgtype.Int4
	Provision              pgtype.Int4
	OtherBankCredit        pgtype.Int4
	OweOtherBank           pgtype.Int4
	OweCentralBank         pgtype.Int4
	ValuablePaper          pgtype.Int4
	PayableInterest        pgtype.Int4
	ReceivableInterest     pgtype.Int4
	Fund                   pgtype.Int4
	UnDistributedIncome    pgtype.Int4
	MinorShareholderProfit pgtype.Int4
	Payable                pgtype.Int4
}

type CashFlow struct {
	ID            int32
	CompanyID     int32
	Quarter       pgtype.Int2
	Year          pgtype.Int2
	InvestCost    pgtype.Int4
	FromInvest    pgtype.Int4
	FromFinancial pgtype.Int4
	FromSale      pgtype.Int4
	FreeCashFlow  pgtype.Int4
}

type Company struct {
	ID          int32
	FullnameVi  pgtype.Text
	CompanyType pgtype.Int4
	Exchange    pgtype.Text
	Ticker      string
}

type Dividend struct {
	ID                     int32
	CompanyID              int32
	ExerciseDate           pgtype.Timestamptz
	CashYear               pgtype.Int2
	CashDividendPercentage pgtype.Float4
	IssueMethod            pgtype.Text
}

type Event struct {
	ID                      int32
	CompanyID               int32
	Price                   pgtype.Int4
	PriceChange             pgtype.Int4
	PriceChangeRatio        pgtype.Float4
	MonthlyPriceChangeRatio pgtype.Float4
	ExRightsDate            pgtype.Timestamptz
	ExerciseDate            pgtype.Timestamptz
	NotifyDate              pgtype.Timestamptz
	RegistrationFinalDate   pgtype.Timestamptz
	EventCode               pgtype.Text
	EventName               pgtype.Text
	EventDescription        pgtype.Text
	Rsi                     pgtype.Float4
	Rs                      pgtype.Float4
}

type IncomeStatement struct {
	ID                             int32
	CompanyID                      int32
	Quarter                        pgtype.Int2
	Year                           pgtype.Int2
	Revenue                        pgtype.Int4
	YearRevenueGrowth              pgtype.Float4
	QuarterRevenueGrowth           pgtype.Float4
	CostOfGoodSold                 pgtype.Int4
	GrossProfit                    pgtype.Int4
	OperationExpense               pgtype.Int4
	OperationProfit                pgtype.Int4
	YearOperationProfitGrowth      pgtype.Float4
	QuarterOperationProfitGrowth   pgtype.Float4
	InterestExpense                pgtype.Int4
	PreTaxProfit                   pgtype.Int4
	PostTaxProfit                  pgtype.Int4
	ShareholderIncome              pgtype.Int4
	YearShareholderIncomeGrowth    pgtype.Float4
	QuarterShareholderIncomeGrowth pgtype.Float4
	InvestProfit                   pgtype.Int4
	ServiceProfit                  pgtype.Int4
	OtherProfit                    pgtype.Int4
	ProvisionExpense               pgtype.Int4
	OperationIncome                pgtype.Int4
	Ebitda                         pgtype.Int4
}

type InsiderDeal struct {
	ID               int32
	CompanyID        int32
	DealPrice        pgtype.Int4
	DealQuantity     pgtype.Int4
	DealRatio        pgtype.Float4
	DealAnnounceDate pgtype.Timestamptz
	DealAction       pgtype.Text
	DealMethod       pgtype.Text
}

type LargeShareholder struct {
	ID              int32
	No              int32
	CompanyID       int32
	ShareOwnPercent pgtype.Float4
	Shareholder     pgtype.Text
}

type News struct {
	ID                      int32
	CompanyID               int32
	Price                   pgtype.Int4
	PriceChange             pgtype.Int4
	PriceChangeRatio        pgtype.Float4
	MonthlyPriceChangeRatio pgtype.Float4
	PublishDate             pgtype.Timestamptz
	Source                  pgtype.Text
	Title                   pgtype.Text
	Rsi                     pgtype.Float4
	Rs                      pgtype.Float4
}

type Officer struct {
	ID         int32
	No         int32
	CompanyID  int32
	OwnPercent pgtype.Float4
	Name       pgtype.Text
	Position   pgtype.Text
}

type Overview struct {
	ID                   int32
	CompanyID            int32
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

type Profile struct {
	ID                 int32
	CompanyID          int32
	BusinessRisk       pgtype.Text
	BusinessStrategies pgtype.Text
	CompanyName        pgtype.Text
	HistoryDev         pgtype.Text
	KeyDevelopments    pgtype.Text
	Profile            pgtype.Text
	Promise            pgtype.Text
}

type Ratio struct {
	ID                      int32
	CompanyID               int32
	Quarter                 pgtype.Int2
	Year                    pgtype.Int2
	PriceToEarning          pgtype.Float4
	PriceToBook             pgtype.Float4
	ValueBeforeEbitda       pgtype.Float4
	Dividend                pgtype.Float4
	Roe                     pgtype.Float4
	Roa                     pgtype.Float4
	DaysReceivable          pgtype.Int4
	DaysInventory           pgtype.Int4
	DaysPayable             pgtype.Int4
	EbitOnInterest          pgtype.Float4
	EarningPerShare         pgtype.Int4
	BookValuePerShare       pgtype.Int4
	InterestMargin          pgtype.Float4
	NonInterestOnToi        pgtype.Float4
	BadDebtPercentage       pgtype.Float4
	ProvisionOnBadDebt      pgtype.Float4
	CostOfFinancing         pgtype.Float4
	EquityOnTotalAsset      pgtype.Float4
	EquityOnLoan            pgtype.Float4
	CostToIncome            pgtype.Float4
	EquityOnLiability       pgtype.Float4
	CurrentPayment          pgtype.Float4
	QuickPayment            pgtype.Float4
	EpsChange               pgtype.Float4
	EbitdaOnStock           pgtype.Int4
	GrossProfitMargin       pgtype.Float4
	OperatingProfitMargin   pgtype.Float4
	PostTaxMargin           pgtype.Float4
	DebtOnEquity            pgtype.Float4
	DebtOnAsset             pgtype.Float4
	DebtOnEbitda            pgtype.Float4
	ShortOnLongDebt         pgtype.Float4
	AssetOnEquity           pgtype.Float4
	CapitalBalance          pgtype.Float4
	CashOnEquity            pgtype.Float4
	CashOnCapitalize        pgtype.Float4
	CashCirculation         pgtype.Float4
	RevenueOnWorkCapital    pgtype.Float4
	CapexOnFixedAsset       pgtype.Float4
	RevenueOnAsset          pgtype.Float4
	PostTaxOnPreTax         pgtype.Float4
	EbitOnRevenue           pgtype.Float4
	PretaxOnEbit            pgtype.Float4
	PreProvisionOnToi       pgtype.Float4
	PostTaxOnToi            pgtype.Float4
	LoanOnEarnAsset         pgtype.Float4
	LoanOnAsset             pgtype.Float4
	LoanOnDeposit           pgtype.Float4
	DepositOnEarnAsset      pgtype.Float4
	BadDebtOnAsset          pgtype.Float4
	LiquidityOnLiability    pgtype.Float4
	PayableOnEquity         pgtype.Float4
	CancelDebt              pgtype.Float4
	EbitdaOnStockChange     pgtype.Float4
	BookValuePerShareChange pgtype.Float4
	CreditGrowth            pgtype.Float4
}

type StockHistorical struct {
	ID        int32
	CompanyID int32
	Close     pgtype.Int4
	High      pgtype.Int4
	Low       pgtype.Int4
	Open      pgtype.Int4
	Volume    pgtype.Int4
	Time      pgtype.Timestamptz
}

type StockIndex struct {
	ID        int32
	IndexName pgtype.Text
}

type StockIndicesCompany struct {
	CompaniesID    int32
	StockIndicesID int32
}

type StockIntraday struct {
	ID                  int32
	CompanyID           int32
	Price               pgtype.Float4
	Volume              pgtype.Int4
	OrderType           pgtype.Text
	PreviousPriceChange pgtype.Float4
	Time                pgtype.Timestamptz
}

type Subsidiary struct {
	ID         int32
	No         int32
	CompanyID  int32
	OwnPercent pgtype.Float4
	Name       pgtype.Text
}
