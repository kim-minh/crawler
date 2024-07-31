package stfetch

import "github.com/jackc/pgx/v5/pgtype"

type company struct {
	Data []struct {
		Ticker      pgtype.Text `json:"code"`
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

type profile struct {
	ID                 any         `json:"id"`
	CompanyName        pgtype.Text `json:"companyName"`
	Ticker             any         `json:"ticker"`
	CompanyProfile     pgtype.Text `json:"companyProfile"`
	HistoryDev         pgtype.Text `json:"historyDev"`
	CompanyPromise     pgtype.Text `json:"companyPromise"`
	BusinessRisk       pgtype.Text `json:"businessRisk"`
	KeyDevelopments    pgtype.Text `json:"keyDevelopments"`
	BusinessStrategies pgtype.Text `json:"businessStrategies"`
}

type shareholders struct {
	Data []struct {
		No         int32         `json:"no"`
		Ticker     pgtype.Text   `json:"ticker"`
		Name       pgtype.Text   `json:"name"`
		OwnPercent pgtype.Float4 `json:"ownPercent"`
	} `json:"listShareHolder"`
}

type insiderdeals struct {
	Data []struct {
		No            int32         `json:"no"`
		Ticker        pgtype.Text   `json:"ticker"`
		AnDate        string        `json:"anDate"`
		DealingMethod int           `json:"dealingMethod"`
		DealingAction string        `json:"dealingAction"`
		Quantity      pgtype.Float4 `json:"quantity"`
		Price         pgtype.Float4 `json:"price"`
		Ratio         pgtype.Float4 `json:"ratio"`
	} `json:"listInsiderDealing"`
}

type subsidiaries struct {
	Data []struct {
		No          int32         `json:"no"`
		Ticker      pgtype.Text   `json:"ticker"`
		CompanyName pgtype.Text   `json:"companyName"`
		OwnPercent  pgtype.Float4 `json:"ownPercent"`
	} `json:"listSubCompany"`
}

type officers struct {
	Data []struct {
		No         int32         `json:"no"`
		Ticker     pgtype.Text   `json:"ticker"`
		Name       pgtype.Text   `json:"name"`
		Position   pgtype.Text   `json:"position"`
		OwnPercent pgtype.Float4 `json:"ownPercent"`
	} `json:"listKeyOfficer"`
}

type events struct {
	Data []struct {
		Rsi                pgtype.Float4 `json:"rsi"`
		Rs                 pgtype.Float4 `json:"rs"`
		ID                 pgtype.Int4   `json:"id"`
		Ticker             pgtype.Text   `json:"ticker"`
		Price              pgtype.Int4   `json:"price"`
		PriceChange        pgtype.Int4   `json:"priceChange"`
		PriceChangeRatio   pgtype.Float4 `json:"priceChangeRatio"`
		PriceChangeRatio1M pgtype.Float4 `json:"priceChangeRatio1M"`
		EventName          pgtype.Text   `json:"eventName"`
		EventCode          pgtype.Text   `json:"eventCode"`
		NotifyDate         pgtype.Text   `json:"notifyDate"`
		ExerDate           pgtype.Text   `json:"exerDate"`
		RegFinalDate       pgtype.Text   `json:"regFinalDate"`
		ExRigthDate        pgtype.Text   `json:"exRigthDate"`
		EventDesc          pgtype.Text   `json:"eventDesc"`
	} `json:"listEventNews"`
}

type news struct {
	Data []struct {
		Rsi                pgtype.Float4 `json:"rsi"`
		Rs                 pgtype.Float4 `json:"rs"`
		Ticker             pgtype.Text   `json:"ticker"`
		Price              pgtype.Int4   `json:"price"`
		PriceChange        pgtype.Int4   `json:"priceChange"`
		PriceChangeRatio   pgtype.Float4 `json:"priceChangeRatio"`
		PriceChangeRatio1M pgtype.Float4 `json:"priceChangeRatio1M"`
		ID                 pgtype.Int4   `json:"id"`
		Title              pgtype.Text   `json:"title"`
		Source             pgtype.Text   `json:"source"`
		PublishDate        pgtype.Text   `json:"publishDate"`
	} `json:"listActivityNews"`
}

type dividends struct {
	Data []struct {
		ExerciseDate           pgtype.Text   `json:"exerciseDate"`
		CashYear               pgtype.Int2   `json:"cashYear"`
		CashDividendPercentage pgtype.Float4 `json:"cashDividendPercentage"`
		IssueMethod            pgtype.Text   `json:"issueMethod"`
		No                     pgtype.Int4   `json:"no"`
		Ticker                 pgtype.Text   `json:"ticker"`
	} `json:"listDividendPaymentHis"`
}
