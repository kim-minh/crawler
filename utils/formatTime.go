package utils

import "time"

func FormatTime(timeStr string) time.Time {
	layout := "02/01/06"
	loc, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	LogError(err)
	parsedTime, err := time.ParseInLocation(layout, timeStr, loc)
	LogError(err)
	return parsedTime
}
