package utils

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func FormatDate(timeStr string) pgtype.Timestamptz {
	if timeStr == "" {
		return pgtype.Timestamptz{}
	}

	layout := "02/01/06"
	loc, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	LogError(err)
	parsedTime, err := time.ParseInLocation(layout, timeStr, loc)
	LogError(err)
	return pgtype.Timestamptz{Time: parsedTime, Valid: true}
}

func FormatTime(timeStr string) pgtype.Timestamptz {
	if timeStr == "" {
		return pgtype.Timestamptz{}
	}

	layout := "2006-01-02 15:04:05"
	loc, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	parsedTime, err := time.ParseInLocation(layout, timeStr, loc)
	LogError(err)
	return pgtype.Timestamptz{Time: parsedTime, Valid: true}
}
